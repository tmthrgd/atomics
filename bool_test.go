// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import "testing"

func TestBoolDefault(t *testing.T) {
	var b Bool
	if b.Load() {
		t.Fatal("invalid default value")
	}
}

func TestNewBool(t *testing.T) {
	if NewBool(false) == nil {
		t.Fatal("NewBool returned nil")
	}
}

func TestBoolUnsafeRaw(t *testing.T) {
	var b Bool
	if b.UnsafeRaw() == nil {
		t.Fatal("UnsafeRaw returned nil")
	}
}

func TestBoolLoad(t *testing.T) {
	if NewBool(false).Load() {
		t.Fatal("Load failed for false")
	}

	if !NewBool(true).Load() {
		t.Fatal("Load failed for true")
	}
}

func TestBoolStore(t *testing.T) {
	var c Bool

	if c.Store(true); !c.Load() {
		t.Fatal("Store failed for true")
	}

	if c.Store(false); c.Load() {
		t.Fatal("Store failed for false")
	}
}

func TestBoolSwap(t *testing.T) {
	if c := NewBool(false); c.Swap(true) || !c.Load() {
		t.Fatal("Swap failed for true")
	}

	if c := NewBool(true); !c.Swap(false) || c.Load() {
		t.Fatal("Swap failed for false")
	}
}

func TestBoolCompareAndSwap(t *testing.T) {
	var c Bool

	if c.CompareAndSwap(true, true) || c.Load() {
		t.Fatal("CompareAndSwap should not have swapped")
	}

	if !c.CompareAndSwap(false, true) || !c.Load() {
		t.Fatal("CompareAndSwap should have swapped")
	}
}

func TestBoolReset(t *testing.T) {
	if c := NewBool(true); !c.Reset() || c.Load() {
		t.Fatal("Reset failed")
	}
}
