// Code generated by MockGen. DO NOT EDIT.
// Source: session/session.go

// Package session is a generated GoMock package.
package session

import (
	gomock "github.com/golang/mock/gomock"
	sessions "github.com/gorilla/sessions"
	echo "github.com/labstack/echo/v4"
	reflect "reflect"
)

// MockSession is a mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockSession) Store() sessions.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store")
	ret0, _ := ret[0].(sessions.Store)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockSessionMockRecorder) Store() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockSession)(nil).Store))
}

// RevokeSession mocks base method
func (m *MockSession) RevokeSession(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSession", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeSession indicates an expected call of RevokeSession
func (mr *MockSessionMockRecorder) RevokeSession(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSession", reflect.TypeOf((*MockSession)(nil).RevokeSession), c)
}
