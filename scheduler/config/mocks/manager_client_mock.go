// Code generated by MockGen. DO NOT EDIT.
// Source: ../../pkg/rpc/manager/client/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	manager "d7y.io/dragonfly/v2/internal/rpc/manager"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockManagerClient is a mock of ManagerClient interface.
type MockManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagerClientMockRecorder
}

// MockManagerClientMockRecorder is the mock recorder for MockManagerClient.
type MockManagerClientMockRecorder struct {
	mock *MockManagerClient
}

// NewMockManagerClient creates a new mock instance.
func NewMockManagerClient(ctrl *gomock.Controller) *MockManagerClient {
	mock := &MockManagerClient{ctrl: ctrl}
	mock.recorder = &MockManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerClient) EXPECT() *MockManagerClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockManagerClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockManagerClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManagerClient)(nil).Close))
}

// GetCdnClusterConfig mocks base method.
func (m *MockManagerClient) GetCdnClusterConfig(ctx context.Context, req *manager.GetClusterConfigRequest, opts ...grpc.CallOption) (*manager.CdnConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCdnClusterConfig", varargs...)
	ret0, _ := ret[0].(*manager.CdnConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCdnClusterConfig indicates an expected call of GetCdnClusterConfig.
func (mr *MockManagerClientMockRecorder) GetCdnClusterConfig(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCdnClusterConfig", reflect.TypeOf((*MockManagerClient)(nil).GetCdnClusterConfig), varargs...)
}

// GetSchedulerClusterConfig mocks base method.
func (m *MockManagerClient) GetSchedulerClusterConfig(ctx context.Context, req *manager.GetClusterConfigRequest, opts ...grpc.CallOption) (*manager.SchedulerConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchedulerClusterConfig", varargs...)
	ret0, _ := ret[0].(*manager.SchedulerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedulerClusterConfig indicates an expected call of GetSchedulerClusterConfig.
func (mr *MockManagerClientMockRecorder) GetSchedulerClusterConfig(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulerClusterConfig", reflect.TypeOf((*MockManagerClient)(nil).GetSchedulerClusterConfig), varargs...)
}

// GetSchedulers mocks base method.
func (m *MockManagerClient) GetSchedulers(ctx context.Context, req *manager.GetSchedulersRequest, opts ...grpc.CallOption) (*manager.SchedulerNodes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchedulers", varargs...)
	ret0, _ := ret[0].(*manager.SchedulerNodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedulers indicates an expected call of GetSchedulers.
func (mr *MockManagerClientMockRecorder) GetSchedulers(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulers", reflect.TypeOf((*MockManagerClient)(nil).GetSchedulers), varargs...)
}

// KeepAlive mocks base method.
func (m *MockManagerClient) KeepAlive(ctx context.Context, req *manager.KeepAliveRequest, opts ...grpc.CallOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "KeepAlive", varargs...)
}

// KeepAlive indicates an expected call of KeepAlive.
func (mr *MockManagerClientMockRecorder) KeepAlive(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAlive", reflect.TypeOf((*MockManagerClient)(nil).KeepAlive), varargs...)
}
