package sql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB_NAME = "benchmark"

func GetDbConn() (*sql.DB, error) {
	return sql.Open("postgres",
		"postgres://roderm:password1234@localhost:5432/"+DB_NAME+"?sslmode=disable",
	)
}
