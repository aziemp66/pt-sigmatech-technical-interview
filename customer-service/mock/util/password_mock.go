// Code generated by MockGen. DO NOT EDIT.
// Source: util/password/type.go
//
// Generated by this command:
//
//	mockgen -package=mock_util -source=util/password/type.go -destination=mock/util/password_mock.go -typed=true
//

// Package mock_util is a generated GoMock package.
package mock_util

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPasswordManager is a mock of PasswordManager interface.
type MockPasswordManager struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordManagerMockRecorder
}

// MockPasswordManagerMockRecorder is the mock recorder for MockPasswordManager.
type MockPasswordManagerMockRecorder struct {
	mock *MockPasswordManager
}

// NewMockPasswordManager creates a new mock instance.
func NewMockPasswordManager(ctrl *gomock.Controller) *MockPasswordManager {
	mock := &MockPasswordManager{ctrl: ctrl}
	mock.recorder = &MockPasswordManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordManager) EXPECT() *MockPasswordManagerMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method.
func (m *MockPasswordManager) CheckPasswordHash(password, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", password, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash.
func (mr *MockPasswordManagerMockRecorder) CheckPasswordHash(password, hash any) *MockPasswordManagerCheckPasswordHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockPasswordManager)(nil).CheckPasswordHash), password, hash)
	return &MockPasswordManagerCheckPasswordHashCall{Call: call}
}

// MockPasswordManagerCheckPasswordHashCall wrap *gomock.Call
type MockPasswordManagerCheckPasswordHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPasswordManagerCheckPasswordHashCall) Return(arg0 error) *MockPasswordManagerCheckPasswordHashCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPasswordManagerCheckPasswordHashCall) Do(f func(string, string) error) *MockPasswordManagerCheckPasswordHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPasswordManagerCheckPasswordHashCall) DoAndReturn(f func(string, string) error) *MockPasswordManagerCheckPasswordHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HashPassword mocks base method.
func (m *MockPasswordManager) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockPasswordManagerMockRecorder) HashPassword(password any) *MockPasswordManagerHashPasswordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockPasswordManager)(nil).HashPassword), password)
	return &MockPasswordManagerHashPasswordCall{Call: call}
}

// MockPasswordManagerHashPasswordCall wrap *gomock.Call
type MockPasswordManagerHashPasswordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPasswordManagerHashPasswordCall) Return(arg0 string, arg1 error) *MockPasswordManagerHashPasswordCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPasswordManagerHashPasswordCall) Do(f func(string) (string, error)) *MockPasswordManagerHashPasswordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPasswordManagerHashPasswordCall) DoAndReturn(f func(string) (string, error)) *MockPasswordManagerHashPasswordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PasswordValidation mocks base method.
func (m *MockPasswordManager) PasswordValidation(password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordValidation", password)
	ret0, _ := ret[0].(error)
	return ret0
}

// PasswordValidation indicates an expected call of PasswordValidation.
func (mr *MockPasswordManagerMockRecorder) PasswordValidation(password any) *MockPasswordManagerPasswordValidationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordValidation", reflect.TypeOf((*MockPasswordManager)(nil).PasswordValidation), password)
	return &MockPasswordManagerPasswordValidationCall{Call: call}
}

// MockPasswordManagerPasswordValidationCall wrap *gomock.Call
type MockPasswordManagerPasswordValidationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPasswordManagerPasswordValidationCall) Return(arg0 error) *MockPasswordManagerPasswordValidationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPasswordManagerPasswordValidationCall) Do(f func(string) error) *MockPasswordManagerPasswordValidationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPasswordManagerPasswordValidationCall) DoAndReturn(f func(string) error) *MockPasswordManagerPasswordValidationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
