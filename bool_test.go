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

func TestBoolRaw(t *testing.T) {
	var b Bool

	if b.Raw() == nil {
		t.Fatal("Raw returned nil")
	}

	if *NewBool(false).Raw() != 0 {
		t.Fatal("Raw failed for false")
	}

	if *NewBool(true).Raw() != 1 {
		t.Fatal("Raw failed for true")
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
	var b Bool

	if b.Store(true); !b.Load() {
		t.Fatal("Store failed for true")
	}

	if b.Store(false); b.Load() {
		t.Fatal("Store failed for false")
	}
}

func TestBoolSwap(t *testing.T) {
	if b := NewBool(false); b.Swap(true) || !b.Load() {
		t.Fatal("Swap failed for true")
	}

	if b := NewBool(true); !b.Swap(false) || b.Load() {
		t.Fatal("Swap failed for false")
	}
}

func TestBoolCompareAndSwap(t *testing.T) {
	var b Bool

	if b.CompareAndSwap(true, true) || b.Load() {
		t.Fatal("CompareAndSwap should not have swapped")
	}

	if !b.CompareAndSwap(false, true) || !b.Load() {
		t.Fatal("CompareAndSwap should have swapped")
	}
}

func TestBoolSet(t *testing.T) {
	var b Bool

	if b.Set() || !b.Load() {
		t.Fatal("Set failed")
	}
}

func TestBoolReset(t *testing.T) {
	if b := NewBool(true); !b.Reset() || b.Load() {
		t.Fatal("Reset failed")
	}
}

func TestBoolString(t *testing.T) {
	if NewBool(false).String() != "false" {
		t.Fatal("String failed for false")
	}

	if NewBool(true).String() != "true" {
		t.Fatal("String failed for true")
	}
}

func BenchmarkNewBool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewBool(true)
	}
}

func BenchmarkBoolLoad(b *testing.B) {
	var v Bool

	for n := 0; n < b.N; n++ {
		v.Load()
	}
}

func BenchmarkBoolStore(b *testing.B) {
	var v Bool

	for n := 0; n < b.N; n++ {
		v.Store(true)
	}
}

func BenchmarkBoolSwap(b *testing.B) {
	var v Bool

	for n := 0; n < b.N; n++ {
		v.Swap(true)
	}
}

func BenchmarkBoolCompareAndSwap(b *testing.B) {
	var v Bool

	for n := 0; n < b.N; n++ {
		v.CompareAndSwap(true, true)
	}
}
