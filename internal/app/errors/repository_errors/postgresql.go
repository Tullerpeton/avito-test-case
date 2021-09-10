package repository_errors

import "github.com/avito-test-case/internal/app/errors"

var (
	PgInternalDbError = errors.Error{
		Message: "Internal database error",
	}
	PgCanNotFind = errors.Error{
		Message: "Can't find data",
	}
	PgIncorrectPaginator = errors.Error{
		Message: "Incorrect paginator",
	}
	PgCanNotUpdate = errors.Error{
		Message: "Can't update data",
	}
)
