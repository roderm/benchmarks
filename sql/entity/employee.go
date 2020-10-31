package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Employee struct {
	Id        string
	Firstname string
	Lastname  string
	Email     string
	Birthdate time.Time
}

func (h *Employee) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}
