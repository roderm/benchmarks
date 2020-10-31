package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mailru/easyjson"
)

type Company struct {
	Id        string
	Name      string
	Branch    string
	Url       string
	Founded   time.Time
	Employees []*Employee
	Products  []*Product
}

func (h *Company) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return easyjson.Unmarshal(buff, h)
}
