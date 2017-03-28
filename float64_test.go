// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import "testing"

func TestFloat64UnsafeLoad(t *testing.T) {
	var c Float64

	if c.unsafeLoad(0) == nil {
		t.Fatal("UnsafeLoad returned nil")
	}
}

func TestFloat64Load(t *testing.T) {
	var c Float64

	if c.Load(0) != 0 {
		t.Fatal("Load returned non-zero for new key")
	}

	c.Store(0, 1)

	if c.Load(0) != 1 {
		t.Fatal("Load did not return value for key")
	}
}

func TestFloat64Store(t *testing.T) {
	var c Float64
	c.Store(0, 1)

	if c.Load(0) != 1 {
		t.Fatal("Store did not set value for key")
	}
}

func TestFloat64Swap(t *testing.T) {
	var c Float64
	c.Store(0, 1)

	if c.Swap(0, 2) != 1 {
		t.Fatal("Swap did not return old value for key")
	}

	if c.Load(0) != 2 {
		t.Fatal("Swap did not set value for key")
	}
}

func TestFloat64CompareAndSwap(t *testing.T) {
	var c Float64
	c.Store(0, 1)

	if c.CompareAndSwap(0, 0, 2) {
		t.Fatal("CompareAndSwap did not compare")
	}

	if !c.CompareAndSwap(0, 1, 2) {
		t.Fatal("CompareAndSwap compare failed")
	}

	if c.Load(0) != 2 {
		t.Fatal("CompareAndSwap did not swap")
	}
}

func TestFloat64Reset(t *testing.T) {
	var c Float64
	c.Store(0, 2)

	if c.Reset(0) != 2 {
		t.Fatal("Reset returned wrong value")
	}

	if c.Load(0) != 0 {
		t.Fatal("Reset failed")
	}
}

func TestFloat64Delete(t *testing.T) {
	var c Float64
	c.Store(0, 2)
	c.Delete(0)

	if c.Load(0) != 0 {
		t.Fatal("Delete failed")
	}
}

func TestFloat64Keys(t *testing.T) {
	var c Float64
	c.Store(0, 1)
	c.Store(1, 2)

	keys := c.Keys()

	if len(keys) != 2 {
		t.Fatal("Keys returned wrong number of keys")
	}

	if keys[0] == keys[1] {
		t.Fatal("Keys contains duplicate key")
	}

	if keys[0] != 0 && keys[1] != 0 {
		t.Fatal("Keys missing key 0")
	}

	if keys[0] != 1 && keys[1] != 1 {
		t.Fatal("Keys missing key 1")
	}
}
