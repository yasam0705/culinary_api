// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "github/culinary_api/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m_2 *MockAuth) Login(ctx context.Context, m *entity.User) (*entity.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Login", ctx, m)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), ctx, m)
}

// Registration mocks base method.
func (m_2 *MockAuth) Registration(ctx context.Context, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Registration", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Registration indicates an expected call of Registration.
func (mr *MockAuthMockRecorder) Registration(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockAuth)(nil).Registration), ctx, m)
}
