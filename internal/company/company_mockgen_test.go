// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/by-sabbir/company-microservice-rest/internal/company (interfaces: Store)

// Package company is a generated GoMock package.
package company

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteCompany mocks base method.
func (m *MockStore) DeleteCompany(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockStoreMockRecorder) DeleteCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockStore)(nil).DeleteCompany), arg0, arg1)
}

// GetCompany mocks base method.
func (m *MockStore) GetCompany(arg0 context.Context, arg1 string) (Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", arg0, arg1)
	ret0, _ := ret[0].(Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockStoreMockRecorder) GetCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockStore)(nil).GetCompany), arg0, arg1)
}

// PartialUpdateCompany mocks base method.
func (m *MockStore) PartialUpdateCompany(arg0 context.Context, arg1 string, arg2 Company) (Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PartialUpdateCompany", arg0, arg1, arg2)
	ret0, _ := ret[0].(Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PartialUpdateCompany indicates an expected call of PartialUpdateCompany.
func (mr *MockStoreMockRecorder) PartialUpdateCompany(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PartialUpdateCompany", reflect.TypeOf((*MockStore)(nil).PartialUpdateCompany), arg0, arg1, arg2)
}

// PostCompany mocks base method.
func (m *MockStore) PostCompany(arg0 context.Context, arg1 Company) (Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCompany", arg0, arg1)
	ret0, _ := ret[0].(Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostCompany indicates an expected call of PostCompany.
func (mr *MockStoreMockRecorder) PostCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCompany", reflect.TypeOf((*MockStore)(nil).PostCompany), arg0, arg1)
}
