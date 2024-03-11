// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/beebeewijaya-tech/go-budget/internal/domain (interfaces: ExpenseUsecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "github.com/beebeewijaya-tech/go-budget/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockExpenseUsecase is a mock of ExpenseUsecase interface.
type MockExpenseUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockExpenseUsecaseMockRecorder
}

// MockExpenseUsecaseMockRecorder is the mock recorder for MockExpenseUsecase.
type MockExpenseUsecaseMockRecorder struct {
	mock *MockExpenseUsecase
}

// NewMockExpenseUsecase creates a new mock instance.
func NewMockExpenseUsecase(ctrl *gomock.Controller) *MockExpenseUsecase {
	mock := &MockExpenseUsecase{ctrl: ctrl}
	mock.recorder = &MockExpenseUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpenseUsecase) EXPECT() *MockExpenseUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockExpenseUsecase) Create(arg0 context.Context, arg1 domain.Expense) (domain.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(domain.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockExpenseUsecaseMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockExpenseUsecase)(nil).Create), arg0, arg1)
}

// DeleteByID mocks base method.
func (m *MockExpenseUsecase) DeleteByID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockExpenseUsecaseMockRecorder) DeleteByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockExpenseUsecase)(nil).DeleteByID), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockExpenseUsecase) GetByID(arg0 context.Context, arg1 string) (domain.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(domain.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockExpenseUsecaseMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockExpenseUsecase)(nil).GetByID), arg0, arg1)
}

// List mocks base method.
func (m *MockExpenseUsecase) List(arg0 context.Context, arg1 string, arg2, arg3 int) ([]domain.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]domain.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockExpenseUsecaseMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockExpenseUsecase)(nil).List), arg0, arg1, arg2, arg3)
}
