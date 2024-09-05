// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/repositories/users.go
//
// Generated by this command:
//
//	mockgen -source=internal/ports/repositories/users.go -destination=internal/repositories/tests/mocks/users_repository_mock.go -package=repository_mocks
//

// Package repository_mocks is a generated GoMock package.
package repository_mocks

import (
	reflect "reflect"

	models "github.com/Arthur-Conti/keepass-project/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUsersRepositoryInterface is a mock of UsersRepositoryInterface interface.
type MockUsersRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryInterfaceMockRecorder
}

// MockUsersRepositoryInterfaceMockRecorder is the mock recorder for MockUsersRepositoryInterface.
type MockUsersRepositoryInterfaceMockRecorder struct {
	mock *MockUsersRepositoryInterface
}

// NewMockUsersRepositoryInterface creates a new mock instance.
func NewMockUsersRepositoryInterface(ctrl *gomock.Controller) *MockUsersRepositoryInterface {
	mock := &MockUsersRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepositoryInterface) EXPECT() *MockUsersRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsersRepositoryInterface) CreateUser(arg0 models.UserModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsersRepositoryInterfaceMockRecorder) CreateUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).CreateUser), arg0)
}

// GetUserByID mocks base method.
func (m *MockUsersRepositoryInterface) GetUserByID(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUsersRepositoryInterfaceMockRecorder) GetUserByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).GetUserByID), arg0)
}

// ListAllUsers mocks base method.
func (m *MockUsersRepositoryInterface) ListAllUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllUsers indicates an expected call of ListAllUsers.
func (mr *MockUsersRepositoryInterfaceMockRecorder) ListAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllUsers", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).ListAllUsers))
}

// UpdateUser mocks base method.
func (m *MockUsersRepositoryInterface) UpdateUser(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUser", arg0)
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUsersRepositoryInterfaceMockRecorder) UpdateUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).UpdateUser), arg0)
}