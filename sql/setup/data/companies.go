package data

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/roderm/benchmarks/sql/entity"
)

func loadCompanyJson() (list []*entity.Company) {
	f, err := ioutil.ReadFile("setup/data/companies.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &list)
	if err != nil {
		panic(err)
	}
	return
}
func loadEmployeeJson() (list []*entity.Employee) {
	f, err := ioutil.ReadFile("setup/data/employees.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &list)
	if err != nil {
		panic(err)
	}
	return
}
func loadProductJson() (list []*entity.Product) {
	f, err := ioutil.ReadFile("setup/data/products.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &list)
	if err != nil {
		panic(err)
	}
	return
}
func GetCompanies(num int, numEmpl int, numProd int) (comp []*entity.Company) {
	myCompanies := loadCompanyJson()
	myEmployees := loadEmployeeJson()
	myProduct := loadProductJson()
	randomCompany := func() (comp *entity.Company) {
		comp = myCompanies[rand.Intn(len(myCompanies)-1)]
		return
	}
	randomEmployee := func() (empl *entity.Employee) {
		empl = myEmployees[rand.Intn(len(myEmployees)-1)]
		return
	}
	randomProduct := func() (prod *entity.Product) {
		prod = myProduct[rand.Intn(len(myProduct)-1)]
		return
	}
	for i_c := 0; i_c < num; i_c++ {
		nc := randomCompany()
		for i_e := 0; i_e < numEmpl; i_e++ {
			nc.Employees = append(nc.Employees, randomEmployee())
		}
		for i_p := 0; i_p < numEmpl; i_p++ {
			nc.Products = append(nc.Products, randomProduct())
		}
		comp = append(comp, nc)
	}
	return
}
