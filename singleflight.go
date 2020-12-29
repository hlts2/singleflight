package singleflight

import (
	"sync"
	"sync/atomic"
)

// Group --
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

// New --
func New() Group {
	return new(group)
}

// Do --
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
