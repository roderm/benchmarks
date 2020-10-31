package jsonagg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/roderm/benchmarks/sql/entity"
)

type SqlLoader struct {
	conn *sql.DB
}

func New(conn *sql.DB) *SqlLoader {
	return &SqlLoader{
		conn: conn,
	}
}

func (s *SqlLoader) Select() ([]*entity.Company, error) {
	rows := []*entity.Company{}
	stmt, err := s.conn.PrepareContext(context.TODO(), `
	SELECT JSON_BUILD_OBJECT(
		'id', "company"."id",
		'name', "company"."name",
		'branch', "company"."branch",
		'url', "company"."url",
		'founded', "company"."founded",
		'employees', "employee"."value",
		'products', "product"."value"
	) as "company"
	FROM
		"company"
		LEFT JOIN (
			SELECT
				"company_id",
				JSON_AGG(JSON_BUILD_OBJECT(
					'id', "employee"."id",
					'firstname', "employee"."firstname",
					'lastname', "employee"."lastname",
					'email', "employee"."email",
					'birthdate', "employee"."birthdate"
				)) as value
			FROM "employee"
			GROUP BY "company_id"
		) AS "employee" on "employee"."company_id" = "company"."id"
		LEFT JOIN (
			SELECT
				"company_id",
				JSON_AGG(JSON_BUILD_OBJECT(
					'id', "product"."id",
					'name', "product"."name",
					'prod_type', "product"."prod_type",
					'manufactured', "product"."manufactured",
					'sold', "product"."sold",
					'price', "product"."price",
					'released', "product"."released"
				)) as "value"
			FROM "product"
			GROUP BY "company_id"
		) AS "product" on "product"."company_id" = "company"."id"
	`)
	if err != nil {
		return rows, err
	}
	companyRows, err := stmt.QueryContext(context.TODO())
	if err != nil {
		return rows, err
	}
	defer companyRows.Close()
	for companyRows.Next() {
		row := new(entity.Company)
		err := companyRows.Scan(row)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (s *SqlLoader) Insert(c *entity.Company) error {
	return fmt.Errorf("Not implemented")
}
