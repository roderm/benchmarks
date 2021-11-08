package protomap

import (
	context "context"
	sql "database/sql"
	driver "database/sql/driver"
	json "encoding/json"
	fmt "fmt"
	pg "github.com/roderm/protoc-gen-go-sqlmap/lib/go/pg"
)

var _ = fmt.Sprintf
var _ = context.TODO
var _ = pg.NONE
var _ = sql.Open
var _ = driver.IsValue
var _ = json.Valid

type CompanyStore struct {
	conn *sql.DB
}

func NewCompanyStore(conn *sql.DB) *CompanyStore {
	return &CompanyStore{conn}
}

func (m *Company) Scan(value interface{}) error {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed %+v", value)
	}
	m.Id = string(buff)
	return nil
}

func (m *Company) Value() (driver.Value, error) {
	return m.Id, nil
}

type queryCompanyConfig struct {
	Store        *CompanyStore
	filter       pg.Where
	beforeReturn []func(map[interface{}]*Company) error
	cb           []func(*Company)
	rows         map[interface{}]*Company

	loadEmployees bool
	optsEmployees []EmployeeOption

	loadProducts bool
	optsProducts []ProductOption

	loadDebitors bool
	optsDebitors []ContractOption

	loadCreditors bool
	optsCreditors []ContractOption
}

type CompanyOption func(*queryCompanyConfig)

func CompanyFilter(filter pg.Where) CompanyOption {
	return func(config *queryCompanyConfig) {
		if config.filter == nil {
			config.filter = filter
		} else {
			config.filter = pg.AND(config.filter, filter)
		}
	}
}

func CompanyOnRow(cb func(*Company)) CompanyOption {
	return func(s *queryCompanyConfig) {
		s.cb = append(s.cb, cb)
	}
}

func CompanyWithEmployees(opts ...EmployeeOption) CompanyOption {
	return func(config *queryCompanyConfig) {
		mapEmployees := make(map[interface{}]*Company)
		config.loadEmployees = true
		config.optsEmployees = opts
		config.cb = append(config.cb, func(row *Company) {
			// repeated
			mapEmployees[row.Id] = row

		})
		config.optsEmployees = append(config.optsEmployees,
			EmployeeOnRow(func(row *Employee) {

				// repeated
				if config.rows[row.Company.Id] != nil {
					config.rows[row.Company.Id].Employees = append(config.rows[row.Company.Id].Employees, row)
				}

			}),
			EmployeeFilter(pg.INCallabel("company_id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapEmployees {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}
func CompanyWithProducts(opts ...ProductOption) CompanyOption {
	return func(config *queryCompanyConfig) {
		mapProducts := make(map[interface{}]*Company)
		config.loadProducts = true
		config.optsProducts = opts
		config.cb = append(config.cb, func(row *Company) {
			// repeated
			mapProducts[row.Id] = row

		})
		config.optsProducts = append(config.optsProducts,
			ProductOnRow(func(row *Product) {

				// repeated
				if config.rows[row.Company.Id] != nil {
					config.rows[row.Company.Id].Products = append(config.rows[row.Company.Id].Products, row)
				}

			}),
			ProductFilter(pg.INCallabel("company_id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapProducts {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}
func CompanyWithDebitors(opts ...ContractOption) CompanyOption {
	return func(config *queryCompanyConfig) {
		mapDebitors := make(map[interface{}]*Company)
		config.loadDebitors = true
		config.optsDebitors = opts
		config.cb = append(config.cb, func(row *Company) {
			// repeated
			mapDebitors[row.Id] = row

		})
		config.optsDebitors = append(config.optsDebitors,
			ContractOnRow(func(row *Contract) {

				// repeated
				if config.rows[row.Debitor.Id] != nil {
					config.rows[row.Debitor.Id].Debitors = append(config.rows[row.Debitor.Id].Debitors, row)
				}

			}),
			ContractFilter(pg.INCallabel("debitor_id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapDebitors {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}
func CompanyWithCreditors(opts ...ContractOption) CompanyOption {
	return func(config *queryCompanyConfig) {
		mapCreditors := make(map[interface{}]*Company)
		config.loadCreditors = true
		config.optsCreditors = opts
		config.cb = append(config.cb, func(row *Company) {
			// repeated
			mapCreditors[row.Id] = row

		})
		config.optsCreditors = append(config.optsCreditors,
			ContractOnRow(func(row *Contract) {

				// repeated
				if config.rows[row.Creditor.Id] != nil {
					config.rows[row.Creditor.Id].Creditors = append(config.rows[row.Creditor.Id].Creditors, row)
				}

			}),
			ContractFilter(pg.INCallabel("creditor_id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapCreditors {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}

func (s *CompanyStore) Company(ctx context.Context, opts ...CompanyOption) (map[interface{}]*Company, error) {
	config := &queryCompanyConfig{
		Store:  s,
		filter: pg.NONE(),
		rows:   make(map[interface{}]*Company),
	}
	for _, o := range opts {
		o(config)
	}
	err := s.selectCompany(ctx, config.filter, func(row *Company) {
		config.rows[row.Id] = row
		for _, cb := range config.cb {
			cb(row)
		}
	})
	if err != nil {
		return config.rows, err
	}

	if config.loadDebitors {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Contract(ctx, config.optsDebitors...)

	}
	if err != nil {
		return config.rows, err
	}

	if config.loadCreditors {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Contract(ctx, config.optsCreditors...)

	}
	if err != nil {
		return config.rows, err
	}

	if config.loadEmployees {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Employee(ctx, config.optsEmployees...)

	}
	if err != nil {
		return config.rows, err
	}

	if config.loadProducts {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Product(ctx, config.optsProducts...)

	}
	if err != nil {
		return config.rows, err
	}

	for _, cb := range config.beforeReturn {
		err = cb(config.rows)
		if err != nil {
			return config.rows, err
		}
	}
	return config.rows, nil
}
func (s *CompanyStore) selectCompany(ctx context.Context, filter pg.Where, withRow func(*Company)) error {
	where, vals := pg.GetWhereClause(filter, nil)
	stmt, err := s.conn.PrepareContext(ctx, ` 
	SELECT "id", "name", "branch", "url", "founded" 
	FROM "company"
	`+where)
	if err != nil {
		return err
	}
	cursor, err := stmt.QueryContext(ctx, vals...)
	if err != nil {
		return err
	}
	defer cursor.Close()
	for cursor.Next() {
		row := new(Company)
		err := cursor.Scan(&row.Id, &row.Name, &row.Branch, &row.Url, &row.Founded)
		if err != nil {
			return err
		}
		withRow(row)
	}
	return nil
}

func (m *Contract) Scan(value interface{}) error {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed %+v", value)
	}
	m.Id = string(buff)
	return nil
}

func (m *Contract) Value() (driver.Value, error) {
	return m.Id, nil
}

type queryContractConfig struct {
	Store        *CompanyStore
	filter       pg.Where
	beforeReturn []func(map[interface{}]*Contract) error
	cb           []func(*Contract)
	rows         map[interface{}]*Contract

	loadDebitor bool
	optsDebitor []CompanyOption

	loadCreditor bool
	optsCreditor []CompanyOption
}

type ContractOption func(*queryContractConfig)

func ContractFilter(filter pg.Where) ContractOption {
	return func(config *queryContractConfig) {
		if config.filter == nil {
			config.filter = filter
		} else {
			config.filter = pg.AND(config.filter, filter)
		}
	}
}

func ContractOnRow(cb func(*Contract)) ContractOption {
	return func(s *queryContractConfig) {
		s.cb = append(s.cb, cb)
	}
}

func ContractWithCreditor(opts ...CompanyOption) ContractOption {
	return func(config *queryContractConfig) {
		mapCreditor := make(map[interface{}]*Contract)
		config.loadCreditor = true
		config.optsCreditor = opts
		config.cb = append(config.cb, func(row *Contract) {
			// one-to-one
			mapCreditor[row.Creditor.Id] = row

		})
		config.optsCreditor = append(config.optsCreditor,
			CompanyOnRow(func(row *Company) {

				// one-to-one
				item := mapCreditor[row.Id]
				if config.rows[item.Id] != nil {
					config.rows[item.Id].Creditor = row
				}

			}),
			CompanyFilter(pg.INCallabel("id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapCreditor {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}
func ContractWithDebitor(opts ...CompanyOption) ContractOption {
	return func(config *queryContractConfig) {
		mapDebitor := make(map[interface{}]*Contract)
		config.loadDebitor = true
		config.optsDebitor = opts
		config.cb = append(config.cb, func(row *Contract) {
			// one-to-one
			mapDebitor[row.Debitor.Id] = row

		})
		config.optsDebitor = append(config.optsDebitor,
			CompanyOnRow(func(row *Company) {

				// one-to-one
				item := mapDebitor[row.Id]
				if config.rows[item.Id] != nil {
					config.rows[item.Id].Debitor = row
				}

			}),
			CompanyFilter(pg.INCallabel("id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapDebitor {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}

func (s *CompanyStore) Contract(ctx context.Context, opts ...ContractOption) (map[interface{}]*Contract, error) {
	config := &queryContractConfig{
		Store:  s,
		filter: pg.NONE(),
		rows:   make(map[interface{}]*Contract),
	}
	for _, o := range opts {
		o(config)
	}
	err := s.selectContract(ctx, config.filter, func(row *Contract) {
		config.rows[row.Id] = row
		for _, cb := range config.cb {
			cb(row)
		}
	})
	if err != nil {
		return config.rows, err
	}

	if config.loadCreditor {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Company(ctx, config.optsCreditor...)

	}
	if err != nil {
		return config.rows, err
	}

	if config.loadDebitor {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Company(ctx, config.optsDebitor...)

	}
	if err != nil {
		return config.rows, err
	}

	for _, cb := range config.beforeReturn {
		err = cb(config.rows)
		if err != nil {
			return config.rows, err
		}
	}
	return config.rows, nil
}
func (s *CompanyStore) selectContract(ctx context.Context, filter pg.Where, withRow func(*Contract)) error {
	where, vals := pg.GetWhereClause(filter, nil)
	stmt, err := s.conn.PrepareContext(ctx, ` 
	SELECT "id", "creditor_id", "debitor_id", "amount" 
	FROM "contract"
	`+where)
	if err != nil {
		return err
	}
	cursor, err := stmt.QueryContext(ctx, vals...)
	if err != nil {
		return err
	}
	defer cursor.Close()
	for cursor.Next() {
		row := new(Contract)
		err := cursor.Scan(&row.Id, &row.Creditor, &row.Debitor, &row.Amount)
		if err != nil {
			return err
		}
		withRow(row)
	}
	return nil
}

func (m *Employee) Scan(value interface{}) error {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed %+v", value)
	}
	m.Id = string(buff)
	return nil
}

func (m *Employee) Value() (driver.Value, error) {
	return m.Id, nil
}

type queryEmployeeConfig struct {
	Store        *CompanyStore
	filter       pg.Where
	beforeReturn []func(map[interface{}]*Employee) error
	cb           []func(*Employee)
	rows         map[interface{}]*Employee

	loadCompany bool
	optsCompany []CompanyOption
}

type EmployeeOption func(*queryEmployeeConfig)

func EmployeeFilter(filter pg.Where) EmployeeOption {
	return func(config *queryEmployeeConfig) {
		if config.filter == nil {
			config.filter = filter
		} else {
			config.filter = pg.AND(config.filter, filter)
		}
	}
}

func EmployeeOnRow(cb func(*Employee)) EmployeeOption {
	return func(s *queryEmployeeConfig) {
		s.cb = append(s.cb, cb)
	}
}

func EmployeeWithCompany(opts ...CompanyOption) EmployeeOption {
	return func(config *queryEmployeeConfig) {
		mapCompany := make(map[interface{}]*Employee)
		config.loadCompany = true
		config.optsCompany = opts
		config.cb = append(config.cb, func(row *Employee) {
			// one-to-one
			mapCompany[row.Company.Id] = row

		})
		config.optsCompany = append(config.optsCompany,
			CompanyOnRow(func(row *Company) {

				// one-to-one
				item := mapCompany[row.Id]
				if config.rows[item.Id] != nil {
					config.rows[item.Id].Company = row
				}

			}),
			CompanyFilter(pg.INCallabel("id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapCompany {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}

func (s *CompanyStore) Employee(ctx context.Context, opts ...EmployeeOption) (map[interface{}]*Employee, error) {
	config := &queryEmployeeConfig{
		Store:  s,
		filter: pg.NONE(),
		rows:   make(map[interface{}]*Employee),
	}
	for _, o := range opts {
		o(config)
	}
	err := s.selectEmployee(ctx, config.filter, func(row *Employee) {
		config.rows[row.Id] = row
		for _, cb := range config.cb {
			cb(row)
		}
	})
	if err != nil {
		return config.rows, err
	}

	if config.loadCompany {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Company(ctx, config.optsCompany...)

	}
	if err != nil {
		return config.rows, err
	}

	for _, cb := range config.beforeReturn {
		err = cb(config.rows)
		if err != nil {
			return config.rows, err
		}
	}
	return config.rows, nil
}
func (s *CompanyStore) selectEmployee(ctx context.Context, filter pg.Where, withRow func(*Employee)) error {
	where, vals := pg.GetWhereClause(filter, nil)
	stmt, err := s.conn.PrepareContext(ctx, ` 
	SELECT "id", "company_id", "firstname", "lastname", "email", "birthdate" 
	FROM "employee"
	`+where)
	if err != nil {
		return err
	}
	cursor, err := stmt.QueryContext(ctx, vals...)
	if err != nil {
		return err
	}
	defer cursor.Close()
	for cursor.Next() {
		row := new(Employee)
		err := cursor.Scan(&row.Id, &row.Company, &row.Firstname, &row.Lastname, &row.Email, &row.Birthdate)
		if err != nil {
			return err
		}
		withRow(row)
	}
	return nil
}

func (m *Product) Scan(value interface{}) error {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed %+v", value)
	}
	m.Id = string(buff)
	return nil
}

func (m *Product) Value() (driver.Value, error) {
	return m.Id, nil
}

type queryProductConfig struct {
	Store        *CompanyStore
	filter       pg.Where
	beforeReturn []func(map[interface{}]*Product) error
	cb           []func(*Product)
	rows         map[interface{}]*Product

	loadCompany bool
	optsCompany []CompanyOption
}

type ProductOption func(*queryProductConfig)

func ProductFilter(filter pg.Where) ProductOption {
	return func(config *queryProductConfig) {
		if config.filter == nil {
			config.filter = filter
		} else {
			config.filter = pg.AND(config.filter, filter)
		}
	}
}

func ProductOnRow(cb func(*Product)) ProductOption {
	return func(s *queryProductConfig) {
		s.cb = append(s.cb, cb)
	}
}

func ProductWithCompany(opts ...CompanyOption) ProductOption {
	return func(config *queryProductConfig) {
		mapCompany := make(map[interface{}]*Product)
		config.loadCompany = true
		config.optsCompany = opts
		config.cb = append(config.cb, func(row *Product) {
			// one-to-one
			mapCompany[row.Company.Id] = row

		})
		config.optsCompany = append(config.optsCompany,
			CompanyOnRow(func(row *Company) {

				// one-to-one
				item := mapCompany[row.Id]
				if config.rows[item.Id] != nil {
					config.rows[item.Id].Company = row
				}

			}),
			CompanyFilter(pg.INCallabel("id", func() []interface{} {
				ids := []interface{}{}
				for id := range mapCompany {
					ids = append(ids, id)
				}
				return ids
			})),
		)
	}
}

func (s *CompanyStore) Product(ctx context.Context, opts ...ProductOption) (map[interface{}]*Product, error) {
	config := &queryProductConfig{
		Store:  s,
		filter: pg.NONE(),
		rows:   make(map[interface{}]*Product),
	}
	for _, o := range opts {
		o(config)
	}
	err := s.selectProduct(ctx, config.filter, func(row *Product) {
		config.rows[row.Id] = row
		for _, cb := range config.cb {
			cb(row)
		}
	})
	if err != nil {
		return config.rows, err
	}

	if config.loadCompany {
		// github.com/roderm/benchmarks/sql/protomap/sql/protomap

		_, err = s.Company(ctx, config.optsCompany...)

	}
	if err != nil {
		return config.rows, err
	}

	for _, cb := range config.beforeReturn {
		err = cb(config.rows)
		if err != nil {
			return config.rows, err
		}
	}
	return config.rows, nil
}
func (s *CompanyStore) selectProduct(ctx context.Context, filter pg.Where, withRow func(*Product)) error {
	where, vals := pg.GetWhereClause(filter, nil)
	stmt, err := s.conn.PrepareContext(ctx, ` 
	SELECT "id", "company_id", "name", "prod_type", "manufactured", "sold", "price", "released" 
	FROM "product"
	`+where)
	if err != nil {
		return err
	}
	cursor, err := stmt.QueryContext(ctx, vals...)
	if err != nil {
		return err
	}
	defer cursor.Close()
	for cursor.Next() {
		row := new(Product)
		err := cursor.Scan(&row.Id, &row.Company, &row.Name, &row.ProdType, &row.Manufactured, &row.Sold, &row.Price, &row.Released)
		if err != nil {
			return err
		}
		withRow(row)
	}
	return nil
}
