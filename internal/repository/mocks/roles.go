// Code generated by MockGen. DO NOT EDIT.
// Source: roles.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	domain "example.com/m/v2/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockRolesStorage is a mock of RolesStorage interface.
type MockRolesStorage struct {
	ctrl     *gomock.Controller
	recorder *MockRolesStorageMockRecorder
}

// MockRolesStorageMockRecorder is the mock recorder for MockRolesStorage.
type MockRolesStorageMockRecorder struct {
	mock *MockRolesStorage
}

// NewMockRolesStorage creates a new mock instance.
func NewMockRolesStorage(ctrl *gomock.Controller) *MockRolesStorage {
	mock := &MockRolesStorage{ctrl: ctrl}
	mock.recorder = &MockRolesStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRolesStorage) EXPECT() *MockRolesStorageMockRecorder {
	return m.recorder
}

// GetRole mocks base method.
func (m *MockRolesStorage) GetRole(arg0 int) (*domain.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0)
	ret0, _ := ret[0].(*domain.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockRolesStorageMockRecorder) GetRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockRolesStorage)(nil).GetRole), arg0)
}

// GetRoles mocks base method.
func (m *MockRolesStorage) GetRoles() ([]domain.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles")
	ret0, _ := ret[0].([]domain.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles.
func (mr *MockRolesStorageMockRecorder) GetRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockRolesStorage)(nil).GetRoles))
}
