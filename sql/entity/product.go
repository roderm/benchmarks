package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Product struct {
	Id           string    `db:"product_id"`
	Name         string    `db:"product_name"`
	ProdType     string    `db:"product_prod_type"`
	Manufactured int       `db:"product_manufactured"`
	Sold         int       `db:"product_sold"`
	Price        float32   `db:"product_price"`
	Released     time.Time `db:"product_released"`
}

func (h *Product) Scan(value interface{}) (err error) {
	buff, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Can't cast %s to []byte", reflect.TypeOf(value))
	}
	return json.Unmarshal(buff, h)
}
