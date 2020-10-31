package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type ArrEmployee []*Employee
type ArrProduct []*Product

type Company struct {
	Id        string
	Name      string
	Branch    string
	Url       string
	Founded   time.Time
	Employees ArrEmployee
	Products  ArrProduct
}

func (h *Company) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}

func (h *ArrEmployee) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}

func (h *ArrProduct) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}
