// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tiflow/engine/pkg/client (interfaces: ExecutorClient)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExecutorClient is a mock of ExecutorClient interface.
type MockExecutorClient struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorClientMockRecorder
}

// MockExecutorClientMockRecorder is the mock recorder for MockExecutorClient.
type MockExecutorClientMockRecorder struct {
	mock *MockExecutorClient
}

// NewMockExecutorClient creates a new mock instance.
func NewMockExecutorClient(ctrl *gomock.Controller) *MockExecutorClient {
	mock := &MockExecutorClient{ctrl: ctrl}
	mock.recorder = &MockExecutorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutorClient) EXPECT() *MockExecutorClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockExecutorClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockExecutorClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockExecutorClient)(nil).Close))
}

// DispatchTask mocks base method.
func (m *MockExecutorClient) DispatchTask(arg0 context.Context, arg1 *DispatchTaskArgs, arg2 func(), arg3 func(error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchTask indicates an expected call of DispatchTask.
func (mr *MockExecutorClientMockRecorder) DispatchTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchTask", reflect.TypeOf((*MockExecutorClient)(nil).DispatchTask), arg0, arg1, arg2, arg3)
}

// RemoveResource mocks base method.
func (m *MockExecutorClient) RemoveResource(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveResource indicates an expected call of RemoveResource.
func (mr *MockExecutorClientMockRecorder) RemoveResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResource", reflect.TypeOf((*MockExecutorClient)(nil).RemoveResource), arg0, arg1, arg2)
}