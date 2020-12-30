# singleflight

[![GoDoc](https://godoc.org/github.com/hlts2/singleflight?status.svg)](http://godoc.org/github.com/hlts2/singleflight)

singleflight is simple golang singleflight library and provides a duplicate function call a suppression

## Requirement

Go 1.15

## Installing

```
go get github.com/hlts2/singleflight
```

## Example

```go
import (
    "log"
    "sync"
    
    "github.com/hlts2/singleflight"
)

var group = New()

func callAPI(key string) (v interface{}, err error, shared bool) {
    v, err, shared = group.Do(key, func() (resp interface{}, err error) {
        // send api request
        return
    })
    return
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            v, err, shared := callAPI("UserListAPI")
            if err == nil {
                log.Printf("response: %v", v)
                log.Printf("shared: %v", shared)
            }
        }()
    }
}
```

## Performance

The following represents [sync/singleflight](https://github.com/golang/sync/tree/master/singleflight) and [hlts2/singleflight](https://github.com/hlts2/singleflight) performance.
The `total` is the number of accesses and the `shared` is the number of shared results.

```
goos: linux
goarch: amd64
pkg: github.com/hlts2/singleflight
Benchmark_hlts2_singleflight-7   	12671460	        96.4 ns/op	      80 B/op	       2 allocs/op
--- BENCH: Benchmark_singleflight-7
    singleflight_bench_test.go:30: singleflight total: 1 shared: 0
    singleflight_bench_test.go:30: singleflight total: 100 shared: 99
    singleflight_bench_test.go:30: singleflight total: 10000 shared: 9999
    singleflight_bench_test.go:30: singleflight total: 1000000 shared: 999999
    singleflight_bench_test.go:30: singleflight total: 12671460 shared: 12671459
Benchmark_sync_singleflight-7    	13313968	       119 ns/op	      11 B/op	       0 allocs/op
--- BENCH: Benchmark_sync_singleflight-7
    singleflight_bench_test.go:53: singleflight total: 1 shared: 0
    singleflight_bench_test.go:53: singleflight total: 100 shared: 0
    singleflight_bench_test.go:53: singleflight total: 10000 shared: 3546
    singleflight_bench_test.go:53: singleflight total: 1000000 shared: 897228
    singleflight_bench_test.go:53: singleflight total: 13313968 shared: 11728991
```
