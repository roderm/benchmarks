// go:generate sqlboiler --wipe psql
package sqlboiler

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/roderm/benchmarks/sql/entity"
	models "github.com/roderm/benchmarks/sql/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Boiler struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Boiler {
	return &Boiler{
		conn: conn,
	}
}

func (s *Boiler) Select() (models.CompanySlice, error) {
	return models.Companies(
		qm.Load(models.CompanyRels.Employees),
		qm.Load(models.CompanyRels.Products),
	).All(context.Background(), s.conn)
}

func (s *Boiler) Insert(c *entity.Company) error {
	return fmt.Errorf("Not implemented")
}
