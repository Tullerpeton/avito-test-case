package balance

import "github.com/avito-test-case/internal/pkg/models"

//go:generate mockgen -destination=./mock/mock_repository.go -package=mock github.com/avito-test-case/internal/pkg/balance Repository

type Repository interface {
	SelectUserBalanceById(userId uint64) (*models.UserBalance, error)
	ImproveUserBalance(userId uint64, value float64) (*models.UserBalance, error)
	DoUserBalanceTransfer(transfer *models.Transfer) (*models.TransferResult, error)
	WithdrawUserBalance(userId uint64, value float64) (*models.UserBalance, error)
}
