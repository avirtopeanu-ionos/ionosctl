// Code generated by MockGen. DO NOT EDIT.
// Source: lan.go

// Package mock_v6 is a generated GoMock package.
package mock_v6

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v6 "github.com/ionos-cloud/ionosctl/pkg/resources/v6"
)

// MockLansService is a mock of LansService interface.
type MockLansService struct {
	ctrl     *gomock.Controller
	recorder *MockLansServiceMockRecorder
}

// MockLansServiceMockRecorder is the mock recorder for MockLansService.
type MockLansServiceMockRecorder struct {
	mock *MockLansService
}

// NewMockLansService creates a new mock instance.
func NewMockLansService(ctrl *gomock.Controller) *MockLansService {
	mock := &MockLansService{ctrl: ctrl}
	mock.recorder = &MockLansServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLansService) EXPECT() *MockLansServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLansService) Create(datacenterId string, input v6.LanPost) (*v6.LanPost, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", datacenterId, input)
	ret0, _ := ret[0].(*v6.LanPost)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockLansServiceMockRecorder) Create(datacenterId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLansService)(nil).Create), datacenterId, input)
}

// Delete mocks base method.
func (m *MockLansService) Delete(datacenterId, lanId string) (*v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", datacenterId, lanId)
	ret0, _ := ret[0].(*v6.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockLansServiceMockRecorder) Delete(datacenterId, lanId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLansService)(nil).Delete), datacenterId, lanId)
}

// Get mocks base method.
func (m *MockLansService) Get(datacenterId, lanId string) (*v6.Lan, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", datacenterId, lanId)
	ret0, _ := ret[0].(*v6.Lan)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockLansServiceMockRecorder) Get(datacenterId, lanId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLansService)(nil).Get), datacenterId, lanId)
}

// List mocks base method.
func (m *MockLansService) List(datacenterId string) (v6.Lans, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", datacenterId)
	ret0, _ := ret[0].(v6.Lans)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockLansServiceMockRecorder) List(datacenterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLansService)(nil).List), datacenterId)
}

// Update mocks base method.
func (m *MockLansService) Update(datacenterId, lanId string, input v6.LanProperties) (*v6.Lan, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", datacenterId, lanId, input)
	ret0, _ := ret[0].(*v6.Lan)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockLansServiceMockRecorder) Update(datacenterId, lanId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLansService)(nil).Update), datacenterId, lanId, input)
}
