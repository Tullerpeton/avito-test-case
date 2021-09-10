package usecase_errors

import "github.com/avito-test-case/internal/app/errors"

var (
	UcUserBalanceNotFound = errors.Error{
		Message: "User balance not found",
	}
	UcCanNotConvertCurrency = errors.Error{
		Message: "Can not convert currency",
	}
	UcInternalServerError = errors.Error{
		Message: "Internal server error",
	}
	UcInsufficientFunds = errors.Error{
		Message: "Insufficient funds on the balance",
	}
)
