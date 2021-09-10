package handler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/avito-test-case/internal/pkg/models"
	transaction_mock "github.com/avito-test-case/internal/pkg/transaction/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestTransactionHandler_GetUserTransactions(t *testing.T) {
	paginator := models.PaginatorTransaction{
		Page:          1,
		CountRecords:  10,
		SortField:     "date",
		SortDirection: "asc",
	}
	userId := uint64(1)
	transactions := make([]*models.Transaction, 0)

	t.Run("GetUserTransactions_success_base", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		balanceUCase := transaction_mock.NewMockUseCase(ctrl)
		balanceHandler := NewHandler(balanceUCase)
		balanceUCase.
			EXPECT().
			GetUserTransactions(userId, &paginator).
			Return(transactions, nil)

		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/v1/user/transaction/%d", userId),
			nil,
		)
		vars := map[string]string{
			"id": strconv.Itoa(int(userId)),
		}
		req = mux.SetURLVars(req, vars)

		q := req.URL.Query()
		q.Add("page", strconv.Itoa(int(paginator.Page)))
		q.Add("records", strconv.Itoa(int(paginator.CountRecords)))
		q.Add("sort", paginator.SortField)
		q.Add("direction", paginator.SortDirection)
		req.URL.RawQuery = q.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(balanceHandler.GetUserTransactions)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK, "incorrect http code")
	})
}
