// Code generated by go run generate.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import (
	"sync/atomic"

	"github.com/golang/sync/syncmap"
)

// Int32 provides a map of atomic counters of type int32.
type Int32 struct {
	m syncmap.Map // map[interface{}]*int32
}

// UnsafeLoad returns a pointer to the counter key.
//
// It is only safe to access the return value with
// methods from the sync/atomic package. It must
// not be manually dereferenced.
func (c *Int32) UnsafeLoad(key interface{}) *int32 {
	v, _ := c.m.LoadOrStore(key, new(int32))
	return v.(*int32)
}

// Load returns the value of the counter key.
func (c *Int32) Load(key interface{}) (val int32) {
	return atomic.LoadInt32(c.UnsafeLoad(key))
}

// Store sets the value of the counter key.
func (c *Int32) Store(key interface{}, val int32) {
	atomic.StoreInt32(c.UnsafeLoad(key), val)
}

// Swap sets the value of the counter key and returns the
// old value.
func (c *Int32) Swap(key interface{}, new int32) (old int32) {
	return atomic.SwapInt32(c.UnsafeLoad(key), new)
}

// CompareAndSwap sets the value of the counter key to new
// but only if it currently has the value old.
func (c *Int32) CompareAndSwap(key interface{}, old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(c.UnsafeLoad(key), old, new)
}

// Add adds delta to the counter key.
func (c *Int32) Add(key interface{}, delta int32) (new int32) {
	return atomic.AddInt32(c.UnsafeLoad(key), delta)
}

// Increment is a wrapper for Add(key, 1).
func (c *Int32) Increment(key interface{}) (new int32) {
	return c.Add(key, 1)
}

// Subtract is a wrapper for Add(key, -delta)
func (c *Int32) Subtract(key interface{}, delta int32) (new int32) {
	return c.Add(key, -delta)
}

// Decrement is a wrapper for Add(key, -1).
func (c *Int32) Decrement(key interface{}) (new int32) {
	return c.Add(key, -1)
}

// Reset is a wrapper for Swap(key, 0).
func (c *Int32) Reset(key interface{}) (old int32) {
	return c.Swap(key, 0)
}

// Delete removes the counter key from the map.
func (c *Int32) Delete(key interface{}) {
	c.m.Delete(key)
}

// Keys returns the list of all counters.
func (c *Int32) Keys() []interface{} {
	var keys []interface{}
	c.m.Range(func(key, val interface{}) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

// UnsafeRange calls f with a pointer to each
// counter.
//
// It is only safe to access val with methods from
// the sync/atomic package. It must not be manually
// dereferenced.
func (c *Int32) UnsafeRange(f func(key interface{}, val *int32) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, val.(*int32))
	})
}

// RangeKeys calls f with the key of each counter.
func (c *Int32) RangeKeys(f func(key interface{}) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key)
	})
}

// RangeLoad calls f with the value of each counter.
func (c *Int32) RangeLoad(f func(key interface{}, val int32) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, atomic.LoadInt32(val.(*int32)))
	})
}

// RangeStore sets each counter to the return value of f.
func (c *Int32) RangeStore(f func(key interface{}) (val int32, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		v, ok := f(key)
		atomic.StoreInt32(val.(*int32), v)
		return ok
	})
}

// RangeAdd adds the return value of f to each counter.
func (c *Int32) RangeAdd(f func(key interface{}) (delta int32, ok bool)) {
	c.m.Range(func(key, val interface{}) bool {
		delta, ok := f(key)
		atomic.AddInt32(val.(*int32), delta)
		return ok
	})
}

// RangeReset resets each counter and calls f with the
// old value.
func (c *Int32) RangeReset(f func(key interface{}, old int32) bool) {
	c.m.Range(func(key, val interface{}) bool {
		return f(key, atomic.SwapInt32(val.(*int32), 0))
	})
}
