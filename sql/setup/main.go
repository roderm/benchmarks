package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
	"github.com/roderm/benchmarks/sql/dataloader"
	"github.com/roderm/benchmarks/sql/setup/data"
)

const DB_NAME = "benchmark"

func MakeSetup() (*sql.DB, error) {
	db, err := CreateDB()
	if err != nil {
		return db, err
	}
	err = Schema(db)
	return db, err
}

func GetDbConn() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://root@roach:26257/"+DB_NAME+"?sslmode=disable")
}

func CreateDB() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "postgres://root@roach:26257/system?sslmode=disable")
	if err != nil {
		return conn, err
	}
	_, err = conn.Exec(`DROP DATABASE IF EXISTS ` + DB_NAME)
	if err != nil {
		return conn, err
	}
	_, err = conn.Exec(`CREATE DATABASE ` + DB_NAME)
	if err != nil {
		return conn, err
	}
	err = conn.Close()
	if err != nil {
		return conn, err
	}
	return GetDbConn()
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

func main() {
	db, err := MakeSetup()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB and Schema recreated.")
	comps := data.GetCompanies(20, 200, 100)
	loader := dataloader.New(db)
	for i, c := range comps {
		err := loader.Insert(c)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[%d] New company inserted \n", i)
	}
}