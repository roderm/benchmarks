package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Product struct {
	Id           string
	Name         string
	ProdType     string
	Manufactured int
	Sold         int
	Price        float32
	Released     time.Time
}

func (h *Product) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}
