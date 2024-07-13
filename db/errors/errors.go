package errors

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func IsUniqueViolation(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == pgerrcode.UniqueViolation
	}
	return false
}
