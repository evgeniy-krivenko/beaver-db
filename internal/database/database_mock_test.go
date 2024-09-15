// Code generated by MockGen. DO NOT EDIT.
// Source: database.go
//
// Generated by this command:
//
//	mockgen -source=database.go -destination=database_mock_test.go -package=database_test
//

// Package database_test is a generated GoMock package.
package database_test

import (
	compute "beaver/internal/database/compute"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockcomputeParser is a mock of computeParser interface.
type MockcomputeParser struct {
	ctrl     *gomock.Controller
	recorder *MockcomputeParserMockRecorder
}

// MockcomputeParserMockRecorder is the mock recorder for MockcomputeParser.
type MockcomputeParserMockRecorder struct {
	mock *MockcomputeParser
}

// NewMockcomputeParser creates a new mock instance.
func NewMockcomputeParser(ctrl *gomock.Controller) *MockcomputeParser {
	mock := &MockcomputeParser{ctrl: ctrl}
	mock.recorder = &MockcomputeParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcomputeParser) EXPECT() *MockcomputeParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockcomputeParser) Parse(arg0 string) (compute.Query, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", arg0)
	ret0, _ := ret[0].(compute.Query)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockcomputeParserMockRecorder) Parse(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockcomputeParser)(nil).Parse), arg0)
}

// Mockstorage is a mock of storage interface.
type Mockstorage struct {
	ctrl     *gomock.Controller
	recorder *MockstorageMockRecorder
}

// MockstorageMockRecorder is the mock recorder for Mockstorage.
type MockstorageMockRecorder struct {
	mock *Mockstorage
}

// NewMockstorage creates a new mock instance.
func NewMockstorage(ctrl *gomock.Controller) *Mockstorage {
	mock := &Mockstorage{ctrl: ctrl}
	mock.recorder = &MockstorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockstorage) EXPECT() *MockstorageMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *Mockstorage) Del(cxt context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", cxt, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockstorageMockRecorder) Del(cxt, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*Mockstorage)(nil).Del), cxt, key)
}

// Get mocks base method.
func (m *Mockstorage) Get(cxt context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", cxt, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockstorageMockRecorder) Get(cxt, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mockstorage)(nil).Get), cxt, key)
}

// Set mocks base method.
func (m *Mockstorage) Set(cxt context.Context, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", cxt, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockstorageMockRecorder) Set(cxt, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*Mockstorage)(nil).Set), cxt, key, value)
}
