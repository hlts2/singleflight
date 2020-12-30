package singleflight

import (
	"sync"
	"sync/atomic"
)

// Group represents an interface that eliminates duplicate calls and executes functions.
type Group interface {
	Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)
}

type result struct {
	wg   sync.WaitGroup
	val  interface{}
	err  error
	dups uint64
}

type group struct {
	m sync.Map
}

// New creates Group instance.
func New() Group {
	return new(group)
}

// Do executes and returns the results of the given function, making sure that only one execution is in-flight for a given key at a time.
// If a duplicate comes in, the duplicate caller waits for the original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
func (g *group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	actual, loaded := g.m.LoadOrStore(key, new(result))
	result := actual.(*result)
	if loaded {
		result.wg.Wait()
		v, err = result.val, result.err
		return v, err, true
	}

	result.wg.Add(1)
	result.val, result.err = fn()
	result.wg.Done()

	return result.val, result.err, atomic.LoadUint64(&result.dups) > 0
}
