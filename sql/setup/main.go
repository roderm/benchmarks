package main

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
	msql "github.com/roderm/benchmarks/sql"
	"github.com/roderm/benchmarks/sql/dataloader"
	"github.com/roderm/benchmarks/sql/setup/data"
)

func MakeSetup() (*sql.DB, error) {
	db, err := CreateDB()
	if err != nil {
		return db, err
	}
	err = Schema(db)
	return db, err
}

func CreateDB() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "postgres://root@localhost:26257/system?sslmode=disable")
	if err != nil {
		return conn, err
	}
	_, err = conn.Exec(`DROP DATABASE IF EXISTS ` + msql.DB_NAME)
	if err != nil {
		return conn, err
	}
	_, err = conn.Exec(`CREATE DATABASE ` + msql.DB_NAME)
	if err != nil {
		return conn, err
	}
	err = conn.Close()
	if err != nil {
		return conn, err
	}
	return msql.GetDbConn()
}

func Schema(conn *sql.DB) error {
	schema, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	_, err = conn.Exec(string(schema))
	return err
}

func InsertData(conn *sql.DB, comps, empls, prods int) error {
	companies := data.GetCompanies(comps, empls, prods)
	loader := dataloader.New(conn)
	for _, c := range companies {
		err := loader.Insert(c)
		if err != nil {
			return err
		}
	}
	return nil
}
