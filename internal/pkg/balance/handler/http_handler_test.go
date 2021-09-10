package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	balance_mock "github.com/avito-test-case/internal/pkg/balance/mock"
	"github.com/avito-test-case/internal/pkg/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestBalanceHandler_GetUserBalance(t *testing.T) {
	userBalance := models.UserBalance{
		Id:       1,
		Balance:  100,
		Currency: "RUB",
	}
	currencyConvert := "USD"

	t.Run("GetUserBalance_success_without_convert", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceUCase.
			EXPECT().
			GetUserBalance(userBalance.Id, "").
			Return(&userBalance, nil)

		balanceHandler := NewHandler(balanceUCase)

		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/v1/user/balance/%d", userBalance.Id),
			nil,
		)
		vars := map[string]string{
			"id": strconv.Itoa(int(userBalance.Id)),
		}
		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.GetUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK, "incorrect http code")
	})

	t.Run("GetUserBalance_success_with_convert", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceHandler := NewHandler(balanceUCase)
		balanceUCase.
			EXPECT().
			GetUserBalance(userBalance.Id, currencyConvert).
			Return(&userBalance, nil)

		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/v1/user/balance/%d", userBalance.Id),
			nil,
		)
		vars := map[string]string{
			"id": strconv.Itoa(int(userBalance.Id)),
		}
		req = mux.SetURLVars(req, vars)

		q := req.URL.Query()
		q.Add("currency", currencyConvert)
		req.URL.RawQuery = q.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.GetUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK, "incorrect http code")
	})

	t.Run("GetUserBalance_bad_args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceHandler := NewHandler(balanceUCase)

		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/v1/user/balance/%d", userBalance.Id),
			nil,
		)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.GetUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest, "incorrect http code")
	})
}

func TestBalanceHandler_ImproveUserBalance(t *testing.T) {
	userBalance := models.UserBalance{
		Id:       1,
		Balance:  100,
		Currency: "RUB",
	}
	improveBalance := models.ImproveBalance{
		Id:    1,
		Value: 100,
	}

	t.Run("ImproveUserBalance_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceUCase.
			EXPECT().
			ImproveUserBalance(&improveBalance).
			Return(&userBalance, nil)

		balanceHandler := NewHandler(balanceUCase)

		userBytes, _ := json.Marshal(improveBalance)
		req, _ := http.NewRequest(
			"POST",
			"/api/v1/user/balance/improve",
			bytes.NewBuffer(userBytes),
		)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.ImproveUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK, "incorrect http code")
	})

	t.Run("ImproveUserBalance_bad_body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceHandler := NewHandler(balanceUCase)

		req, _ := http.NewRequest(
			"POST",
			"/api/v1/user/balance/improve",
			bytes.NewBuffer(nil),
		)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.ImproveUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest, "incorrect http code")
	})
}

func TestBalanceHandler_WithdrawUserBalance(t *testing.T) {
	userBalance := models.UserBalance{
		Id:       1,
		Balance:  100,
		Currency: "RUB",
	}
	withdrawBalance := models.WithdrawBalance{
		Id:    1,
		Value: 100,
	}

	t.Run("WithdrawUserBalance_success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceUCase.
			EXPECT().
			WithdrawUserBalance(&withdrawBalance).
			Return(&userBalance, nil)

		balanceHandler := NewHandler(balanceUCase)

		userBytes, _ := json.Marshal(withdrawBalance)
		req, _ := http.NewRequest(
			"POST",
			"/api/v1/user/balance/withdraw",
			bytes.NewBuffer(userBytes),
		)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.WithdrawUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK, "incorrect http code")
	})

	t.Run("WithdrawUserBalance_bad_body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := balance_mock.NewMockUseCase(ctrl)
		balanceHandler := NewHandler(balanceUCase)

		req, _ := http.NewRequest(
			"POST",
			"/api/v1/user/balance/withdraw",
			bytes.NewBuffer(nil),
		)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.WithdrawUserBalance)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest, "incorrect http code")
	})
}
