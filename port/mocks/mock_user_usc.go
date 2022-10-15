// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_usc.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/w-woong/user/dto"
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

// FindUserByID mocks base method.
func (m *MockUserUsc) FindUserByID(ID string) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ID)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserUscMockRecorder) FindUserByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserUsc)(nil).FindUserByID), ID)
}

// ModifyUser mocks base method.
func (m *MockUserUsc) ModifyUser(ID string, input dto.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyUser", ID, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyUser indicates an expected call of ModifyUser.
func (mr *MockUserUscMockRecorder) ModifyUser(ID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyUser", reflect.TypeOf((*MockUserUsc)(nil).ModifyUser), ID, input)
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
func (m *MockUserUsc) RemoveUser(ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUserUscMockRecorder) RemoveUser(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUserUsc)(nil).RemoveUser), ID)
}