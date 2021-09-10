package repository

import (
	"fmt"

	"github.com/avito-test-case/internal/app/errors/repository_errors"
	"github.com/avito-test-case/internal/pkg/balance"
	"github.com/avito-test-case/internal/pkg/models"

	"github.com/jmoiron/sqlx"
)

type PostgresqlRepository struct {
	db *sqlx.DB
}

func NewSessionPostgresqlRepository(db *sqlx.DB) balance.Repository {
	return &PostgresqlRepository{
		db: db,
	}
}

func (r *PostgresqlRepository) SelectUserBalanceById(userId uint64) (*models.UserBalance, error) {
	row := r.db.QueryRow(
		"SELECT user_id, balance "+
			"FROM balances "+
			"WHERE user_id = $1",
		userId,
	)

	userBalance := models.UserBalance{Currency: "RUB"}
	if err := row.Scan(
		&userBalance.Id,
		&userBalance.Balance,
	); err != nil {
		return nil, repository_errors.PgCanNotFind
	}

	return &userBalance, nil
}

func (r *PostgresqlRepository) ImproveUserBalance(userId uint64, value float64) (*models.UserBalance, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, repository_errors.PgInternalDbError
	}

	row := tx.QueryRow(
		"INSERT INTO balances AS b (user_id, balance)"+
			"VALUES ($1, $2) "+
			"ON CONFLICT (user_id) "+
			"DO UPDATE SET "+
			"balance = b.balance + $2 "+
			"WHERE b.user_id = $1 "+
			"RETURNING b.balance",
		userId,
		value,
	)

	userBalance := models.UserBalance{Id: userId, Currency: "RUB"}
	if err = row.Scan(
		&userBalance.Balance,
	); err != nil {
		tx.Rollback()
		return nil, repository_errors.PgCanNotUpdate
	}

	_, err = tx.Exec(
		"INSERT INTO transactions(type_operation, user_id, context, value) "+
			"SELECT id, $1, $2, $3 "+
			"FROM transaction_types "+
			"WHERE title = $4",
		userId,
		fmt.Sprintf("user [id=%d] single operation", userId),
		value,
		"improve",
	)

	if err != nil {
		tx.Rollback()
		return nil, repository_errors.PgInternalDbError
	}

	err = tx.Commit()
	if err != nil {
		return nil, repository_errors.PgInternalDbError
	}

	return &userBalance, nil
}

func (r *PostgresqlRepository) WithdrawUserBalance(userId uint64, value float64) (*models.UserBalance, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, repository_errors.PgInternalDbError
	}

	row := tx.QueryRow(
		"UPDATE balances AS b SET "+
			"balance = b.balance - $2 "+
			"WHERE b.user_id = $1 "+
			"RETURNING b.balance",
		userId,
		value,
	)

	userBalance := models.UserBalance{Id: userId, Currency: "RUB"}
	if err = row.Scan(
		&userBalance.Balance,
	); err != nil {
		var balance float64
		if err := tx.QueryRow(
			"SELECT balance "+
				"FROM balances "+
				"WHERE user_id = $1",
			userId,
		).Scan(&balance); err != nil {
			tx.Rollback()
			return nil, repository_errors.PgCanNotFind
		}

		tx.Rollback()
		return nil, repository_errors.PgCanNotUpdate
	}

	_, err = tx.Exec(
		"INSERT INTO transactions(type_operation, user_id, context, value) "+
			"SELECT id, $1, $2, $3 "+
			"FROM transaction_types "+
			"WHERE title = $4",
		userId,
		fmt.Sprintf("user [id=%d] single operation", userId),
		value,
		"withdraw",
	)

	if err != nil {
		tx.Rollback()
		return nil, repository_errors.PgInternalDbError
	}

	err = tx.Commit()
	if err != nil {
		return nil, repository_errors.PgInternalDbError
	}

	return &userBalance, nil
}

func (r *PostgresqlRepository) DoUserBalanceTransfer(transfer *models.Transfer) (*models.TransferResult, error) {
	transferResult := models.TransferResult{
		Sender:   models.UserBalance{Currency: "RUB"},
		Receiver: models.UserBalance{Currency: "RUB"},
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, repository_errors.PgInternalDbError
	}

	err = tx.QueryRow(
		"UPDATE balances SET "+
			"balance = balance - $2 "+
			"WHERE user_id = $1 "+
			"RETURNING balance",
		transfer.SenderId,
		transfer.Value,
	).Scan(transferResult.Sender.Balance)

	if err != nil {
		var balance float64
		if err := tx.QueryRow(
			"SELECT balance "+
				"FROM balances "+
				"WHERE user_id = $1",
			transfer.SenderId,
		).Scan(&balance); err != nil {
			tx.Rollback()
			return nil, repository_errors.PgCanNotFind
		}

		tx.Rollback()
		return nil, repository_errors.PgCanNotUpdate
	}

	err = tx.QueryRow(
		"INSERT INTO balances (user_id, balance)"+
			"VALUES ($1, $2) "+
			"ON CONFLICT (user_id) "+
			"UPDATE balances SET "+
			"balance = balance + $2 "+
			"WHERE user_id = $1 "+
			"RETURNING balance",
		transfer.ReceiverId,
		transfer.Value,
	).Scan(transferResult.Receiver.Balance)

	if err != nil {
		var balance float64
		if err := tx.QueryRow(
			"SELECT balance "+
				"FROM balances "+
				"WHERE user_id = $1",
			transfer.SenderId,
		).Scan(&balance); err != nil {
			tx.Rollback()
			return nil, repository_errors.PgCanNotFind
		}

		tx.Rollback()
		return nil, repository_errors.PgCanNotUpdate
	}

	typeOperation := "transfer"
	_, err = tx.Exec(
		"INSERT INTO transactions(type_operation, user_id, context, value) "+
			"SELECT name, $1, $2, $3 "+
			"FROM transaction_types "+
			"WHERE title = $4",
		transfer.ReceiverId,
		fmt.Sprintf("user [id=%d] received money from user [id=%d]", transfer.ReceiverId, transfer.SenderId),
		transfer.Value,
		typeOperation,
	)

	if err != nil {
		tx.Rollback()
		return nil, repository_errors.PgInternalDbError
	}

	_, err = tx.Exec(
		"INSERT INTO transactions(type_operation, user_id, context, value) "+
			"SELECT name, $1, $2, $3 "+
			"FROM transaction_types "+
			"WHERE title = $4",
		transfer.SenderId,
		fmt.Sprintf("user [id=%d] sent money to user [id=%d]", transfer.SenderId, transfer.ReceiverId),
		transfer.Value,
		typeOperation,
	)

	if err != nil {
		tx.Rollback()
		return nil, repository_errors.PgInternalDbError
	}

	err = tx.Commit()
	if err != nil {
		return nil, repository_errors.PgInternalDbError
	}

	return &transferResult, nil
}
