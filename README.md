# generic benchmarks
```go test -bench=. -benchtime=100000000x -count=5 -benchmem``` \
goos: darwin \
goarch: amd64 \
pkg: genericBench \
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz

BenchmarkSum/SumSwitchInt-16            100000000               16.98 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchInt-16            100000000               11.54 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchInt-16            100000000               11.61 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchInt-16            100000000               11.41 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchInt-16            100000000               11.26 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumInts-16                 100000000               10.48 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumInts-16                 100000000               10.35 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumInts-16                 100000000               10.44 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumInts-16                 100000000               10.82 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumInts-16                 100000000               11.14 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericInts-16          100000000               10.36 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericInts-16          100000000               10.36 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericInts-16          100000000               10.31 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericInts-16          100000000               10.41 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericInts-16          100000000               10.37 ns/op            8 B/op          1 allocs/op
 \
 \
BenchmarkSum/SumSwitchFloat32-16        100000000                9.960 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat32-16        100000000                8.835 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat32-16        100000000                8.756 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat32-16        100000000                8.799 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat32-16        100000000                8.948 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumFloats32-16             100000000                9.633 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumFloats32-16             100000000               11.60 ns/op            4 B/op          1 allocs/op \
BenchmarkSum/SumFloats32-16             100000000               13.21 ns/op            4 B/op          1 allocs/op \
BenchmarkSum/SumFloats32-16             100000000                9.515 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumFloats32-16             100000000                9.725 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats32-16      100000000                8.881 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats32-16      100000000                9.252 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats32-16      100000000                9.089 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats32-16      100000000                9.456 ns/op           4 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats32-16      100000000                9.977 ns/op           4 B/op          1 allocs/op \
 \
 \
BenchmarkSum/SumSwitchFloat64-16        100000000               15.34 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat64-16        100000000               12.09 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat64-16        100000000               13.19 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat64-16        100000000               12.35 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumSwitchFloat64-16        100000000               14.09 ns/op            8 B/op          1 allocs/op


BenchmarkSum/SumFloats64-16             100000000               11.56 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumFloats64-16             100000000               11.19 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumFloats64-16             100000000               11.32 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumFloats64-16             100000000               10.54 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumFloats64-16             100000000               10.60 ns/op            8 B/op          1 allocs/op

BenchmarkSum/SumGenericFloats64-16      100000000               11.01 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats64-16      100000000               11.12 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats64-16      100000000               10.99 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats64-16      100000000               11.65 ns/op            8 B/op          1 allocs/op \
BenchmarkSum/SumGenericFloats64-16      100000000               11.46 ns/op            8 B/op          1 allocs/op \
PASS \
ok      genericBench    50.901s \

