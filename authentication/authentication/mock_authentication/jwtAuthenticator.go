// Code generated by MockGen. DO NOT EDIT.
// Source: jwtAuthenticator.go

// Package mock_authentication is a generated GoMock package.
package mock_authentication

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHashStore is a mock of HashStore interface
type MockHashStore struct {
	ctrl     *gomock.Controller
	recorder *MockHashStoreMockRecorder
}

// MockHashStoreMockRecorder is the mock recorder for MockHashStore
type MockHashStoreMockRecorder struct {
	mock *MockHashStore
}

// NewMockHashStore creates a new mock instance
func NewMockHashStore(ctrl *gomock.Controller) *MockHashStore {
	mock := &MockHashStore{ctrl: ctrl}
	mock.recorder = &MockHashStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHashStore) EXPECT() *MockHashStoreMockRecorder {
	return m.recorder
}

// GetHash mocks base method
func (m *MockHashStore) GetHash(i int64) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetHash", i)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHash indicates an expected call of GetHash
func (mr *MockHashStoreMockRecorder) GetHash(i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*MockHashStore)(nil).GetHash), i)
}

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetUserId mocks base method
func (m *MockUserService) GetUserId(email string) (int64, error) {
	ret := m.ctrl.Call(m, "GetUserId", email)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId
func (mr *MockUserServiceMockRecorder) GetUserId(email interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockUserService)(nil).GetUserId), email)
}
