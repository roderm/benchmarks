// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProducts(t *testing.T) {
	t.Parallel()

	query := Products()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProductsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProductsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Products().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProductsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProductSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProductsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProductExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Product exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProductExists to return true, but got false.")
	}
}

func testProductsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	productFound, err := FindProduct(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if productFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProductsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Products().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProductsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Products().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProductsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	productOne := &Product{}
	productTwo := &Product{}
	if err = randomize.Struct(seed, productOne, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}
	if err = randomize.Struct(seed, productTwo, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = productOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = productTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Products().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProductsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	productOne := &Product{}
	productTwo := &Product{}
	if err = randomize.Struct(seed, productOne, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}
	if err = randomize.Struct(seed, productTwo, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = productOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = productTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func productBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func productAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Product) error {
	*o = Product{}
	return nil
}

func testProductsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Product{}
	o := &Product{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, productDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Product object: %s", err)
	}

	AddProductHook(boil.BeforeInsertHook, productBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	productBeforeInsertHooks = []ProductHook{}

	AddProductHook(boil.AfterInsertHook, productAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	productAfterInsertHooks = []ProductHook{}

	AddProductHook(boil.AfterSelectHook, productAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	productAfterSelectHooks = []ProductHook{}

	AddProductHook(boil.BeforeUpdateHook, productBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	productBeforeUpdateHooks = []ProductHook{}

	AddProductHook(boil.AfterUpdateHook, productAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	productAfterUpdateHooks = []ProductHook{}

	AddProductHook(boil.BeforeDeleteHook, productBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	productBeforeDeleteHooks = []ProductHook{}

	AddProductHook(boil.AfterDeleteHook, productAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	productAfterDeleteHooks = []ProductHook{}

	AddProductHook(boil.BeforeUpsertHook, productBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	productBeforeUpsertHooks = []ProductHook{}

	AddProductHook(boil.AfterUpsertHook, productAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	productAfterUpsertHooks = []ProductHook{}
}

func testProductsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProductsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(productColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProductToOneCompanyUsingCompany(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Product
	var foreign Company

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, companyDBTypes, false, companyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Company struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CompanyID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Company().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProductSlice{&local}
	if err = local.L.LoadCompany(ctx, tx, false, (*[]*Product)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Company == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Company = nil
	if err = local.L.LoadCompany(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Company == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProductToOneSetOpCompanyUsingCompany(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Product
	var b, c Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Company{&b, &c} {
		err = a.SetCompany(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Company != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Products[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CompanyID != x.ID {
			t.Error("foreign key was wrong value", a.CompanyID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CompanyID))
		reflect.Indirect(reflect.ValueOf(&a.CompanyID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CompanyID != x.ID {
			t.Error("foreign key was wrong value", a.CompanyID, x.ID)
		}
	}
}

func testProductsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProductsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProductSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProductsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Products().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	productDBTypes = map[string]string{`ID`: `uuid`, `CompanyID`: `uuid`, `Name`: `character varying`, `ProdType`: `character varying`, `Manufactured`: `integer`, `Sold`: `integer`, `Price`: `double precision`, `Released`: `timestamp without time zone`}
	_              = bytes.MinRead
)

func testProductsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(productPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(productAllColumns) == len(productPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, productDBTypes, true, productPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProductsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(productAllColumns) == len(productPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Product{}
	if err = randomize.Struct(seed, o, productDBTypes, true, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, productDBTypes, true, productPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(productAllColumns, productPrimaryKeyColumns) {
		fields = productAllColumns
	} else {
		fields = strmangle.SetComplement(
			productAllColumns,
			productPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ProductSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProductsUpsert(t *testing.T) {
	t.Parallel()

	if len(productAllColumns) == len(productPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Product{}
	if err = randomize.Struct(seed, &o, productDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Product: %s", err)
	}

	count, err := Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, productDBTypes, false, productPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Product: %s", err)
	}

	count, err = Products().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
