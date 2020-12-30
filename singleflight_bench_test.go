package singleflight

import (
	"sync/atomic"
	"testing"

	"golang.org/x/sync/singleflight"
)

func Benchmark_singleflight(b *testing.B) {
	sf := New()

	var (
		cnt   int64
		total int64
	)

	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&total, 1)
		_, _, shared := sf.Do("key", func() (interface{}, error) {
			return nil, nil
		})
		if shared {
			atomic.AddInt64(&cnt, 1)
		}
	}

	b.Logf("singleflight total: %d shared: %d", atomic.LoadInt64(&total), atomic.LoadInt64(&cnt))
}

func Benchmark_sync_singleflight(b *testing.B) {
	sf := new(singleflight.Group)

	var (
		cnt   int64
		total int64
	)

	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&total, 1)
		_, _, shared := sf.Do("key", func() (interface{}, error) {
			return nil, nil
		})
		if shared {
			atomic.AddInt64(&cnt, 1)
		}
	}

	b.Logf("singleflight total: %d shared: %d", atomic.LoadInt64(&total), atomic.LoadInt64(&cnt))
}
