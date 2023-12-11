// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	reflect "reflect"

	entity "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DeleteBySub mocks base method.
func (m *MockRepository) DeleteBySub(sub string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBySub", sub)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBySub indicates an expected call of DeleteBySub.
func (mr *MockRepositoryMockRecorder) DeleteBySub(sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBySub", reflect.TypeOf((*MockRepository)(nil).DeleteBySub), sub)
}

// ExistsByEmail mocks base method.
func (m *MockRepository) ExistsByEmail(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByEmail", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByEmail indicates an expected call of ExistsByEmail.
func (mr *MockRepositoryMockRecorder) ExistsByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByEmail", reflect.TypeOf((*MockRepository)(nil).ExistsByEmail), email)
}

// ExistsByUserID mocks base method.
func (m *MockRepository) ExistsByUserID(userID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByUserID", userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByUserID indicates an expected call of ExistsByUserID.
func (mr *MockRepositoryMockRecorder) ExistsByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByUserID", reflect.TypeOf((*MockRepository)(nil).ExistsByUserID), userID)
}

// Insert mocks base method.
func (m *MockRepository) Insert(user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepositoryMockRecorder) Insert(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), user)
}

// UpdateBySub mocks base method.
func (m *MockRepository) UpdateBySub(user *entity.User, sub string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBySub", user, sub)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBySub indicates an expected call of UpdateBySub.
func (mr *MockRepositoryMockRecorder) UpdateBySub(user, sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBySub", reflect.TypeOf((*MockRepository)(nil).UpdateBySub), user, sub)
}
