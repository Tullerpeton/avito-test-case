package transaction

import "github.com/avito-test-case/internal/pkg/models"

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock github.com/avito-test-case/internal/pkg/transaction UseCase

type UseCase interface {
	GetUserTransactions(userId uint64, paginator *models.PaginatorTransaction) ([]*models.Transaction, error)
}
