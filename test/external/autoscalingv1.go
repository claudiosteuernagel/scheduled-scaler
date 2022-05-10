// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/autoscaling/v1 (interfaces: AutoscalingV1Interface,HorizontalPodAutoscalerInterface)

// Package mock_external is a generated GoMock package.
package mock_external

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/autoscaling/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockAutoscalingV1Interface is a mock of AutoscalingV1Interface interface
type MockAutoscalingV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAutoscalingV1InterfaceMockRecorder
}

// MockAutoscalingV1InterfaceMockRecorder is the mock recorder for MockAutoscalingV1Interface
type MockAutoscalingV1InterfaceMockRecorder struct {
	mock *MockAutoscalingV1Interface
}

// NewMockAutoscalingV1Interface creates a new mock instance
func NewMockAutoscalingV1Interface(ctrl *gomock.Controller) *MockAutoscalingV1Interface {
	mock := &MockAutoscalingV1Interface{ctrl: ctrl}
	mock.recorder = &MockAutoscalingV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAutoscalingV1Interface) EXPECT() *MockAutoscalingV1InterfaceMockRecorder {
	return m.recorder
}

// HorizontalPodAutoscalers mocks base method
func (m *MockAutoscalingV1Interface) HorizontalPodAutoscalers(arg0 string) v11.HorizontalPodAutoscalerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HorizontalPodAutoscalers", arg0)
	ret0, _ := ret[0].(v11.HorizontalPodAutoscalerInterface)
	return ret0
}

// HorizontalPodAutoscalers indicates an expected call of HorizontalPodAutoscalers
func (mr *MockAutoscalingV1InterfaceMockRecorder) HorizontalPodAutoscalers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HorizontalPodAutoscalers", reflect.TypeOf((*MockAutoscalingV1Interface)(nil).HorizontalPodAutoscalers), arg0)
}

// RESTClient mocks base method
func (m *MockAutoscalingV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockAutoscalingV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAutoscalingV1Interface)(nil).RESTClient))
}

// MockHorizontalPodAutoscalerInterface is a mock of HorizontalPodAutoscalerInterface interface
type MockHorizontalPodAutoscalerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHorizontalPodAutoscalerInterfaceMockRecorder
}

// MockHorizontalPodAutoscalerInterfaceMockRecorder is the mock recorder for MockHorizontalPodAutoscalerInterface
type MockHorizontalPodAutoscalerInterfaceMockRecorder struct {
	mock *MockHorizontalPodAutoscalerInterface
}

// NewMockHorizontalPodAutoscalerInterface creates a new mock instance
func NewMockHorizontalPodAutoscalerInterface(ctrl *gomock.Controller) *MockHorizontalPodAutoscalerInterface {
	mock := &MockHorizontalPodAutoscalerInterface{ctrl: ctrl}
	mock.recorder = &MockHorizontalPodAutoscalerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHorizontalPodAutoscalerInterface) EXPECT() *MockHorizontalPodAutoscalerInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockHorizontalPodAutoscalerInterface) Create(arg0 context.Context, arg1 *v1.HorizontalPodAutoscaler, arg2 v10.CreateOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockHorizontalPodAutoscalerInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method
func (m *MockHorizontalPodAutoscalerInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockHorizontalPodAutoscalerInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockHorizontalPodAutoscalerInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.HorizontalPodAutoscalerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscalerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method
func (m *MockHorizontalPodAutoscalerInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockHorizontalPodAutoscalerInterface) Update(arg0 context.Context, arg1 *v1.HorizontalPodAutoscaler, arg2 v10.UpdateOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method
func (m *MockHorizontalPodAutoscalerInterface) UpdateStatus(arg0 context.Context, arg1 *v1.HorizontalPodAutoscaler, arg2 v10.UpdateOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method
func (m *MockHorizontalPodAutoscalerInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Watch), arg0, arg1)
}
