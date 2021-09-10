package repository

import (
	"fmt"

	"github.com/avito-test-case/internal/app/errors/repository_errors"
	"github.com/avito-test-case/internal/pkg/models"
	"github.com/avito-test-case/internal/pkg/transaction"

	"github.com/jmoiron/sqlx"
)

type PostgresqlRepository struct {
	db *sqlx.DB
}

func NewSessionPostgresqlRepository(db *sqlx.DB) transaction.Repository {
	return &PostgresqlRepository{
		db: db,
	}
}

// Create sort string from paginator options
func (r *PostgresqlRepository) createSortString(sortField, sortDirection string) (string, error) {
	// Select order target
	var orderTarget string
	switch sortField {
	case models.TransactionSumSort:
		orderTarget = "value"
	case models.TransactionDateAddedSort:
		orderTarget = "created_at"
	default:
		return "", repository_errors.PgIncorrectPaginator
	}

	// Select order direction
	var orderDirection string
	switch sortDirection {
	case models.PaginatorASC:
		orderDirection = "ASC"
	case models.PaginatorDESC:
		orderDirection = "DESC"
	default:
		return "", repository_errors.PgIncorrectPaginator
	}

	return fmt.Sprintf("ORDER BY %s %s ", orderTarget, orderDirection), nil
}

func (r *PostgresqlRepository) SelectAllTransactions(userId uint64, paginator *models.PaginatorTransaction) ([]*models.Transaction, error) {
	sortString, err := r.createSortString(paginator.SortField, paginator.SortDirection)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(
		"SELECT title, created_at, context, value "+
			"FROM transactions t "+
			"JOIN transaction_types tt ON t.type_operation = tt.id "+
			"WHERE user_id = $3 "+
			sortString+
			"LIMIT $1 OFFSET $2",
		paginator.CountRecords,
		paginator.CountRecords*(paginator.Page-1),
		userId,
	)

	if err != nil {
		return nil, repository_errors.PgCanNotFind
	}
	defer rows.Close()

	transactions := make([]*models.Transaction, 0)
	for rows.Next() {
		transaction := &models.Transaction{}
		err = rows.Scan(
			&transaction.TypeOperation,
			&transaction.CreatedAt,
			&transaction.Context,
			&transaction.Value,
		)

		if err != nil {
			return nil, repository_errors.PgCanNotFind
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
