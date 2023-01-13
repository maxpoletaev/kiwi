// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/maxpoletaev/kiwi/storage"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBackend) Get(key string) ([]storage.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]storage.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBackendMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBackend)(nil).Get), key)
}

// Put mocks base method.
func (m *MockBackend) Put(key string, value storage.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockBackendMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockBackend)(nil).Put), key, value)
}

// MockScanIterator is a mock of ScanIterator interface.
type MockScanIterator struct {
	ctrl     *gomock.Controller
	recorder *MockScanIteratorMockRecorder
}

// MockScanIteratorMockRecorder is the mock recorder for MockScanIterator.
type MockScanIteratorMockRecorder struct {
	mock *MockScanIterator
}

// NewMockScanIterator creates a new mock instance.
func NewMockScanIterator(ctrl *gomock.Controller) *MockScanIterator {
	mock := &MockScanIterator{ctrl: ctrl}
	mock.recorder = &MockScanIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanIterator) EXPECT() *MockScanIteratorMockRecorder {
	return m.recorder
}

// HasNext mocks base method.
func (m *MockScanIterator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockScanIteratorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockScanIterator)(nil).HasNext))
}

// Next mocks base method.
func (m *MockScanIterator) Next() (string, storage.Value) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(storage.Value)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockScanIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockScanIterator)(nil).Next))
}
