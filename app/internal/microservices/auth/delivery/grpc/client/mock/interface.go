// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "getme-backend/internal/microservices/auth/sessions/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthCheckerClient is a mock of AuthCheckerClient interface.
type MockAuthCheckerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCheckerClientMockRecorder
}

// MockAuthCheckerClientMockRecorder is the mock recorder for MockAuthCheckerClient.
type MockAuthCheckerClientMockRecorder struct {
	mock *MockAuthCheckerClient
}

// NewMockAuthCheckerClient creates a new mock instance.
func NewMockAuthCheckerClient(ctrl *gomock.Controller) *MockAuthCheckerClient {
	mock := &MockAuthCheckerClient{ctrl: ctrl}
	mock.recorder = &MockAuthCheckerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCheckerClient) EXPECT() *MockAuthCheckerClientMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockAuthCheckerClient) Check(ctx context.Context, sessionID string) (models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx, sessionID)
	ret0, _ := ret[0].(models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockAuthCheckerClientMockRecorder) Check(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAuthCheckerClient)(nil).Check), ctx, sessionID)
}

// CheckWithDelete mocks base method.
func (m *MockAuthCheckerClient) CheckWithDelete(ctx context.Context, tokenID string) (models.ResultByToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckWithDelete", ctx, tokenID)
	ret0, _ := ret[0].(models.ResultByToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckWithDelete indicates an expected call of CheckWithDelete.
func (mr *MockAuthCheckerClientMockRecorder) CheckWithDelete(ctx, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckWithDelete", reflect.TypeOf((*MockAuthCheckerClient)(nil).CheckWithDelete), ctx, tokenID)
}

// Create mocks base method.
func (m *MockAuthCheckerClient) Create(ctx context.Context, userID int64) (models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID)
	ret0, _ := ret[0].(models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthCheckerClientMockRecorder) Create(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthCheckerClient)(nil).Create), ctx, userID)
}

// CreateByToken mocks base method.
func (m *MockAuthCheckerClient) CreateByToken(ctx context.Context, tokenID string, userID int64) (models.ResultByToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateByToken", ctx, tokenID, userID)
	ret0, _ := ret[0].(models.ResultByToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateByToken indicates an expected call of CreateByToken.
func (mr *MockAuthCheckerClientMockRecorder) CreateByToken(ctx, tokenID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateByToken", reflect.TypeOf((*MockAuthCheckerClient)(nil).CreateByToken), ctx, tokenID, userID)
}

// Delete mocks base method.
func (m *MockAuthCheckerClient) Delete(ctx context.Context, sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthCheckerClientMockRecorder) Delete(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthCheckerClient)(nil).Delete), ctx, sessionID)
}
