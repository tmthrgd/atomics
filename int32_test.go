// Code generated by go run generate-tests.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestInt32Default(t *testing.T) {
	var v Int32

	if v.Load() != 0 {
		t.Fatal("invalid default value")
	}
}

func TestNewInt32(t *testing.T) {
	if NewInt32(0) == nil {
		t.Fatal("NewInt32 returned nil")
	}
}

func TestInt32Raw(t *testing.T) {
	var v Int32

	if v.Raw() == nil {
		t.Fatal("Raw returned nil")
	}

	if err := quick.Check(func(v int32) bool {
		return *NewInt32(v).Raw() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Load(t *testing.T) {
	if err := quick.Check(func(v int32) bool {
		return NewInt32(v).Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Store(t *testing.T) {
	if err := quick.Check(func(v int32) bool {
		var a Int32
		a.Store(v)
		return a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Swap(t *testing.T) {
	if err := quick.Check(func(old, new int32) bool {
		a := NewInt32(old)
		return a.Swap(new) == old && a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32CompareAndSwap(t *testing.T) {
	if err := quick.Check(func(old, new int32) bool {
		a := NewInt32(old)
		return !a.CompareAndSwap(-old, new) &&
			a.Load() == old &&
			a.CompareAndSwap(old, new) &&
			a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Add(t *testing.T) {
	if err := quick.Check(func(v, delta int32) bool {
		a := NewInt32(v)
		v += delta
		return a.Add(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Increment(t *testing.T) {
	if err := quick.Check(func(v int32) bool {
		a := NewInt32(v)
		v++
		return a.Increment() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Subtract(t *testing.T) {
	if err := quick.Check(func(v, delta int32) bool {
		a := NewInt32(v)
		v -= delta
		return a.Subtract(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Decrement(t *testing.T) {
	if err := quick.Check(func(v int32) bool {
		a := NewInt32(v)
		v--
		return a.Decrement() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32Reset(t *testing.T) {
	if err := quick.Check(func(v int32) bool {
		a := NewInt32(v)
		return a.Reset() == v && a.Load() == 0
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInt32String(t *testing.T) {
	if err := quick.Check(func(v int32) bool {
		return NewInt32(v).String() == fmt.Sprint(v)
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNewInt32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewInt32(0)
	}
}

func BenchmarkInt32Load(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Load()
	}
}

func BenchmarkInt32Store(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Store(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkInt32Swap(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Swap(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkInt32CompareAndSwap(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.CompareAndSwap(0, 0)
	}
}

func BenchmarkInt32Add(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Add(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkInt32Increment(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Increment()
	}
}

func BenchmarkInt32Subtract(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Subtract(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkInt32Decrement(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Decrement()
	}
}

func BenchmarkInt32Reset(b *testing.B) {
	var v Int32

	for n := 0; n < b.N; n++ {
		v.Reset()
	}
}
