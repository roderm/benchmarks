# SQL Benchmarking
To load n+1 problematic data from a database, there are a lot of diffrent solutions. I wanted to start a selector-generator based on protobuf defintions and decided to first get some benchmarks on the diffrent solutions.

## usage
You can directly use docker-compose (`docker-compose up -d`) to start and afterwards connect to the benchmark container with `docker-compose exec benchmark bash` and go to `/golang/src/github.com/roderm/benchmark/sql`:
- Setup the database and load it with data: `go run setup/main.go` (you can modify the amount of companies/employes/product in here, currentls really slow)
- Test the dataloader: `go test -bench=Dataloader`
- Test the JSON (currently no working): `go test -bench=JSON`

## dataloader
The dataloader executes a query on each table and maps it to the golang-struct

## jsonagg
A lot of databases out there have some functions to return a query-result in json. with this method a single query can hold all data and can directly be mapped to a golang struct. (no circles)

PROBLEM: empty dates are store in cockroachdb as "0001-01-01T00:00:00" and can't be read from json.Unmarshal
```
Scan error on column index 0, name "company": parsing time ""0001-01-01T00:00:00"" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "Z07:00"
```

# Results: 
## 20 companies, 200 employees, 100 products
```
goos: linux
goarch: amd64
pkg: github.com/roderm/benchmarks/sql
BenchmarkJSON-4                 1000000000               0.222 ns/op
BenchmarkDataloader-4           1000000000               0.109 ns/op
PASS
ok      github.com/roderm/benchmarks/sql        4.264s
```