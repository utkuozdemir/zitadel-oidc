// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/oidc/pkg/op (interfaces: Signer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	jose "gopkg.in/square/go-jose.v2"
	reflect "reflect"
)

// MockSigner is a mock of Signer interface
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// Health mocks base method
func (m *MockSigner) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health
func (mr *MockSignerMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockSigner)(nil).Health), arg0)
}

// SignatureAlgorithm mocks base method
func (m *MockSigner) SignatureAlgorithm() jose.SignatureAlgorithm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignatureAlgorithm")
	ret0, _ := ret[0].(jose.SignatureAlgorithm)
	return ret0
}

// SignatureAlgorithm indicates an expected call of SignatureAlgorithm
func (mr *MockSignerMockRecorder) SignatureAlgorithm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignatureAlgorithm", reflect.TypeOf((*MockSigner)(nil).SignatureAlgorithm))
}

// Signer mocks base method
func (m *MockSigner) Signer() jose.Signer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signer")
	ret0, _ := ret[0].(jose.Signer)
	return ret0
}

// Signer indicates an expected call of Signer
func (mr *MockSignerMockRecorder) Signer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signer", reflect.TypeOf((*MockSigner)(nil).Signer))
}
