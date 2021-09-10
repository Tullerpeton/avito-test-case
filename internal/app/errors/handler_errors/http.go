package handler_errors

import "github.com/avito-test-case/internal/app/errors"

var (
	HttpIncorrectRequestBody = errors.Error{
		Message: "Body of request is incorrect",
	}
	HttpIncorrectRequestParams = errors.Error{
		Message: "Params of request is incorrect",
	}
	HttpUserBalanceNotFound = errors.Error{
		Message: "User balance not found",
	}
	HttpCanNotConvertCurrency = errors.Error{
		Message: "User can not convert currency",
	}
	HttpCanNotMakeOperation = errors.Error{
		Message: "User can not make operation",
	}
	HttpInternalServerError = errors.Error{
		Message: "Internal server error",
	}
)
