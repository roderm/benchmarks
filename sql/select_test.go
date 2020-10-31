package sql

import (
	"testing"

	"github.com/roderm/benchmarks/sql/dataloader"
	"github.com/roderm/benchmarks/sql/jsonagg"
	setup "github.com/roderm/benchmarks/sql/setup"
)

func BenchmarkJSON(b *testing.B) {
	db, err := setup.GetDbConn()
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

func BenchmarkDataloader(b *testing.B) {
	db, err := setup.GetDbConn()
	if err != nil {
		b.Fatal(err)
	}
	rows, err := dataloader.New(db).Select()
	if err != nil {
		b.Fatal(err)
	}
	if len(rows) == 0 {
		b.Fatal("No rows received")
	}
}
