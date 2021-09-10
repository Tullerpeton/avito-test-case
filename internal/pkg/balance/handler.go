package balance

import "net/http"

//go:generate mockgen -destination=./mock/mock_handler.go -package=mock github.com/avito-test-case/internal/pkg/balance Handler

type Handler interface {
	GetUserBalance(w http.ResponseWriter, r *http.Request)
	WithdrawUserBalance(w http.ResponseWriter, r *http.Request)
	ImproveUserBalance(w http.ResponseWriter, r *http.Request)
	MakeUserBalanceTransfer(w http.ResponseWriter, r *http.Request)
}
