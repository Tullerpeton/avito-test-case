package handler

import (
	"net/http"
	"strconv"

	"github.com/avito-test-case/internal/app/errors/handler_errors"
	"github.com/avito-test-case/internal/app/errors/usecase_errors"
	"github.com/avito-test-case/internal/pkg/models"
	"github.com/avito-test-case/internal/pkg/transaction"
	"github.com/avito-test-case/pkg/tools/http_utils"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	transactionUCase transaction.UseCase
}

func NewHandler(transactionUCase transaction.UseCase) transaction.Handler {
	return &TransactionHandler{
		transactionUCase: transactionUCase,
	}
}

func (h *TransactionHandler) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil || userId < 0 {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestParams, http.StatusBadRequest)
		return
	}

	paginator := &models.PaginatorTransaction{}
	paginator.Page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || paginator.Page < 0 {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestParams, http.StatusBadRequest)
		return
	}

	paginator.CountRecords, err = strconv.Atoi(r.URL.Query().Get("records"))
	if err != nil || paginator.CountRecords <= 0 {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestParams, http.StatusBadRequest)
		return
	}

	paginator.SortDirection = r.URL.Query().Get("direction")
	if paginator.SortDirection == "" {
		paginator.SortDirection = "asc"
	}
	paginator.SortField = r.URL.Query().Get("sort")
	if paginator.SortField == "" {
		paginator.SortField = "date"
	}

	transactions, err := h.transactionUCase.GetUserTransactions(uint64(userId), paginator)
	switch err {
	case usecase_errors.UcUserBalanceNotFound:
		http_utils.SetJSONResponse(w, handler_errors.HttpUserBalanceNotFound, http.StatusForbidden)
		return
	case usecase_errors.UcInternalServerError:
		http_utils.SetJSONResponse(w, handler_errors.HttpInternalServerError, http.StatusInternalServerError)
		return
	default:
		http_utils.SetJSONResponse(w, transactions, http.StatusOK)
	}
}
