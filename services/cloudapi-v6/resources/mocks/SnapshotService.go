// Code generated by MockGen. DO NOT EDIT.
// Source: snapshot.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
)

// MockSnapshotsService is a mock of SnapshotsService interface.
type MockSnapshotsService struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsServiceMockRecorder
}

// MockSnapshotsServiceMockRecorder is the mock recorder for MockSnapshotsService.
type MockSnapshotsServiceMockRecorder struct {
	mock *MockSnapshotsService
}

// NewMockSnapshotsService creates a new mock instance.
func NewMockSnapshotsService(ctrl *gomock.Controller) *MockSnapshotsService {
	mock := &MockSnapshotsService{ctrl: ctrl}
	mock.recorder = &MockSnapshotsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotsService) EXPECT() *MockSnapshotsServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSnapshotsService) Create(datacenterId, volumeId, name, description, licenceType string, secAuthProtection bool, params resources.QueryParams) (*resources.Snapshot, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", datacenterId, volumeId, name, description, licenceType, secAuthProtection, params)
	ret0, _ := ret[0].(*resources.Snapshot)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockSnapshotsServiceMockRecorder) Create(datacenterId, volumeId, name, description, licenceType, secAuthProtection, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSnapshotsService)(nil).Create), datacenterId, volumeId, name, description, licenceType, secAuthProtection, params)
}

// Delete mocks base method.
func (m *MockSnapshotsService) Delete(snapshotId string, params resources.QueryParams) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", snapshotId, params)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSnapshotsServiceMockRecorder) Delete(snapshotId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSnapshotsService)(nil).Delete), snapshotId, params)
}

// Get mocks base method.
func (m *MockSnapshotsService) Get(snapshotId string, params resources.QueryParams) (*resources.Snapshot, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", snapshotId, params)
	ret0, _ := ret[0].(*resources.Snapshot)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockSnapshotsServiceMockRecorder) Get(snapshotId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSnapshotsService)(nil).Get), snapshotId, params)
}

// List mocks base method.
func (m *MockSnapshotsService) List(params resources.ListQueryParams) (resources.Snapshots, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", params)
	ret0, _ := ret[0].(resources.Snapshots)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockSnapshotsServiceMockRecorder) List(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSnapshotsService)(nil).List), params)
}

// Restore mocks base method.
func (m *MockSnapshotsService) Restore(datacenterId, volumeId, snapshotId string, params resources.QueryParams) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", datacenterId, volumeId, snapshotId, params)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Restore indicates an expected call of Restore.
func (mr *MockSnapshotsServiceMockRecorder) Restore(datacenterId, volumeId, snapshotId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockSnapshotsService)(nil).Restore), datacenterId, volumeId, snapshotId, params)
}

// Update mocks base method.
func (m *MockSnapshotsService) Update(snapshotId string, snapshotProp resources.SnapshotProperties, params resources.QueryParams) (*resources.Snapshot, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", snapshotId, snapshotProp, params)
	ret0, _ := ret[0].(*resources.Snapshot)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockSnapshotsServiceMockRecorder) Update(snapshotId, snapshotProp, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSnapshotsService)(nil).Update), snapshotId, snapshotProp, params)
}
