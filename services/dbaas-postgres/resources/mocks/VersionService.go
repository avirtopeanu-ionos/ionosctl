// Code generated by MockGen. DO NOT EDIT.
// Source: version.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/services/dbaas-postgres/resources"
)

// MockVersionsService is a mock of VersionsService interface.
type MockVersionsService struct {
	ctrl     *gomock.Controller
	recorder *MockVersionsServiceMockRecorder
}

// MockVersionsServiceMockRecorder is the mock recorder for MockVersionsService.
type MockVersionsServiceMockRecorder struct {
	mock *MockVersionsService
}

// NewMockVersionsService creates a new mock instance.
func NewMockVersionsService(ctrl *gomock.Controller) *MockVersionsService {
	mock := &MockVersionsService{ctrl: ctrl}
	mock.recorder = &MockVersionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionsService) EXPECT() *MockVersionsServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVersionsService) Get(clusterId string) (resources.PostgresVersionList, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", clusterId)
	ret0, _ := ret[0].(resources.PostgresVersionList)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockVersionsServiceMockRecorder) Get(clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVersionsService)(nil).Get), clusterId)
}

// List mocks base method.
func (m *MockVersionsService) List() (resources.PostgresVersionList, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(resources.PostgresVersionList)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockVersionsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVersionsService)(nil).List))
}
