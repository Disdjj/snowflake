# snowflake

support IDC / traditional snowflake mode

Rent Mode has a better performance than Spin Mode and not afraid of time drift.

```shell
BenchmarkParallel10
BenchmarkParallel10-19                  22684210                82.00 ns/op
BenchmarkParallel100
BenchmarkParallel100-19                 20711701               119.8 ns/op
BenchmarkParallel1000
BenchmarkParallel1000-19                15499429               121.7 ns/op
BenchmarkParallel10000
BenchmarkParallel10000-19               10542180               112.0 ns/op
BenchmarkParallel100000
BenchmarkParallel100000-19              17447278                71.31 ns/op
BenchmarkParallelSpin10
BenchmarkParallelSpin10-19               4744905               256.6 ns/op
BenchmarkParallelSpin100
BenchmarkParallelSpin100-19              4787672               257.5 ns/op
BenchmarkParallelSpin1000
BenchmarkParallelSpin1000-19             4763870               254.6 ns/op
BenchmarkParallelSpin10000
BenchmarkParallelSpin10000-19            4614619               254.0 ns/op
BenchmarkParallelSpin100000
BenchmarkParallelSpin100000-19           4710369               254.4 ns/op
```
