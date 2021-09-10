package usecase

import (
	"github.com/avito-test-case/internal/app/errors/repository_errors"
	"github.com/avito-test-case/internal/app/errors/usecase_errors"
	"github.com/avito-test-case/internal/pkg/balance"
	"github.com/avito-test-case/internal/pkg/models"
	"github.com/avito-test-case/pkg/tools/currency_converter"
)

type BalanceUseCase struct {
	balanceRepo balance.Repository
}

func NewUseCase(balanceRepo balance.Repository) balance.UseCase {
	return &BalanceUseCase{
		balanceRepo: balanceRepo,
	}
}

func (u *BalanceUseCase) GetUserBalance(userId uint64, currencyToConvert string) (*models.UserBalance, error) {
	userBalance, err := u.balanceRepo.SelectUserBalanceById(userId)
	if err != nil {
		return nil, usecase_errors.UcUserBalanceNotFound
	}

	if currencyToConvert != "" {
		userBalance.Balance, err = currency_converter.ConvertCurrency(userBalance.Balance, currencyToConvert)
		if err != nil {
			return nil, usecase_errors.UcCanNotConvertCurrency
		}
		userBalance.Currency = currencyToConvert
	}

	return userBalance, nil
}

func (u *BalanceUseCase) WithdrawUserBalance(withdrawBalance *models.WithdrawBalance) (*models.UserBalance, error) {
	updatedUserBalance, err := u.balanceRepo.WithdrawUserBalance(withdrawBalance.Id, withdrawBalance.Value)
	switch err {
	case nil:
		return updatedUserBalance, nil
	case repository_errors.PgCanNotUpdate:
		return nil, usecase_errors.UcInsufficientFunds
	case repository_errors.PgCanNotFind:
		return nil, usecase_errors.UcUserBalanceNotFound
	default:
		return nil, usecase_errors.UcInternalServerError
	}
}

func (u *BalanceUseCase) ImproveUserBalance(improveBalance *models.ImproveBalance) (*models.UserBalance, error) {
	updatedUserBalance, err := u.balanceRepo.ImproveUserBalance(improveBalance.Id, improveBalance.Value)
	switch err {
	case nil:
		return updatedUserBalance, nil
	case repository_errors.PgCanNotUpdate:
		return nil, usecase_errors.UcUserBalanceNotFound
	default:
		return nil, usecase_errors.UcInternalServerError
	}
}

func (u *BalanceUseCase) MakeUserBalanceTransfer(transfer *models.Transfer) (*models.TransferResult, error) {
	transferResult, err := u.balanceRepo.DoUserBalanceTransfer(transfer)
	switch err {
	case nil:
		return transferResult, nil
	case repository_errors.PgCanNotUpdate:
		return nil, usecase_errors.UcInsufficientFunds
	case repository_errors.PgCanNotFind:
		return nil, usecase_errors.UcUserBalanceNotFound
	default:
		return nil, usecase_errors.UcInternalServerError
	}
}
