// Code generated by MockGen. DO NOT EDIT.
// Source: openai.go

// Package mock_card is a generated GoMock package.
package mock_card

import (
	reflect "reflect"

	entity "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockInteractor is a mock of Interactor interface.
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor.
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance.
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CreateCard mocks base method.
func (m *MockInteractor) CreateCard(card entity.Card) (entity.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", card)
	ret0, _ := ret[0].(entity.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCard indicates an expected call of CreateCard.
func (mr *MockInteractorMockRecorder) CreateCard(card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockInteractor)(nil).CreateCard), card)
}

// CreateCards mocks base method.
func (m *MockInteractor) CreateCards(cards []entity.Card) ([]entity.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCards", cards)
	ret0, _ := ret[0].([]entity.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCards indicates an expected call of CreateCards.
func (mr *MockInteractorMockRecorder) CreateCards(cards interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCards", reflect.TypeOf((*MockInteractor)(nil).CreateCards), cards)
}

// DeleteCard mocks base method.
func (m *MockInteractor) DeleteCard(id, sub string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", id, sub)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockInteractorMockRecorder) DeleteCard(id, sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockInteractor)(nil).DeleteCard), id, sub)
}

// ListCards mocks base method.
func (m *MockInteractor) ListCards(sub string) ([]entity.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCards", sub)
	ret0, _ := ret[0].([]entity.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCards indicates an expected call of ListCards.
func (mr *MockInteractorMockRecorder) ListCards(sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCards", reflect.TypeOf((*MockInteractor)(nil).ListCards), sub)
}

// UpdateCard mocks base method.
func (m *MockInteractor) UpdateCard(card entity.Card, id, sub string) (entity.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", card, id, sub)
	ret0, _ := ret[0].(entity.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockInteractorMockRecorder) UpdateCard(card, id, sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockInteractor)(nil).UpdateCard), card, id, sub)
}
