// Code generated by MockGen. DO NOT EDIT.
// Source: proxy.go

// Package mock_cron is a generated GoMock package.
package mock_cron

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cron "github.com/robfig/cron"
)

// MockCronProxy is a mock of CronProxy interface.
type MockCronProxy struct {
	ctrl     *gomock.Controller
	recorder *MockCronProxyMockRecorder
}

// MockCronProxyMockRecorder is the mock recorder for MockCronProxy.
type MockCronProxyMockRecorder struct {
	mock *MockCronProxy
}

// NewMockCronProxy creates a new mock instance.
func NewMockCronProxy(ctrl *gomock.Controller) *MockCronProxy {
	mock := &MockCronProxy{ctrl: ctrl}
	mock.recorder = &MockCronProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCronProxy) EXPECT() *MockCronProxyMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCronProxy) Create(timeZone string) (*cron.Cron, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", timeZone)
	ret0, _ := ret[0].(*cron.Cron)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCronProxyMockRecorder) Create(timeZone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCronProxy)(nil).Create), timeZone)
}

// Parse mocks base method.
func (m *MockCronProxy) Parse(spec string) (cron.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", spec)
	ret0, _ := ret[0].(cron.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockCronProxyMockRecorder) Parse(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockCronProxy)(nil).Parse), spec)
}

// Push mocks base method.
func (m *MockCronProxy) Push(c *cron.Cron, time string, call func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Push", c, time, call)
}

// Push indicates an expected call of Push.
func (mr *MockCronProxyMockRecorder) Push(c, time, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockCronProxy)(nil).Push), c, time, call)
}

// Start mocks base method.
func (m *MockCronProxy) Start(c *cron.Cron) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", c)
}

// Start indicates an expected call of Start.
func (mr *MockCronProxyMockRecorder) Start(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCronProxy)(nil).Start), c)
}

// Stop mocks base method.
func (m *MockCronProxy) Stop(c *cron.Cron) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", c)
}

// Stop indicates an expected call of Stop.
func (mr *MockCronProxyMockRecorder) Stop(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCronProxy)(nil).Stop), c)
}
