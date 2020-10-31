package dataloader

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/roderm/benchmarks/sql/entity"
)

func Select() ([]*entity.Company, error) {
	var conn sql.Conn
	rows := []*entity.Company{}
	companyIds := []string{}
	emplTmp := make(map[string][]*entity.Employee)
	prodTmp := make(map[string][]*entity.Product)
	stmtCompany, err := conn.PrepareContext(context.TODO(), `SELECT id, name, branch, url, founded FROM company`)
	if err != nil {
		return rows, err
	}
	stmtEmployee, err := conn.PrepareContext(context.TODO(), `SELECT company_id, id, firstname, lastname, email, birthdate FROM employee`) // WHERE company_id IN (?, ?, ?)
	if err != nil {
		return rows, err
	}
	stmtProduct, err := conn.PrepareContext(context.TODO(), `SELECT company_id, id, name, product_type, manufactured, sold, price, released FROM product`) // WHERE company_id IN (?, ?, ?)
	if err != nil {
		return rows, err
	}
	companyRows, err := stmtCompany.QueryContext(context.TODO())
	if err != nil {
		return rows, err
	}

	defer companyRows.Close()
	for companyRows.NextResultSet() {
		var row entity.Company
		err := companyRows.Scan(&row.Id, &row.Name, &row.Branch, &row.Url)
		if err != nil {
			return rows, err
		}
		rows = append(rows, &row)
		companyIds = append(companyIds, row.Id)
		emplTmp[row.Id] = row.Employees
		prodTmp[row.Id] = row.Products
	}

	emplRows, err := stmtEmployee.QueryContext(context.TODO(), companyIds)
	if err != nil {
		return rows, err
	}
	defer emplRows.Close()
	for emplRows.NextResultSet() {
		var cid string
		var row entity.Employee
		err := emplRows.Scan(&cid, &row.Id, &row.Firstname, &row.Lastname, &row.Email, &row.Birthdate)
		if err != nil {
			return rows, err
		}
		emplTmp[cid] = append(emplTmp[cid], &row)
	}

	prodRows, err := stmtProduct.QueryContext(context.TODO(), companyIds)
	if err != nil {
		return rows, err
	}
	defer prodRows.Close()
	for prodRows.NextResultSet() {
		var cid string
		var row entity.Product
		err := prodRows.Scan(&cid, &row.Id, &row.Name, &row.ProdType, &row.Manufactured, &row.Sold, &row.Price, &row.Released)
		if err != nil {
			return nil, err
		}
		prodTmp[cid] = append(prodTmp[cid], &row)
	}
	return rows, nil
}

func Insert(c *entity.Company) error {
	return fmt.Errorf("Not implemented")
}
