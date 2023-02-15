# encode

simple custom encoding - playing around with golang, for fun and (very little) profit
 
```
ref https://pkg.go.dev/cmd/go/internal/test

go test --bench=. --count 2 --benchmem
goos: windows
goarch: amd64
pkg: github.com/SJMakin/encode
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkEncode/input_size_44-4          1742977               696.3 ns/op           160 B/op          2 allocs/op
BenchmarkEncode/input_size_44-4          1651494               636.1 ns/op           160 B/op          2 allocs/op
BenchmarkEncode/input_size_220-4          443614              2665 ns/op             736 B/op          2 allocs/op
BenchmarkEncode/input_size_220-4          547399              2832 ns/op             736 B/op          2 allocs/op
BenchmarkEncode/input_size_681-4          147622              7482 ns/op            2176 B/op          2 allocs/op
BenchmarkEncode/input_size_681-4          150399              7305 ns/op            2176 B/op          2 allocs/op
BenchmarkEncodeOptimized/input_size_44-4                 1696236               602.1 ns/op           328 B/op          6 allocs/op
BenchmarkEncodeOptimized/input_size_44-4                 1960818               581.8 ns/op           328 B/op          6 allocs/op
BenchmarkEncodeOptimized/input_size_220-4                 627112              1804 ns/op            1400 B/op          8 allocs/op
BenchmarkEncodeOptimized/input_size_220-4                 601832              1837 ns/op            1400 B/op          8 allocs/op
BenchmarkEncodeOptimized/input_size_681-4                 227094              4711 ns/op            4472 B/op         10 allocs/op
BenchmarkEncodeOptimized/input_size_681-4                 218846              4762 ns/op            4472 B/op         10 allocs/op
BenchmarkDecode/input_size_44-4                          1535334               852.8 ns/op           120 B/op          4 allocs/op
BenchmarkDecode/input_size_44-4                          1451750               831.1 ns/op           120 B/op          4 allocs/op
BenchmarkDecode/input_size_220-4                          325080              3996 ns/op             504 B/op          6 allocs/op
BenchmarkDecode/input_size_220-4                          303199              3408 ns/op             504 B/op          6 allocs/op
BenchmarkDecode/input_size_681-4                          119646              9108 ns/op            2040 B/op          8 allocs/op
BenchmarkDecode/input_size_681-4                          125892              9326 ns/op            2040 B/op          8 allocs/op
BenchmarkDecodeOptimized/input_size_44-4                 2266228               521.9 ns/op           168 B/op          5 allocs/op
BenchmarkDecodeOptimized/input_size_44-4                 2189098               564.0 ns/op           168 B/op          5 allocs/op
BenchmarkDecodeOptimized/input_size_220-4                 716668              1692 ns/op             728 B/op          7 allocs/op
BenchmarkDecodeOptimized/input_size_220-4                 623214              1660 ns/op             728 B/op          7 allocs/op
BenchmarkDecodeOptimized/input_size_681-4                 261568              4554 ns/op            2744 B/op          9 allocs/op
BenchmarkDecodeOptimized/input_size_681-4                 233595              4583 ns/op            2744 B/op          9 allocs/op
BenchmarkDecodeOptimizedSimplified/input_size_44-4               2401389               518.3 ns/op           168 B/op          5 allocs/op
BenchmarkDecodeOptimizedSimplified/input_size_44-4               2236362               529.4 ns/op           168 B/op          5 allocs/op
BenchmarkDecodeOptimizedSimplified/input_size_220-4              1000000              1683 ns/op             728 B/op          7 allocs/op
BenchmarkDecodeOptimizedSimplified/input_size_220-4               602091              1690 ns/op             728 B/op          7 allocs/op
BenchmarkDecodeOptimizedSimplified/input_size_681-4               279962              4499 ns/op            2744 B/op          9 allocs/op
BenchmarkDecodeOptimizedSimplified/input_size_681-4               279685              4570 ns/op            2744 B/op          9 allocs/op
PASS
ok      github.com/SJMakin/encode       45.025s
 
```
