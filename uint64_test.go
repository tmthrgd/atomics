// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import "testing"

func TestUnsafeLoad(t *testing.T) {
	var c Uint64

	if c.UnsafeLoad(0) == nil {
		t.Fatal("UnsafeLoad returned nil")
	}
}

func TestLoad(t *testing.T) {
	var c Uint64

	if c.Load(0) != 0 {
		t.Fatal("Load returned non-zero for new key")
	}

	c.Store(0, 1)

	if c.Load(0) != 1 {
		t.Fatal("Load did not return value for key")
	}
}

func TestStore(t *testing.T) {
	var c Uint64
	c.Store(0, 1)

	if c.Load(0) != 1 {
		t.Fatal("Store did not set value for key")
	}
}

func TestSwap(t *testing.T) {
	var c Uint64
	c.Store(0, 1)

	if c.Swap(0, 2) != 1 {
		t.Fatal("Swap did not return old value for key")
	}

	if c.Load(0) != 2 {
		t.Fatal("Swap did not set value for key")
	}
}

func TestCompareAndSwap(t *testing.T) {
	var c Uint64
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

func TestAdd(t *testing.T) {
	var c Uint64
	c.Add(0, 2)

	if c.Load(0) != 2 {
		t.Fatal("Add failed")
	}
}

func TestIncrement(t *testing.T) {
	var c Uint64
	c.Increment(0)

	if c.Load(0) != 1 {
		t.Fatal("Increment failed")
	}
}

func TestSubtract(t *testing.T) {
	var c Uint64
	c.Store(0, 2)
	c.Subtract(0, 1)

	if c.Load(0) != 1 {
		t.Fatal("Subtract failed")
	}

	c.Subtract(0, 2)

	if c.Load(0) != ^uint64(0) {
		t.Fatal("Subtract failed")
	}
}

func TestDecrement(t *testing.T) {
	var c Uint64
	c.Store(0, 2)
	c.Decrement(0)

	if c.Load(0) != 1 {
		t.Fatal("Decrement failed")
	}

	c.Decrement(0)
	c.Decrement(0)

	if c.Load(0) != ^uint64(0) {
		t.Fatal("Decrement failed")
	}
}

func TestReset(t *testing.T) {
	var c Uint64
	c.Store(0, 2)

	if c.Reset(0) != 2 {
		t.Fatal("Reset returned wrong value")
	}

	if c.Load(0) != 0 {
		t.Fatal("Reset failed")
	}
}

func TestDelete(t *testing.T) {
	var c Uint64
	c.Store(0, 2)
	c.Delete(0)

	if c.Load(0) != 0 {
		t.Fatal("Delete failed")
	}
}

func TestKeys(t *testing.T) {
	var c Uint64
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
