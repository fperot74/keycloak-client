// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-client/v2/toolbox (interfaces: Logger)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=./mock/logger.go -package=mock -mock_names=Logger=Logger github.com/cloudtrust/keycloak-client/v2/toolbox Logger
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// Logger is a mock of Logger interface.
type Logger struct {
	ctrl     *gomock.Controller
	recorder *LoggerMockRecorder
	isgomock struct{}
}

// LoggerMockRecorder is the mock recorder for Logger.
type LoggerMockRecorder struct {
	mock *Logger
}

// NewLogger creates a new mock instance.
func NewLogger(ctrl *gomock.Controller) *Logger {
	mock := &Logger{ctrl: ctrl}
	mock.recorder = &LoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Logger) EXPECT() *LoggerMockRecorder {
	return m.recorder
}

// Warn mocks base method.
func (m *Logger) Warn(ctx context.Context, keyvals ...any) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range keyvals {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *LoggerMockRecorder) Warn(ctx any, keyvals ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, keyvals...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*Logger)(nil).Warn), varargs...)
}
