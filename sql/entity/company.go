package entity

import (
	"fmt"
	"reflect"
	"time"

	"encoding/json"
)

type Company struct {
	Id        string    `db:"company_id"`
	Name      string    `db:"company_name"`
	Branch    string    `db:"company_branch"`
	Url       string    `db:"company_url"`
	Founded   time.Time `db:"company_founded"`
	Employees []*Employee
	Products  []*Product
}

func (h *Company) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}
