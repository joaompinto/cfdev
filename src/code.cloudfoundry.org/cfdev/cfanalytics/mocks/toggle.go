// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cfanalytics (interfaces: Toggle)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockToggle is a mock of Toggle interface
type MockToggle struct {
	ctrl     *gomock.Controller
	recorder *MockToggleMockRecorder
}

// MockToggleMockRecorder is the mock recorder for MockToggle
type MockToggleMockRecorder struct {
	mock *MockToggle
}

// NewMockToggle creates a new mock instance
func NewMockToggle(ctrl *gomock.Controller) *MockToggle {
	mock := &MockToggle{ctrl: ctrl}
	mock.recorder = &MockToggleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockToggle) EXPECT() *MockToggleMockRecorder {
	return m.recorder
}

// Defined mocks base method
func (m *MockToggle) Defined() bool {
	ret := m.ctrl.Call(m, "Defined")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Defined indicates an expected call of Defined
func (mr *MockToggleMockRecorder) Defined() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Defined", reflect.TypeOf((*MockToggle)(nil).Defined))
}

// Get mocks base method
func (m *MockToggle) Get() bool {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockToggleMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockToggle)(nil).Get))
}

// GetProps mocks base method
func (m *MockToggle) GetProps() map[string]interface{} {
	ret := m.ctrl.Call(m, "GetProps")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// GetProps indicates an expected call of GetProps
func (mr *MockToggleMockRecorder) GetProps() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProps", reflect.TypeOf((*MockToggle)(nil).GetProps))
}

// Set mocks base method
func (m *MockToggle) Set(arg0 bool) error {
	ret := m.ctrl.Call(m, "Set", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockToggleMockRecorder) Set(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockToggle)(nil).Set), arg0)
}
