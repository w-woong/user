// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_usc.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/w-woong/common/dto"
)

// MockUserUsc is a mock of UserUsc interface.
type MockUserUsc struct {
	ctrl     *gomock.Controller
	recorder *MockUserUscMockRecorder
}

// MockUserUscMockRecorder is the mock recorder for MockUserUsc.
type MockUserUscMockRecorder struct {
	mock *MockUserUsc
}

// NewMockUserUsc creates a new mock instance.
func NewMockUserUsc(ctrl *gomock.Controller) *MockUserUsc {
	mock := &MockUserUsc{ctrl: ctrl}
	mock.recorder = &MockUserUscMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsc) EXPECT() *MockUserUscMockRecorder {
	return m.recorder
}

// FindByLoginID mocks base method.
func (m *MockUserUsc) FindByLoginID(ctx context.Context, loginSource, loginID string) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLoginID", ctx, loginSource, loginID)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLoginID indicates an expected call of FindByLoginID.
func (mr *MockUserUscMockRecorder) FindByLoginID(ctx, loginSource, loginID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLoginID", reflect.TypeOf((*MockUserUsc)(nil).FindByLoginID), ctx, loginSource, loginID)
}

// FindUser mocks base method.
func (m *MockUserUsc) FindUser(ctx context.Context, id string) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, id)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserUscMockRecorder) FindUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserUsc)(nil).FindUser), ctx, id)
}

// LoginWithPassword mocks base method.
func (m *MockUserUsc) LoginWithPassword(ctx context.Context, loginID, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginWithPassword", ctx, loginID, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginWithPassword indicates an expected call of LoginWithPassword.
func (mr *MockUserUscMockRecorder) LoginWithPassword(ctx, loginID, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWithPassword", reflect.TypeOf((*MockUserUsc)(nil).LoginWithPassword), ctx, loginID, password)
}

// ModifyUser mocks base method.
func (m *MockUserUsc) ModifyUser(ctx context.Context, userNew dto.User) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyUser", ctx, userNew)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyUser indicates an expected call of ModifyUser.
func (mr *MockUserUscMockRecorder) ModifyUser(ctx, userNew interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyUser", reflect.TypeOf((*MockUserUsc)(nil).ModifyUser), ctx, userNew)
}

// RegisterUser mocks base method.
func (m *MockUserUsc) RegisterUser(ctx context.Context, input dto.User) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, input)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserUscMockRecorder) RegisterUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserUsc)(nil).RegisterUser), ctx, input)
}

// RemoveUser mocks base method.
func (m *MockUserUsc) RemoveUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUserUscMockRecorder) RemoveUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUserUsc)(nil).RemoveUser), ctx, id)
}

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticator) Authenticate(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticatorMockRecorder) Authenticate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticator)(nil).Authenticate), ctx)
}
