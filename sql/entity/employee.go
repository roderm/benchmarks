package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Employee struct {
	Id        string    `db:"employee_id"`
	Firstname string    `db:"employee_firstname"`
	Lastname  string    `db:"employee_lastname"`
	Email     string    `db:"employee_email"`
	Birthdate time.Time `db:"employee_birthdate"`
}

func (h *Employee) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}
