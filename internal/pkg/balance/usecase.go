package balance

import "github.com/avito-test-case/internal/pkg/models"

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock github.com/avito-test-case/internal/pkg/balance UseCase

type UseCase interface {
	GetUserBalance(userId uint64, currency string) (*models.UserBalance, error)
	WithdrawUserBalance(withdrawBalance *models.WithdrawBalance) (*models.UserBalance, error)
	ImproveUserBalance(improveBalance *models.ImproveBalance) (*models.UserBalance, error)
	MakeUserBalanceTransfer(transfer *models.Transfer) (*models.TransferResult, error)
}
