package dataloader

import (
	"context"
	"database/sql"

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
func (s *SqlLoader) selectProducts(companyIds []interface{}, prodTmp map[string]*[]*entity.Product) error {
	stmtProduct, err := s.conn.PrepareContext(context.TODO(), `
	SELECT company_id, id, name, prod_type, manufactured, sold, price, released 
	FROM product
	WHERE company_id IN (`+joinN(len(companyIds), "?", ",")+`);`)
	if err != nil {
		return err
	}

	prodRows, err := stmtProduct.QueryContext(context.TODO(), companyIds...)
	if err != nil {
		return err
	}
	defer prodRows.Close()
	for prodRows.Next() {
		var cid string
		var row entity.Product
		err := prodRows.Scan(&cid, &row.Id, &row.Name, &row.ProdType, &row.Manufactured, &row.Sold, &row.Price, &row.Released)
		if err != nil {
			return err
		}
		*prodTmp[cid] = append(*prodTmp[cid], &row)
	}
	return nil
}
func (s *SqlLoader) selectEmployee(companyIds []interface{}, emplTmp map[string]*[]*entity.Employee) error {
	stmtEmployee, err := s.conn.PrepareContext(context.TODO(), `
	SELECT company_id, id, firstname, lastname, email, birthdate 
	FROM employee
	WHERE company_id IN (`+joinN(len(companyIds), "?", ",")+`);`)
	if err != nil {
		return err
	}
	emplRows, err := stmtEmployee.QueryContext(context.TODO(), companyIds...)
	if err != nil {
		return err
	}
	defer emplRows.Close()
	for emplRows.Next() {
		var cid string
		var row entity.Employee
		err := emplRows.Scan(&cid, &row.Id, &row.Firstname, &row.Lastname, &row.Email, &row.Birthdate)
		if err != nil {
			return err
		}
		*emplTmp[cid] = append(*emplTmp[cid], &row)
	}
	return nil
}
func (s *SqlLoader) Select() ([]*entity.Company, error) {
	rows := []*entity.Company{}
	companyIds := []interface{}{}
	emplTmp := make(map[string]*[]*entity.Employee)
	prodTmp := make(map[string]*[]*entity.Product)
	// grp := mutext.NewWaitGroup()
	stmtCompany, err := s.conn.PrepareContext(context.TODO(), `
	SELECT id, name, branch, url, founded 
	FROM company;`)
	if err != nil {
		return rows, err
	}

	companyRows, err := stmtCompany.QueryContext(context.TODO())
	if err != nil {
		return rows, err
	}

	defer companyRows.Close()
	for companyRows.Next() {
		var row entity.Company
		err := companyRows.Scan(&row.Id, &row.Name, &row.Branch, &row.Url, &row.Founded)
		if err != nil {
			return rows, err
		}
		rows = append(rows, &row)
		companyIds = append(companyIds, row.Id)
		emplTmp[row.Id] = &row.Employees
		prodTmp[row.Id] = &row.Products
	}

	// defer grp.Done()
	err = s.selectEmployee(companyIds, emplTmp)
	if err != nil {
		return rows, err
	}

	err = s.selectProducts(companyIds, prodTmp)

	return rows, err
}

func (s *SqlLoader) Insert(c *entity.Company) error {
	stmt, err := s.conn.PrepareContext(context.TODO(), `
		INSERT INTO "company"("name", "branch", "url", "founded")
		VALUES ($1, $2, $3, $4)
		RETURNING "id"`)
	if err != nil {
		return err
	}
	err = stmt.QueryRow(c.Name, c.Branch, c.Url, c.Founded).Scan(&c.Id)
	if err != nil {
		return err
	}
	for _, empl := range c.Employees {
		stmt, err := s.conn.PrepareContext(context.TODO(), `
		INSERT INTO "employee"("company_id", "firstname", "lastname", "email", "birthdate")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`)
		if err != nil {
			return err
		}
		err = stmt.QueryRow(c.Id, empl.Firstname, empl.Lastname, empl.Email, empl.Birthdate).Scan(&empl.Id)
		if err != nil {
			return err
		}
	}
	for _, prod := range c.Products {
		stmt, err := s.conn.PrepareContext(context.TODO(), `
		INSERT INTO "product"("company_id", "name", "prod_type", "manufactured", "sold", "price", "released")
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING "id";`)
		if err != nil {
			return err
		}
		err = stmt.QueryRow(c.Id, prod.Name, prod.ProdType, prod.Manufactured, prod.Sold, prod.Price, prod.Released).Scan(&prod.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
