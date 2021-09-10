package balance

import "github.com/avito-test-case/internal/pkg/models"

type UseCase interface {
	GetUserBalance(userId uint64, currency string) (*models.UserBalance, error)
	WithdrawUserBalance(withdrawBalance *models.WithdrawBalance) (*models.UserBalance, error)
	ImproveUserBalance(improveBalance *models.WithdrawBalance) (*models.UserBalance, error)
	MakeUserBalanceTransfer(transfer *models.Transfer) (*models.TransferResult, error)
}
