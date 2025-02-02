// Code generated by MockGen. DO NOT EDIT.
// Source: nic.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
)

// MockNicsService is a mock of NicsService interface.
type MockNicsService struct {
	ctrl     *gomock.Controller
	recorder *MockNicsServiceMockRecorder
}

// MockNicsServiceMockRecorder is the mock recorder for MockNicsService.
type MockNicsServiceMockRecorder struct {
	mock *MockNicsService
}

// NewMockNicsService creates a new mock instance.
func NewMockNicsService(ctrl *gomock.Controller) *MockNicsService {
	mock := &MockNicsService{ctrl: ctrl}
	mock.recorder = &MockNicsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNicsService) EXPECT() *MockNicsServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNicsService) Create(datacenterId, serverId string, input resources.Nic, params resources.QueryParams) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", datacenterId, serverId, input, params)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockNicsServiceMockRecorder) Create(datacenterId, serverId, input, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNicsService)(nil).Create), datacenterId, serverId, input, params)
}

// Delete mocks base method.
func (m *MockNicsService) Delete(datacenterId, serverId, nicId string, params resources.QueryParams) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", datacenterId, serverId, nicId, params)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockNicsServiceMockRecorder) Delete(datacenterId, serverId, nicId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNicsService)(nil).Delete), datacenterId, serverId, nicId, params)
}

// Get mocks base method.
func (m *MockNicsService) Get(datacenterId, serverId, nicId string, params resources.QueryParams) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", datacenterId, serverId, nicId, params)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockNicsServiceMockRecorder) Get(datacenterId, serverId, nicId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNicsService)(nil).Get), datacenterId, serverId, nicId, params)
}

// List mocks base method.
func (m *MockNicsService) List(datacenterId, serverId string, params resources.ListQueryParams) (resources.Nics, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", datacenterId, serverId, params)
	ret0, _ := ret[0].(resources.Nics)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockNicsServiceMockRecorder) List(datacenterId, serverId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNicsService)(nil).List), datacenterId, serverId, params)
}

// Update mocks base method.
func (m *MockNicsService) Update(datacenterId, serverId, nicId string, input resources.NicProperties, params resources.QueryParams) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", datacenterId, serverId, nicId, input, params)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockNicsServiceMockRecorder) Update(datacenterId, serverId, nicId, input, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNicsService)(nil).Update), datacenterId, serverId, nicId, input, params)
}
