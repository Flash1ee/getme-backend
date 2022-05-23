// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package mock is a generated GoMock package.
package mock

import (
	dto "getme-backend/internal/app/user/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockauthChecker is a mock of authChecker interface.
type MockauthChecker struct {
	ctrl     *gomock.Controller
	recorder *MockauthCheckerMockRecorder
}

// MockauthCheckerMockRecorder is the mock recorder for MockauthChecker.
type MockauthCheckerMockRecorder struct {
	mock *MockauthChecker
}

// NewMockauthChecker creates a new mock instance.
func NewMockauthChecker(ctrl *gomock.Controller) *MockauthChecker {
	mock := &MockauthChecker{ctrl: ctrl}
	mock.recorder = &MockauthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockauthChecker) EXPECT() *MockauthCheckerMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockauthChecker) Check(data *dto.UserAuthUsecase) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthSimple", data)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockauthCheckerMockRecorder) Check(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthSimple", reflect.TypeOf((*MockauthChecker)(nil).Check), data)
}
