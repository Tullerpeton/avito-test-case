package transaction

import "net/http"

type Handler interface {
	GetUserTransactions(w http.ResponseWriter, r *http.Request)
}
