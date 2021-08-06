// Code generated by MockGen. DO NOT EDIT.
// Source: natgateway.go

// Package mock_v6 is a generated GoMock package.
package mock_v6

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v6 "github.com/ionos-cloud/ionosctl/pkg/resources/v6"
)

// MockNatGatewaysService is a mock of NatGatewaysService interface.
type MockNatGatewaysService struct {
	ctrl     *gomock.Controller
	recorder *MockNatGatewaysServiceMockRecorder
}

// MockNatGatewaysServiceMockRecorder is the mock recorder for MockNatGatewaysService.
type MockNatGatewaysServiceMockRecorder struct {
	mock *MockNatGatewaysService
}

// NewMockNatGatewaysService creates a new mock instance.
func NewMockNatGatewaysService(ctrl *gomock.Controller) *MockNatGatewaysService {
	mock := &MockNatGatewaysService{ctrl: ctrl}
	mock.recorder = &MockNatGatewaysServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatGatewaysService) EXPECT() *MockNatGatewaysServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNatGatewaysService) Create(datacenterId string, input v6.NatGateway) (*v6.NatGateway, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", datacenterId, input)
	ret0, _ := ret[0].(*v6.NatGateway)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockNatGatewaysServiceMockRecorder) Create(datacenterId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNatGatewaysService)(nil).Create), datacenterId, input)
}

// CreateFlowLog mocks base method.
func (m *MockNatGatewaysService) CreateFlowLog(datacenterId, natGatewayId string, input v6.FlowLog) (*v6.FlowLog, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlowLog", datacenterId, natGatewayId, input)
	ret0, _ := ret[0].(*v6.FlowLog)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFlowLog indicates an expected call of CreateFlowLog.
func (mr *MockNatGatewaysServiceMockRecorder) CreateFlowLog(datacenterId, natGatewayId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlowLog", reflect.TypeOf((*MockNatGatewaysService)(nil).CreateFlowLog), datacenterId, natGatewayId, input)
}

// CreateRule mocks base method.
func (m *MockNatGatewaysService) CreateRule(datacenterId, natGatewayId string, input v6.NatGatewayRule) (*v6.NatGatewayRule, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRule", datacenterId, natGatewayId, input)
	ret0, _ := ret[0].(*v6.NatGatewayRule)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRule indicates an expected call of CreateRule.
func (mr *MockNatGatewaysServiceMockRecorder) CreateRule(datacenterId, natGatewayId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRule", reflect.TypeOf((*MockNatGatewaysService)(nil).CreateRule), datacenterId, natGatewayId, input)
}

// Delete mocks base method.
func (m *MockNatGatewaysService) Delete(datacenterId, natGatewayId string) (*v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", datacenterId, natGatewayId)
	ret0, _ := ret[0].(*v6.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockNatGatewaysServiceMockRecorder) Delete(datacenterId, natGatewayId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNatGatewaysService)(nil).Delete), datacenterId, natGatewayId)
}

// DeleteFlowLog mocks base method.
func (m *MockNatGatewaysService) DeleteFlowLog(datacenterId, natGatewayId, flowlogId string) (*v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFlowLog", datacenterId, natGatewayId, flowlogId)
	ret0, _ := ret[0].(*v6.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFlowLog indicates an expected call of DeleteFlowLog.
func (mr *MockNatGatewaysServiceMockRecorder) DeleteFlowLog(datacenterId, natGatewayId, flowlogId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFlowLog", reflect.TypeOf((*MockNatGatewaysService)(nil).DeleteFlowLog), datacenterId, natGatewayId, flowlogId)
}

// DeleteRule mocks base method.
func (m *MockNatGatewaysService) DeleteRule(datacenterId, natGatewayId, ruleId string) (*v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRule", datacenterId, natGatewayId, ruleId)
	ret0, _ := ret[0].(*v6.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRule indicates an expected call of DeleteRule.
func (mr *MockNatGatewaysServiceMockRecorder) DeleteRule(datacenterId, natGatewayId, ruleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockNatGatewaysService)(nil).DeleteRule), datacenterId, natGatewayId, ruleId)
}

// Get mocks base method.
func (m *MockNatGatewaysService) Get(datacenterId, natGatewayId string) (*v6.NatGateway, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", datacenterId, natGatewayId)
	ret0, _ := ret[0].(*v6.NatGateway)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockNatGatewaysServiceMockRecorder) Get(datacenterId, natGatewayId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNatGatewaysService)(nil).Get), datacenterId, natGatewayId)
}

// GetFlowLog mocks base method.
func (m *MockNatGatewaysService) GetFlowLog(datacenterId, natGatewayId, flowlogId string) (*v6.FlowLog, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowLog", datacenterId, natGatewayId, flowlogId)
	ret0, _ := ret[0].(*v6.FlowLog)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFlowLog indicates an expected call of GetFlowLog.
func (mr *MockNatGatewaysServiceMockRecorder) GetFlowLog(datacenterId, natGatewayId, flowlogId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowLog", reflect.TypeOf((*MockNatGatewaysService)(nil).GetFlowLog), datacenterId, natGatewayId, flowlogId)
}

// GetRule mocks base method.
func (m *MockNatGatewaysService) GetRule(datacenterId, natGatewayId, ruleId string) (*v6.NatGatewayRule, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRule", datacenterId, natGatewayId, ruleId)
	ret0, _ := ret[0].(*v6.NatGatewayRule)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRule indicates an expected call of GetRule.
func (mr *MockNatGatewaysServiceMockRecorder) GetRule(datacenterId, natGatewayId, ruleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockNatGatewaysService)(nil).GetRule), datacenterId, natGatewayId, ruleId)
}

// List mocks base method.
func (m *MockNatGatewaysService) List(datacenterId string) (v6.NatGateways, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", datacenterId)
	ret0, _ := ret[0].(v6.NatGateways)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockNatGatewaysServiceMockRecorder) List(datacenterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNatGatewaysService)(nil).List), datacenterId)
}

// ListFlowLogs mocks base method.
func (m *MockNatGatewaysService) ListFlowLogs(datacenterId, natGatewayId string) (v6.FlowLogs, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFlowLogs", datacenterId, natGatewayId)
	ret0, _ := ret[0].(v6.FlowLogs)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFlowLogs indicates an expected call of ListFlowLogs.
func (mr *MockNatGatewaysServiceMockRecorder) ListFlowLogs(datacenterId, natGatewayId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFlowLogs", reflect.TypeOf((*MockNatGatewaysService)(nil).ListFlowLogs), datacenterId, natGatewayId)
}

// ListRules mocks base method.
func (m *MockNatGatewaysService) ListRules(datacenterId, natGatewayId string) (v6.NatGatewayRules, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRules", datacenterId, natGatewayId)
	ret0, _ := ret[0].(v6.NatGatewayRules)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRules indicates an expected call of ListRules.
func (mr *MockNatGatewaysServiceMockRecorder) ListRules(datacenterId, natGatewayId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockNatGatewaysService)(nil).ListRules), datacenterId, natGatewayId)
}

// Update mocks base method.
func (m *MockNatGatewaysService) Update(datacenterId, natGatewayId string, input v6.NatGatewayProperties) (*v6.NatGateway, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", datacenterId, natGatewayId, input)
	ret0, _ := ret[0].(*v6.NatGateway)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockNatGatewaysServiceMockRecorder) Update(datacenterId, natGatewayId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNatGatewaysService)(nil).Update), datacenterId, natGatewayId, input)
}

// UpdateFlowLog mocks base method.
func (m *MockNatGatewaysService) UpdateFlowLog(datacenterId, natGatewayId, flowlogId string, input *v6.FlowLogProperties) (*v6.FlowLog, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlowLog", datacenterId, natGatewayId, flowlogId, input)
	ret0, _ := ret[0].(*v6.FlowLog)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateFlowLog indicates an expected call of UpdateFlowLog.
func (mr *MockNatGatewaysServiceMockRecorder) UpdateFlowLog(datacenterId, natGatewayId, flowlogId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlowLog", reflect.TypeOf((*MockNatGatewaysService)(nil).UpdateFlowLog), datacenterId, natGatewayId, flowlogId, input)
}

// UpdateRule mocks base method.
func (m *MockNatGatewaysService) UpdateRule(datacenterId, natGatewayId, ruleId string, input v6.NatGatewayRuleProperties) (*v6.NatGatewayRule, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRule", datacenterId, natGatewayId, ruleId, input)
	ret0, _ := ret[0].(*v6.NatGatewayRule)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateRule indicates an expected call of UpdateRule.
func (mr *MockNatGatewaysServiceMockRecorder) UpdateRule(datacenterId, natGatewayId, ruleId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRule", reflect.TypeOf((*MockNatGatewaysService)(nil).UpdateRule), datacenterId, natGatewayId, ruleId, input)
}
