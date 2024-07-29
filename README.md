# Futils

A collection of common functions but with better performance, fewer allocations and fewer dependencies.

## Benchmarks

Environment: 11th Gen Intel(R) Core(TM) i7-11390H @ 3.40GHz

```bash
Benchmark_ToUpper/futils-8                              18346155                66.31 ns/op           80 B/op           1 allocs/op
Benchmark_ToUpper/IfToUpper-Upper-8                     26770777                52.41 ns/op           0 B/op            0 allocs/op
Benchmark_ToUpper/IfToUpper-Mixed-8                     18458344                72.57 ns/op           80 B/op           1 allocs/op
Benchmark_ToUpper/default-8                             4597606                 273.3 ns/op           80 B/op           1 allocs/op
Benchmark_ToLower/futils-8                              16207501                73.19 ns/op           80 B/op           1 allocs/op
Benchmark_ToLower/IfToLower-Lower-8                     25445340                48.53 ns/op           0 B/op            0 allocs/op
Benchmark_ToLower/IfToLower-Mixed-8                     17034793                73.19 ns/op           80 B/op           1 allocs/op
Benchmark_ToLower/default-8                             5193901                 221.7 ns/op           80 B/op           1 allocs/op
Benchmark_CalculateTimestamp/futils-8                   1000000000              0.3110 ns/op          0 B/op            0 allocs/op
Benchmark_CalculateTimestamp/default-8                  38381926                35.75 ns/op           0 B/op            0 allocs/op
Benchmark_CalculateTimestamp/futils_asserted-8          2281803                 508.4 ns/op           0 B/op            0 allocs/op
Benchmark_CalculateTimestamp/default_asserted-8         2345218                 535.7 ns/op           0 B/op            0 allocs/op
Benchmark_UUID/futils-8                                 23780588                48.58 ns/op           48 B/op           1 allocs/op
Benchmark_UUID/default-8                                3420697                 339.0 ns/op           168 B/op          6 allocs/op
```