package errors

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsUniqueViolation(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == pgerrcode.UniqueViolation
	}
	return false
}
