package transaction

import "github.com/avito-test-case/internal/pkg/models"

type UseCase interface {
	GetUserTransactions(userId uint64, paginator *models.PaginatorTransaction) ([]*models.Transaction, error)
}
