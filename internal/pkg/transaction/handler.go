package transaction

import "net/http"

//go:generate mockgen -destination=./mock/mock_handler.go -package=mock github.com/avito-test-case/internal/pkg/transaction Handler

type Handler interface {
	GetUserTransactions(w http.ResponseWriter, r *http.Request)
}
