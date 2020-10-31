package sql

import (
	"fmt"
	"testing"

	"github.com/roderm/benchmarks/sql/data"
	"github.com/roderm/benchmarks/sql/dataloader"
	"github.com/roderm/benchmarks/sql/jsonagg"
)

func BenchmarkDBSetup(b *testing.B) {
	db, err := MakeSetup()
	if err != nil {
		b.Fatal(err)
	}
	comps := data.GetCompanies(5, 4, 3)
	b.ResetTimer()
	loader := dataloader.New(db)
	for _, c := range comps {
		err := loader.Insert(c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	db, err := GetDbConn()
	if err != nil {
		panic(err)
	}
	rows, err := jsonagg.New(db).Select()
	if err != nil {
		panic(err)
	}
	for _, c := range rows {
		fmt.Println(c.Name)
	}
}

func BenchmarkDataloader(b *testing.B) {
	db, err := GetDbConn()
	if err != nil {
		panic(err)
	}
	rows, err := dataloader.New(db).Select()
	if err != nil {
		panic(err)
	}
	if len(rows) == 0 {
		panic("no rows received")
	}
}
