package transaction

import "github.com/avito-test-case/internal/pkg/models"

//go:generate mockgen -destination=./mock/mock_repository.go -package=mock github.com/avito-test-case/internal/pkg/transaction Repository

type Repository interface {
	SelectAllTransactions(userId uint64, paginator *models.PaginatorTransaction) ([]*models.Transaction, error)
}
