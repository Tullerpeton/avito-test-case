// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/avito-test-case/internal/pkg/balance (interfaces: UseCase)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/avito-test-case/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// GetUserBalance mocks base method.
func (m *MockUseCase) GetUserBalance(arg0 uint64, arg1 string) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBalance", arg0, arg1)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBalance indicates an expected call of GetUserBalance.
func (mr *MockUseCaseMockRecorder) GetUserBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBalance", reflect.TypeOf((*MockUseCase)(nil).GetUserBalance), arg0, arg1)
}

// ImproveUserBalance mocks base method.
func (m *MockUseCase) ImproveUserBalance(arg0 *models.ImproveBalance) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImproveUserBalance", arg0)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImproveUserBalance indicates an expected call of ImproveUserBalance.
func (mr *MockUseCaseMockRecorder) ImproveUserBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImproveUserBalance", reflect.TypeOf((*MockUseCase)(nil).ImproveUserBalance), arg0)
}

// MakeUserBalanceTransfer mocks base method.
func (m *MockUseCase) MakeUserBalanceTransfer(arg0 *models.Transfer) (*models.TransferResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeUserBalanceTransfer", arg0)
	ret0, _ := ret[0].(*models.TransferResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeUserBalanceTransfer indicates an expected call of MakeUserBalanceTransfer.
func (mr *MockUseCaseMockRecorder) MakeUserBalanceTransfer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeUserBalanceTransfer", reflect.TypeOf((*MockUseCase)(nil).MakeUserBalanceTransfer), arg0)
}

// WithdrawUserBalance mocks base method.
func (m *MockUseCase) WithdrawUserBalance(arg0 *models.WithdrawBalance) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawUserBalance", arg0)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawUserBalance indicates an expected call of WithdrawUserBalance.
func (mr *MockUseCaseMockRecorder) WithdrawUserBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawUserBalance", reflect.TypeOf((*MockUseCase)(nil).WithdrawUserBalance), arg0)
}