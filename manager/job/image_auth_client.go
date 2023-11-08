package job

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/docker/distribution"
	"github.com/docker/distribution/registry/client/auth"
	"github.com/docker/distribution/registry/client/transport"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/registry"
)

type BasicAuth struct {
	Username string
	Password string
	Auth     string
}

func (b *BasicAuth) token() string {
	auth := fmt.Sprintf("%s:%s", b.Username, b.Password)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type ImageAuthClient struct {
	client           *http.Client
	basic            *BasicAuth
	headerModifier   http.Header
	tokenInterceptor *InterceptorTokenHandler
}

type ImageAuthClientOption func(*ImageAuthClient)

func WithBasicAuth(b *BasicAuth) ImageAuthClientOption {
	return func(o *ImageAuthClient) {
		o.basic = b
	}
}

func WithImageRepo(repo string) ImageAuthClientOption {
	return func(o *ImageAuthClient) {
		o.basic.Auth = repo
	}
}

func WithHeaderModifier(h http.Header) ImageAuthClientOption {
	return func(o *ImageAuthClient) {
		o.headerModifier = h
	}
}

func WithClient(c *http.Client) ImageAuthClientOption {
	return func(o *ImageAuthClient) {
		o.client = c
	}
}

func NewImageAuthClient(image *preheatImage, opts ...ImageAuthClientOption) (*ImageAuthClient, error) {
	d := &ImageAuthClient{}
	for _, opt := range opts {
		opt(d)
	}

	if d.client == nil {
		d.client = &http.Client{}
	}

	if d.client.Transport == nil {
		direct := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}

		d.client.Transport = &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			DialContext:         direct.DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig:     &tls.Config{RootCAs: x509.NewCertPool()},
			DisableKeepAlives:   true,
		}
	}

	baseTransport := d.client.Transport
	authTransport := transport.NewTransport(baseTransport, transport.NewHeaderRequestModifier(d.headerModifier))
	challengeManager, err := registry.PingV2Registry(&url.URL{Scheme: image.protocol, Host: image.domain}, authTransport)
	if err != nil {
		return nil, err
	}

	scope := auth.RepositoryScope{
		Repository: image.name,
		Actions:    []string{"pull"},
	}
	creds := registry.NewStaticCredentialStore(&types.AuthConfig{Username: d.basic.Username, Password: d.basic.Password, Auth: d.basic.token()})
	tokenHandlerOptions := auth.TokenHandlerOptions{
		Transport:   authTransport,
		Credentials: creds,
		Scopes:      []auth.Scope{scope},
		ClientID:    registry.AuthClientID,
	}
	tokenHandler := auth.NewTokenHandlerWithOptions(tokenHandlerOptions)
	basicHandler := auth.NewBasicHandler(creds)
	interceptor := NewInterceptorTokenHandler()
	d.tokenInterceptor = interceptor
	d.client.Transport = transport.NewTransport(baseTransport, auth.NewAuthorizer(challengeManager, tokenHandler, interceptor, basicHandler))
	return d, nil
}

func (d *ImageAuthClient) Do(req *http.Request) (*http.Response, error) {
	return d.client.Do(req)
}

func (d *ImageAuthClient) GetBearerToken() string {
	return d.tokenInterceptor.GetAuthToken()
}

// GetManifestMediaTypeAcceptHeader
// get manifest/tag need accept header
func GetManifestMediaTypeAcceptHeader(h http.Header) http.Header {
	if h == nil {
		h = http.Header{}
	}
	header := h.Clone()
	for _, v := range distribution.ManifestMediaTypes() {
		header.Add("Accept", v)
	}

	return header
}

// InterceptorTokenHandler is a token interceptor
// intercept bearer token from auth handler
type InterceptorTokenHandler struct {
	auth.AuthenticationHandler
	token string
}

func NewInterceptorTokenHandler() *InterceptorTokenHandler {
	return &InterceptorTokenHandler{}
}

func (h *InterceptorTokenHandler) Scheme() string {
	return "bearer"
}

func (h *InterceptorTokenHandler) AuthorizeRequest(req *http.Request, params map[string]string) error {
	h.token = req.Header.Get("Authorization")
	return nil
}

func (h *InterceptorTokenHandler) GetAuthToken() string {
	return h.token
}