// Code generated by MockGen. DO NOT EDIT.
// Source: users.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	domain "example.com/m/v2/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUsersStorage is a mock of UsersStorage interface.
type MockUsersStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUsersStorageMockRecorder
}

// MockUsersStorageMockRecorder is the mock recorder for MockUsersStorage.
type MockUsersStorageMockRecorder struct {
	mock *MockUsersStorage
}

// NewMockUsersStorage creates a new mock instance.
func NewMockUsersStorage(ctrl *gomock.Controller) *MockUsersStorage {
	mock := &MockUsersStorage{ctrl: ctrl}
	mock.recorder = &MockUsersStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersStorage) EXPECT() *MockUsersStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockUsersStorage) Add(arg0 domain.User) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockUsersStorageMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUsersStorage)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockUsersStorage) Delete(arg0 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersStorageMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsersStorage)(nil).Delete), arg0)
}

// Edit mocks base method.
func (m *MockUsersStorage) Edit(arg0 domain.User) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Edit indicates an expected call of Edit.
func (mr *MockUsersStorageMockRecorder) Edit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockUsersStorage)(nil).Edit), arg0)
}

// GetUser mocks base method.
func (m *MockUsersStorage) GetUser(arg0 int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUsersStorageMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsersStorage)(nil).GetUser), arg0)
}

// GetUsers mocks base method.
func (m *MockUsersStorage) GetUsers() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUsersStorageMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUsersStorage)(nil).GetUsers))
}
