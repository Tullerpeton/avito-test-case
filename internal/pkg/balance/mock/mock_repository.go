// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/avito-test-case/internal/pkg/balance (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/avito-test-case/internal/pkg/models"
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

// DoUserBalanceTransfer mocks base method.
func (m *MockRepository) DoUserBalanceTransfer(arg0 *models.Transfer) (*models.TransferResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoUserBalanceTransfer", arg0)
	ret0, _ := ret[0].(*models.TransferResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoUserBalanceTransfer indicates an expected call of DoUserBalanceTransfer.
func (mr *MockRepositoryMockRecorder) DoUserBalanceTransfer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoUserBalanceTransfer", reflect.TypeOf((*MockRepository)(nil).DoUserBalanceTransfer), arg0)
}

// ImproveUserBalance mocks base method.
func (m *MockRepository) ImproveUserBalance(arg0 uint64, arg1 float64) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImproveUserBalance", arg0, arg1)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImproveUserBalance indicates an expected call of ImproveUserBalance.
func (mr *MockRepositoryMockRecorder) ImproveUserBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImproveUserBalance", reflect.TypeOf((*MockRepository)(nil).ImproveUserBalance), arg0, arg1)
}

// SelectUserBalanceById mocks base method.
func (m *MockRepository) SelectUserBalanceById(arg0 uint64) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserBalanceById", arg0)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserBalanceById indicates an expected call of SelectUserBalanceById.
func (mr *MockRepositoryMockRecorder) SelectUserBalanceById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserBalanceById", reflect.TypeOf((*MockRepository)(nil).SelectUserBalanceById), arg0)
}

// WithdrawUserBalance mocks base method.
func (m *MockRepository) WithdrawUserBalance(arg0 uint64, arg1 float64) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawUserBalance", arg0, arg1)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawUserBalance indicates an expected call of WithdrawUserBalance.
func (mr *MockRepositoryMockRecorder) WithdrawUserBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawUserBalance", reflect.TypeOf((*MockRepository)(nil).WithdrawUserBalance), arg0, arg1)
}