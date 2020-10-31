package jsonagg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/roderm/benchmarks/sql/entity"
)

func Select() ([]*entity.Company, error) {
	var conn sql.Conn
	rows := []*entity.Company{}
	stmt, err := conn.PrepareContext(context.TODO(), `
	SELECT JSON_BUILD_OBJECT(
		'id', company.id,
		'name', company.name,
		'branch', company.branch,
		'url', company.url,
		'founded', company.founded,
		'employees', employee.value,
		'products', product.value
	) as company
	FROM
		company
		LEFT JOIN (
			SELECT
				company_id,
				JSON_AGG(JSON_BUILD_OBJECT(
					'id', employee.id,
					'firstname', employee.firstname,
					'lastname', employee.lastname,
					'email', employee.email,
					'birthdate', employee.birthdate
				)) as value
			FROM employee
			GROUP BY company_id
		) AS employee on employee.company_id = company.id
		LEFT JOIN (
			SELECT
				company_id,
				JSON_AGG(JSON_BUILD_OBJECT(
					'id', product.id,
					'name', product.name,
					'prod_type', product.prod_type,
					'manufactered', product.manufactered,
					'sold', product.sold,
					'price', product.price,
					'released', product.released
				)) as value
			FROM product
			GROUP BY company_id
		) AS product on product.company_id = company.id
	`)
	if err != nil {
		return rows, err
	}
	companyRows, err := stmt.QueryContext(context.TODO())
	if err != nil {
		return rows, err
	}
	defer companyRows.Close()
	for companyRows.NextResultSet() {
		var row *entity.Company
		err := companyRows.Scan(row)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func Insert(c *entity.Company) error {
	return fmt.Errorf("Not implemented")
}
