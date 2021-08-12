package sql

import (
	"testing"

	"github.com/roderm/benchmarks/sql/carta_mapping"
	"github.com/roderm/benchmarks/sql/dataloader"
	"github.com/roderm/benchmarks/sql/jsonagg"
)

// func BenchmarkSQLMap(b *testing.B) {
// 	db, err := GetDbConn()
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	store := protomap.NewStore(db)
// 	rows, err := store.Company(context.TODO(), protomap.CompanyWithEmployee(), protomap.CompanyWithProduct())
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	for _, r := range rows {
// 		if len(r.Employees) == 0 {
// 			b.Fatal("No Employees received")
// 		}
// 		if len(r.Products) == 0 {
// 			b.Fatal("No Products received")
// 		}
// 	}
// }

func BenchmarkDataloader(b *testing.B) {
	db, err := GetDbConn()
	if err != nil {
		b.Fatal(err)
	}
	rows, err := dataloader.New(db).Select()
	if err != nil {
		b.Fatal(err)
	}
	for _, r := range rows {
		if len(r.Employees) == 0 {
			b.Fatal("No Employees received")
		}
		if len(r.Products) == 0 {
			b.Fatal("No Products received")
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	db, err := GetDbConn()
	if err != nil {
		b.Fatal(err)
	}
	rows, err := jsonagg.New(db).Select()
	if err != nil {
		b.Fatal(err)
	}
	if len(rows) == 0 {
		b.Fatal("No rows received")
	}
}

func BenchmarkCarta(b *testing.B) {
	db, err := GetDbConn()
	if err != nil {
		b.Fatal(err)
	}
	c := carta_mapping.New(db)
	rows, err := c.Select()
	if err != nil {
		b.Fatal(err)
	}
	for _, r := range rows {
		if len(r.Employees) == 0 {
			b.Fatal("No Employees received")
		}
		if len(r.Products) == 0 {
			b.Fatal("No Products received")
		}
	}
}

func TestCarta(t *testing.T) {
	db, err := GetDbConn()
	if err != nil {
		t.Fatal(err)
	}
	c := carta_mapping.New(db)
	rows, err := c.Select()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rows {
		if len(r.Employees) == 0 {
			t.Fatal("No Employees received")
		}
		if len(r.Products) == 0 {
			t.Fatal("No Products received")
		}
	}
}
