// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/start (interfaces: Cache)

// Package mocks is a generated GoMock package.
package mocks

import (
	resource "code.cloudfoundry.org/cfdev/resource"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Sync mocks base method
func (m *MockCache) Sync(arg0 resource.Catalog) error {
	ret := m.ctrl.Call(m, "Sync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockCacheMockRecorder) Sync(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockCache)(nil).Sync), arg0)
}
