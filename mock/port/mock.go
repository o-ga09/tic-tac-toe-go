// Code generated by MockGen. DO NOT EDIT.
// Source: ./gateway/port/interface.go

// Package mock_port is a generated GoMock package.
package mock_port

import (
	domain "go-tic-tac-toe/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGameInterface is a mock of GameInterface interface.
type MockGameInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGameInterfaceMockRecorder
}

// MockGameInterfaceMockRecorder is the mock recorder for MockGameInterface.
type MockGameInterfaceMockRecorder struct {
	mock *MockGameInterface
}

// NewMockGameInterface creates a new mock instance.
func NewMockGameInterface(ctrl *gomock.Controller) *MockGameInterface {
	mock := &MockGameInterface{ctrl: ctrl}
	mock.recorder = &MockGameInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameInterface) EXPECT() *MockGameInterfaceMockRecorder {
	return m.recorder
}

// Display mocks base method.
func (m *MockGameInterface) Display(b *domain.Board) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Display", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Display indicates an expected call of Display.
func (mr *MockGameInterfaceMockRecorder) Display(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Display", reflect.TypeOf((*MockGameInterface)(nil).Display), b)
}

// Input mocks base method.
func (m *MockGameInterface) Input(arg0 int) (domain.Koma, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Input", arg0)
	ret0, _ := ret[0].(domain.Koma)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Input indicates an expected call of Input.
func (mr *MockGameInterfaceMockRecorder) Input(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Input", reflect.TypeOf((*MockGameInterface)(nil).Input), arg0)
}

// IsWin mocks base method.
func (m *MockGameInterface) IsWin(arg0 domain.Board) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWin", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWin indicates an expected call of IsWin.
func (mr *MockGameInterfaceMockRecorder) IsWin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWin", reflect.TypeOf((*MockGameInterface)(nil).IsWin), arg0)
}
