# Pubsub
In go we have channels for multithreaded communication between go-routines. In my first benchmark-test I'm wondering how fast they are compared to callbacks.

## plain
without any go-routines

## routines
start for every pub-message a new routine

## workers
add a bund of workers to not explode goroutines


## random results:
```
goos: darwin
goarch: amd64
pkg: github.com/roderm/benchmarks/pubsub
BenchmarkTopicChannel-4                               30          63507643 ns/op
BenchmarkTopicCallback-4                             500           2941913 ns/op
BenchmarkSubsChannel-4                              3000            489339 ns/op
BenchmarkSubsCallback-4                            50000             29364 ns/op
BenchmarkMsgsChannel-4                               300           5446927 ns/op
BenchmarkMsgsCallback-4                             5000            294184 ns/op
BenchmarkTopicChannelPool-4                           20          74223541 ns/op
BenchmarkTopicCallbackPool-4                          30          40068874 ns/op
BenchmarkSubsChannelPool-4                          2000           1090099 ns/op
BenchmarkSubsCallbackPool-4                         3000            493873 ns/op
BenchmarkMsgsChannelPool-4                           100          11359228 ns/op
BenchmarkMsgsCallbackPool-4                          300           5091090 ns/op
BenchmarkTopicChannelRoutine-4                        20         141445163 ns/op
BenchmarkTopicCallbackRoutine-4                       50          57290107 ns/op
BenchmarkSubsChannelRoutine-4                       2000            813597 ns/op
BenchmarkSubsCallbackRoutine-4                      3000            410324 ns/op
BenchmarkMsgsChannelRoutine-4                        200           8279642 ns/op
BenchmarkMsgsCallbackRoutineRoutine-4                300           4221270 ns/op
PASS
ok      github.com/roderm/benchmarks/pubsub     39.504s
```