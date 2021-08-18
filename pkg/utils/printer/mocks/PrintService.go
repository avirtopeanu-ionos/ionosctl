// Code generated by MockGen. DO NOT EDIT.
// Source: ../../utils/printer/printer.go

// Package mock_printer is a generated GoMock package.
package mock_printer

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPrintService is a mock of PrintService interface.
type MockPrintService struct {
	ctrl     *gomock.Controller
	recorder *MockPrintServiceMockRecorder
}

// MockPrintServiceMockRecorder is the mock recorder for MockPrintService.
type MockPrintServiceMockRecorder struct {
	mock *MockPrintService
}

// NewMockPrintService creates a new mock instance.
func NewMockPrintService(ctrl *gomock.Controller) *MockPrintService {
	mock := &MockPrintService{ctrl: ctrl}
	mock.recorder = &MockPrintServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrintService) EXPECT() *MockPrintServiceMockRecorder {
	return m.recorder
}

// GetStderr mocks base method.
func (m *MockPrintService) GetStderr() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStderr")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// GetStderr indicates an expected call of GetStderr.
func (mr *MockPrintServiceMockRecorder) GetStderr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStderr", reflect.TypeOf((*MockPrintService)(nil).GetStderr))
}

// GetStdout mocks base method.
func (m *MockPrintService) GetStdout() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdout")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// GetStdout indicates an expected call of GetStdout.
func (mr *MockPrintServiceMockRecorder) GetStdout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdout", reflect.TypeOf((*MockPrintService)(nil).GetStdout))
}

// Print mocks base method.
func (m *MockPrintService) Print(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Print indicates an expected call of Print.
func (mr *MockPrintServiceMockRecorder) Print(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPrintService)(nil).Print), arg0)
}

// SetStderr mocks base method.
func (m *MockPrintService) SetStderr(arg0 io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStderr", arg0)
}

// SetStderr indicates an expected call of SetStderr.
func (mr *MockPrintServiceMockRecorder) SetStderr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStderr", reflect.TypeOf((*MockPrintService)(nil).SetStderr), arg0)
}

// SetStdout mocks base method.
func (m *MockPrintService) SetStdout(arg0 io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStdout", arg0)
}

// SetStdout indicates an expected call of SetStdout.
func (mr *MockPrintServiceMockRecorder) SetStdout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStdout", reflect.TypeOf((*MockPrintService)(nil).SetStdout), arg0)
}

// Verbose mocks base method.
func (m *MockPrintService) Verbose(format string, a ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "Verbose", varargs...)
}

// Verbose indicates an expected call of Verbose.
func (mr *MockPrintServiceMockRecorder) Verbose(format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verbose", reflect.TypeOf((*MockPrintService)(nil).Verbose), varargs...)
}
