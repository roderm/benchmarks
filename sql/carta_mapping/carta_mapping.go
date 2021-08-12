package carta_mapping

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackskj/carta"
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
	SELECT 
		"company"."id" AS "company_id",
		"company"."name" AS "company_name",
		"company"."branch" AS "company_branch",
		"company"."url" AS "company_url",
		"company"."founded" AS "company_founded",
		"employee"."id" AS employee_id,
		"employee"."firstname" AS employee_firstname,
		"employee"."lastname" AS employee_lastname,
		"employee"."email" AS employee_email,
		"employee"."birthdate" AS employee_birthdate,
		"product"."id" AS product_id,
		"product"."name" AS product_name,
		"product"."prod_type" AS product_prod_type,
		"product"."manufactured" AS product_manufactured,
		"product"."sold" AS product_sold,
		"product"."price" AS product_price,
		"product"."released" AS product_released
	FROM
		"company"
		LEFT JOIN "employee" ON "employee"."company_id" = "company"."id"
		LEFT JOIN "product" ON "product"."company_id" = "company"."id"
	`)
	if err != nil {
		return rows, err
	}
	companyRows, err := stmt.QueryContext(context.TODO())
	if err != nil {
		return rows, err
	}
	defer companyRows.Close()
	err = carta.Map(companyRows, &rows)
	return rows, err
}

func (s *SqlLoader) Insert(c *entity.Company) error {
	return fmt.Errorf("Not implemented")
}
