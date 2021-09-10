package handler

import (
	"net/http"
	"strconv"

	"github.com/avito-test-case/internal/app/errors/handler_errors"
	"github.com/avito-test-case/internal/app/errors/usecase_errors"
	"github.com/avito-test-case/internal/pkg/balance"
	"github.com/avito-test-case/internal/pkg/models"
	"github.com/avito-test-case/pkg/tools/http_utils"

	"github.com/gorilla/mux"
)

type BalanceHandler struct {
	balanceUCase balance.UseCase
}

func NewHandler(balanceUCase balance.UseCase) balance.Handler {
	return &BalanceHandler{
		balanceUCase: balanceUCase,
	}
}

func (h *BalanceHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil || userId < 0 {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestParams, http.StatusBadRequest)
		return
	}
	currencyToConvert := r.URL.Query().Get("currency")

	userBalance, err := h.balanceUCase.GetUserBalance(uint64(userId), currencyToConvert)
	switch err {
	case usecase_errors.UcUserBalanceNotFound:
		http_utils.SetJSONResponse(w, handler_errors.HttpUserBalanceNotFound, http.StatusForbidden)
		return
	case usecase_errors.UcCanNotConvertCurrency:
		http_utils.SetJSONResponse(w, handler_errors.HttpCanNotConvertCurrency, http.StatusConflict)
		return
	default:
		http_utils.SetJSONResponse(w, userBalance, http.StatusOK)
	}
}

func (h *BalanceHandler) WithdrawUserBalance(w http.ResponseWriter, r *http.Request) {
	withdrawBalance := &models.WithdrawBalance{}
	if err := http_utils.ParseBody(r, withdrawBalance); err != nil {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestBody, http.StatusBadRequest)
		return
	}

	updatedUserBalance, err := h.balanceUCase.WithdrawUserBalance(withdrawBalance)
	switch err {
	case usecase_errors.UcUserBalanceNotFound:
		http_utils.SetJSONResponse(w, handler_errors.HttpUserBalanceNotFound, http.StatusForbidden)
		return
	case usecase_errors.UcInsufficientFunds:
		http_utils.SetJSONResponse(w, handler_errors.HttpCanNotMakeOperation, http.StatusConflict)
		return
	case usecase_errors.UcInternalServerError:
		http_utils.SetJSONResponse(w, handler_errors.HttpInternalServerError, http.StatusInternalServerError)
		return
	default:
		http_utils.SetJSONResponse(w, updatedUserBalance, http.StatusOK)
	}
}

func (h *BalanceHandler) ImproveUserBalance(w http.ResponseWriter, r *http.Request) {
	improveBalance := &models.ImproveBalance{}
	if err := http_utils.ParseBody(r, improveBalance); err != nil {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestBody, http.StatusBadRequest)
		return
	}

	updatedUserBalance, err := h.balanceUCase.ImproveUserBalance(improveBalance)
	switch err {
	case usecase_errors.UcUserBalanceNotFound:
		http_utils.SetJSONResponse(w, handler_errors.HttpUserBalanceNotFound, http.StatusForbidden)
		return
	case usecase_errors.UcInternalServerError:
		http_utils.SetJSONResponse(w, handler_errors.HttpInternalServerError, http.StatusInternalServerError)
		return
	default:
		http_utils.SetJSONResponse(w, updatedUserBalance, http.StatusOK)
	}
}

func (h *BalanceHandler) MakeUserBalanceTransfer(w http.ResponseWriter, r *http.Request) {
	transfer := &models.Transfer{}
	if err := http_utils.ParseBody(r, transfer); err != nil {
		http_utils.SetJSONResponse(w, handler_errors.HttpIncorrectRequestBody, http.StatusBadRequest)
		return
	}

	transferResult, err := h.balanceUCase.MakeUserBalanceTransfer(transfer)
	switch err {
	case usecase_errors.UcUserBalanceNotFound:
		http_utils.SetJSONResponse(w, handler_errors.HttpCanNotMakeOperation, http.StatusConflict)
		return
	case usecase_errors.UcInternalServerError:
		http_utils.SetJSONResponse(w, handler_errors.HttpInternalServerError, http.StatusInternalServerError)
		return
	default:
		http_utils.SetJSONResponse(w, transferResult, http.StatusOK)
	}
}
