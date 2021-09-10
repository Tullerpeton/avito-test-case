package balance

import "net/http"

type Handler interface {
	GetUserBalance(w http.ResponseWriter, r *http.Request)
	WithdrawUserBalance(w http.ResponseWriter, r *http.Request)
	ImproveUserBalance(w http.ResponseWriter, r *http.Request)
	MakeUserBalanceTransfer(w http.ResponseWriter, r *http.Request)
}
