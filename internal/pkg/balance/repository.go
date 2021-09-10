package balance

import "github.com/avito-test-case/internal/pkg/models"

type Repository interface {
	SelectUserBalanceById(userId uint64) (*models.UserBalance, error)
	ImproveUserBalance(userId uint64, value float64) (*models.UserBalance, error)
	DoUserBalanceTransfer(transfer *models.Transfer) (*models.TransferResult, error)
	WithdrawUserBalance(userId uint64, value float64) (*models.UserBalance, error)
}
