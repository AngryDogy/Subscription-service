// Code generated by MockGen. DO NOT EDIT.
// Source: proxy_client.go

// Package proxy is a generated GoMock package.
package proxy

import (
	entity "dev/master/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateKey mocks base method.
func (m *MockClient) CreateKey(address string) (*entity.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", address)
	ret0, _ := ret[0].(*entity.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockClientMockRecorder) CreateKey(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockClient)(nil).CreateKey), address)
}

// DeleteKey mocks base method.
func (m *MockClient) DeleteKey(address, keyId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", address, keyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockClientMockRecorder) DeleteKey(address, keyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockClient)(nil).DeleteKey), address, keyId)
}
