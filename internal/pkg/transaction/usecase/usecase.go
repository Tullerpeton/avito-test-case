package usecase

import (
	"github.com/avito-test-case/internal/app/errors/repository_errors"
	"github.com/avito-test-case/internal/app/errors/usecase_errors"
	"github.com/avito-test-case/internal/pkg/models"
	"github.com/avito-test-case/internal/pkg/transaction"
)

type TransactionUseCase struct {
	transactionRepo transaction.Repository
}

func NewUseCase(transactionRepo transaction.Repository) transaction.UseCase {
	return &TransactionUseCase{
		transactionRepo: transactionRepo,
	}
}

func (u *TransactionUseCase) GetUserTransactions(userId uint64, paginator *models.PaginatorTransaction) ([]*models.Transaction, error) {
	transactions, err := u.transactionRepo.SelectAllTransactions(userId, paginator)
	switch err {
	case nil:
		return transactions, nil
	case repository_errors.PgCanNotFind:
		return nil, usecase_errors.UcUserBalanceNotFound
	default:
		return nil, usecase_errors.UcInternalServerError
	}
}
