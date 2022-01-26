// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Company is an object representing the database table.
type Company struct {
	ID      string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name    string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Branch  string      `boil:"branch" json:"branch" toml:"branch" yaml:"branch"`
	URL     null.String `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`
	Founded null.Time   `boil:"founded" json:"founded,omitempty" toml:"founded" yaml:"founded,omitempty"`

	R *companyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L companyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompanyColumns = struct {
	ID      string
	Name    string
	Branch  string
	URL     string
	Founded string
}{
	ID:      "id",
	Name:    "name",
	Branch:  "branch",
	URL:     "url",
	Founded: "founded",
}

var CompanyTableColumns = struct {
	ID      string
	Name    string
	Branch  string
	URL     string
	Founded string
}{
	ID:      "company.id",
	Name:    "company.name",
	Branch:  "company.branch",
	URL:     "company.url",
	Founded: "company.founded",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var CompanyWhere = struct {
	ID      whereHelperstring
	Name    whereHelperstring
	Branch  whereHelperstring
	URL     whereHelpernull_String
	Founded whereHelpernull_Time
}{
	ID:      whereHelperstring{field: "\"company\".\"id\""},
	Name:    whereHelperstring{field: "\"company\".\"name\""},
	Branch:  whereHelperstring{field: "\"company\".\"branch\""},
	URL:     whereHelpernull_String{field: "\"company\".\"url\""},
	Founded: whereHelpernull_Time{field: "\"company\".\"founded\""},
}

// CompanyRels is where relationship names are stored.
var CompanyRels = struct {
	Employees string
	Products  string
}{
	Employees: "Employees",
	Products:  "Products",
}

// companyR is where relationships are stored.
type companyR struct {
	Employees EmployeeSlice `boil:"Employees" json:"Employees" toml:"Employees" yaml:"Employees"`
	Products  ProductSlice  `boil:"Products" json:"Products" toml:"Products" yaml:"Products"`
}

// NewStruct creates a new relationship struct
func (*companyR) NewStruct() *companyR {
	return &companyR{}
}

// companyL is where Load methods for each relationship are stored.
type companyL struct{}

var (
	companyAllColumns            = []string{"id", "name", "branch", "url", "founded"}
	companyColumnsWithoutDefault = []string{"name", "branch", "url", "founded"}
	companyColumnsWithDefault    = []string{"id"}
	companyPrimaryKeyColumns     = []string{"id"}
)

type (
	// CompanySlice is an alias for a slice of pointers to Company.
	// This should almost always be used instead of []Company.
	CompanySlice []*Company
	// CompanyHook is the signature for custom Company hook methods
	CompanyHook func(context.Context, boil.ContextExecutor, *Company) error

	companyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	companyType                 = reflect.TypeOf(&Company{})
	companyMapping              = queries.MakeStructMapping(companyType)
	companyPrimaryKeyMapping, _ = queries.BindMapping(companyType, companyMapping, companyPrimaryKeyColumns)
	companyInsertCacheMut       sync.RWMutex
	companyInsertCache          = make(map[string]insertCache)
	companyUpdateCacheMut       sync.RWMutex
	companyUpdateCache          = make(map[string]updateCache)
	companyUpsertCacheMut       sync.RWMutex
	companyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var companyBeforeInsertHooks []CompanyHook
var companyBeforeUpdateHooks []CompanyHook
var companyBeforeDeleteHooks []CompanyHook
var companyBeforeUpsertHooks []CompanyHook

var companyAfterInsertHooks []CompanyHook
var companyAfterSelectHooks []CompanyHook
var companyAfterUpdateHooks []CompanyHook
var companyAfterDeleteHooks []CompanyHook
var companyAfterUpsertHooks []CompanyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Company) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Company) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Company) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Company) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Company) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Company) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Company) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Company) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Company) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompanyHook registers your hook function for all future operations.
func AddCompanyHook(hookPoint boil.HookPoint, companyHook CompanyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		companyBeforeInsertHooks = append(companyBeforeInsertHooks, companyHook)
	case boil.BeforeUpdateHook:
		companyBeforeUpdateHooks = append(companyBeforeUpdateHooks, companyHook)
	case boil.BeforeDeleteHook:
		companyBeforeDeleteHooks = append(companyBeforeDeleteHooks, companyHook)
	case boil.BeforeUpsertHook:
		companyBeforeUpsertHooks = append(companyBeforeUpsertHooks, companyHook)
	case boil.AfterInsertHook:
		companyAfterInsertHooks = append(companyAfterInsertHooks, companyHook)
	case boil.AfterSelectHook:
		companyAfterSelectHooks = append(companyAfterSelectHooks, companyHook)
	case boil.AfterUpdateHook:
		companyAfterUpdateHooks = append(companyAfterUpdateHooks, companyHook)
	case boil.AfterDeleteHook:
		companyAfterDeleteHooks = append(companyAfterDeleteHooks, companyHook)
	case boil.AfterUpsertHook:
		companyAfterUpsertHooks = append(companyAfterUpsertHooks, companyHook)
	}
}

// One returns a single company record from the query.
func (q companyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Company, error) {
	o := &Company{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for company")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Company records from the query.
func (q companyQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompanySlice, error) {
	var o []*Company

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Company slice")
	}

	if len(companyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Company records in the query.
func (q companyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count company rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q companyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if company exists")
	}

	return count > 0, nil
}

// Employees retrieves all the employee's Employees with an executor.
func (o *Company) Employees(mods ...qm.QueryMod) employeeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"employee\".\"company_id\"=?", o.ID),
	)

	query := Employees(queryMods...)
	queries.SetFrom(query.Query, "\"employee\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"employee\".*"})
	}

	return query
}

// Products retrieves all the product's Products with an executor.
func (o *Company) Products(mods ...qm.QueryMod) productQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"product\".\"company_id\"=?", o.ID),
	)

	query := Products(queryMods...)
	queries.SetFrom(query.Query, "\"product\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"product\".*"})
	}

	return query
}

// LoadEmployees allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (companyL) LoadEmployees(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCompany interface{}, mods queries.Applicator) error {
	var slice []*Company
	var object *Company

	if singular {
		object = maybeCompany.(*Company)
	} else {
		slice = *maybeCompany.(*[]*Company)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &companyR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &companyR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`employee`),
		qm.WhereIn(`employee.company_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load employee")
	}

	var resultSlice []*Employee
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice employee")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on employee")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for employee")
	}

	if len(employeeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Employees = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &employeeR{}
			}
			foreign.R.Company = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CompanyID {
				local.R.Employees = append(local.R.Employees, foreign)
				if foreign.R == nil {
					foreign.R = &employeeR{}
				}
				foreign.R.Company = local
				break
			}
		}
	}

	return nil
}

// LoadProducts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (companyL) LoadProducts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCompany interface{}, mods queries.Applicator) error {
	var slice []*Company
	var object *Company

	if singular {
		object = maybeCompany.(*Company)
	} else {
		slice = *maybeCompany.(*[]*Company)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &companyR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &companyR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`product`),
		qm.WhereIn(`product.company_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load product")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on product")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for product")
	}

	if len(productAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Products = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &productR{}
			}
			foreign.R.Company = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CompanyID {
				local.R.Products = append(local.R.Products, foreign)
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.Company = local
				break
			}
		}
	}

	return nil
}

// AddEmployees adds the given related objects to the existing relationships
// of the company, optionally inserting them as new records.
// Appends related to o.R.Employees.
// Sets related.R.Company appropriately.
func (o *Company) AddEmployees(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Employee) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CompanyID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"employee\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"company_id"}),
				strmangle.WhereClause("\"", "\"", 2, employeePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CompanyID = o.ID
		}
	}

	if o.R == nil {
		o.R = &companyR{
			Employees: related,
		}
	} else {
		o.R.Employees = append(o.R.Employees, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &employeeR{
				Company: o,
			}
		} else {
			rel.R.Company = o
		}
	}
	return nil
}

// AddProducts adds the given related objects to the existing relationships
// of the company, optionally inserting them as new records.
// Appends related to o.R.Products.
// Sets related.R.Company appropriately.
func (o *Company) AddProducts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Product) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CompanyID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"product\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"company_id"}),
				strmangle.WhereClause("\"", "\"", 2, productPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CompanyID = o.ID
		}
	}

	if o.R == nil {
		o.R = &companyR{
			Products: related,
		}
	} else {
		o.R.Products = append(o.R.Products, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &productR{
				Company: o,
			}
		} else {
			rel.R.Company = o
		}
	}
	return nil
}

// Companies retrieves all the records using an executor.
func Companies(mods ...qm.QueryMod) companyQuery {
	mods = append(mods, qm.From("\"company\""))
	return companyQuery{NewQuery(mods...)}
}

// FindCompany retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompany(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Company, error) {
	companyObj := &Company{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"company\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, companyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from company")
	}

	if err = companyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return companyObj, err
	}

	return companyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Company) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	companyInsertCacheMut.RLock()
	cache, cached := companyInsertCache[key]
	companyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			companyAllColumns,
			companyColumnsWithDefault,
			companyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(companyType, companyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(companyType, companyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"company\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"company\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into company")
	}

	if !cached {
		companyInsertCacheMut.Lock()
		companyInsertCache[key] = cache
		companyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Company.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Company) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	companyUpdateCacheMut.RLock()
	cache, cached := companyUpdateCache[key]
	companyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			companyAllColumns,
			companyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update company, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"company\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, companyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(companyType, companyMapping, append(wl, companyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update company row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for company")
	}

	if !cached {
		companyUpdateCacheMut.Lock()
		companyUpdateCache[key] = cache
		companyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q companyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for company")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for company")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompanySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"company\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, companyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in company slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all company")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Company) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	companyUpsertCacheMut.RLock()
	cache, cached := companyUpsertCache[key]
	companyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			companyAllColumns,
			companyColumnsWithDefault,
			companyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			companyAllColumns,
			companyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert company, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(companyPrimaryKeyColumns))
			copy(conflict, companyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"company\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(companyType, companyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(companyType, companyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert company")
	}

	if !cached {
		companyUpsertCacheMut.Lock()
		companyUpsertCache[key] = cache
		companyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Company record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Company) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Company provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), companyPrimaryKeyMapping)
	sql := "DELETE FROM \"company\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from company")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for company")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q companyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no companyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from company")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompanySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(companyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"company\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, companyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from company slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company")
	}

	if len(companyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Company) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompany(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompanySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompanySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"company\".* FROM \"company\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, companyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompanySlice")
	}

	*o = slice

	return nil
}

// CompanyExists checks if the Company row exists.
func CompanyExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"company\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if company exists")
	}

	return exists, nil
}
