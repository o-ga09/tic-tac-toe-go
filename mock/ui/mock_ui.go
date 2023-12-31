// Code generated by MockGen. DO NOT EDIT.
// Source: ./UI/ui.go

// Package mock_ui is a generated GoMock package.
package mock_ui

import (
	ui "go-tic-tac-toe/UI"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIODriver is a mock of IODriver interface.
type MockIODriver struct {
	ctrl     *gomock.Controller
	recorder *MockIODriverMockRecorder
}

// MockIODriverMockRecorder is the mock recorder for MockIODriver.
type MockIODriverMockRecorder struct {
	mock *MockIODriver
}

// NewMockIODriver creates a new mock instance.
func NewMockIODriver(ctrl *gomock.Controller) *MockIODriver {
	mock := &MockIODriver{ctrl: ctrl}
	mock.recorder = &MockIODriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIODriver) EXPECT() *MockIODriverMockRecorder {
	return m.recorder
}

// Display mocks base method.
func (m *MockIODriver) Display(arg0 ui.DriverBoard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Display", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Display indicates an expected call of Display.
func (mr *MockIODriverMockRecorder) Display(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Display", reflect.TypeOf((*MockIODriver)(nil).Display), arg0)
}

// Input mocks base method.
func (m *MockIODriver) Input(arg0 io.Reader) (ui.DriverKoma, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Input", arg0)
	ret0, _ := ret[0].(ui.DriverKoma)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Input indicates an expected call of Input.
func (mr *MockIODriverMockRecorder) Input(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Input", reflect.TypeOf((*MockIODriver)(nil).Input), arg0)
}
