package mysql_utils

import (
	"ReshmaKolekar/bookstore_users/util/errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.Rest_Error {
	mysqlErr, OK := err.(*mysql.MySQLError)
	if !OK {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no matching record found with given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while parsing err: %s", err.Error()))
	}

	switch mysqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error while processing request")
}
