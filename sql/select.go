package sql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB_NAME = "benchmark"

func GetDbConn() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://root@localhost:26257/"+DB_NAME+"?sslmode=disable")
}
