package transaction

import "github.com/avito-test-case/internal/pkg/models"

type Repository interface {
	SelectAllTransactions(userId uint64, paginator *models.PaginatorTransaction) ([]*models.Transaction, error)
}
