package utilits

import (
	"github.com/jmoiron/sqlx"
)

type ExpectedConnections struct {
	SqlConnection *sqlx.DB
}
