// Code generated by MockGen. DO NOT EDIT.
// Source: targetgroup.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
)

// MockTargetGroupsService is a mock of TargetGroupsService interface.
type MockTargetGroupsService struct {
	ctrl     *gomock.Controller
	recorder *MockTargetGroupsServiceMockRecorder
}

// MockTargetGroupsServiceMockRecorder is the mock recorder for MockTargetGroupsService.
type MockTargetGroupsServiceMockRecorder struct {
	mock *MockTargetGroupsService
}

// NewMockTargetGroupsService creates a new mock instance.
func NewMockTargetGroupsService(ctrl *gomock.Controller) *MockTargetGroupsService {
	mock := &MockTargetGroupsService{ctrl: ctrl}
	mock.recorder = &MockTargetGroupsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetGroupsService) EXPECT() *MockTargetGroupsServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTargetGroupsService) Create(tg resources.TargetGroup, params resources.QueryParams) (*resources.TargetGroup, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", tg, params)
	ret0, _ := ret[0].(*resources.TargetGroup)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockTargetGroupsServiceMockRecorder) Create(tg, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTargetGroupsService)(nil).Create), tg, params)
}

// Delete mocks base method.
func (m *MockTargetGroupsService) Delete(targetGroupId string, params resources.QueryParams) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", targetGroupId, params)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTargetGroupsServiceMockRecorder) Delete(targetGroupId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTargetGroupsService)(nil).Delete), targetGroupId, params)
}

// Get mocks base method.
func (m *MockTargetGroupsService) Get(targetGroupId string, params resources.QueryParams) (*resources.TargetGroup, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", targetGroupId, params)
	ret0, _ := ret[0].(*resources.TargetGroup)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockTargetGroupsServiceMockRecorder) Get(targetGroupId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTargetGroupsService)(nil).Get), targetGroupId, params)
}

// List mocks base method.
func (m *MockTargetGroupsService) List(params resources.ListQueryParams) (resources.TargetGroups, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", params)
	ret0, _ := ret[0].(resources.TargetGroups)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockTargetGroupsServiceMockRecorder) List(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTargetGroupsService)(nil).List), params)
}

// Update mocks base method.
func (m *MockTargetGroupsService) Update(targetGroupId string, input *resources.TargetGroupProperties, params resources.QueryParams) (*resources.TargetGroup, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", targetGroupId, input, params)
	ret0, _ := ret[0].(*resources.TargetGroup)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockTargetGroupsServiceMockRecorder) Update(targetGroupId, input, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTargetGroupsService)(nil).Update), targetGroupId, input, params)
}
