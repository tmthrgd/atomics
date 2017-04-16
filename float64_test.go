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

func TestFloat64Default(t *testing.T) {
	var v Float64

	if v.Load() != 0 {
		t.Fatal("invalid default value")
	}
}

func TestNewFloat64(t *testing.T) {
	if NewFloat64(0) == nil {
		t.Fatal("NewFloat64 returned nil")
	}
}

func TestFloat64UnsafeRaw(t *testing.T) {
	var v Float64

	if v.UnsafeRaw() == nil {
		t.Fatal("UnsafeRaw returned nil")
	}
}

func TestFloat64Load(t *testing.T) {
	if err := quick.Check(func(v float64) bool {
		return NewFloat64(v).Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Store(t *testing.T) {
	if err := quick.Check(func(v float64) bool {
		var a Float64
		a.Store(v)
		return a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Swap(t *testing.T) {
	if err := quick.Check(func(old, new float64) bool {
		a := NewFloat64(old)
		return a.Swap(new) == old && a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64CompareAndSwap(t *testing.T) {
	if err := quick.Check(func(old, new float64) bool {
		a := NewFloat64(old)
		return !a.CompareAndSwap(-old, new) &&
			a.Load() == old &&
			a.CompareAndSwap(old, new) &&
			a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Add(t *testing.T) {
	if err := quick.Check(func(v, delta float64) bool {
		a := NewFloat64(v)
		v += delta
		return a.Add(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Increment(t *testing.T) {
	if err := quick.Check(func(v float64) bool {
		a := NewFloat64(v)
		v++
		return a.Increment() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Subtract(t *testing.T) {
	if err := quick.Check(func(v, delta float64) bool {
		a := NewFloat64(v)
		v -= delta
		return a.Subtract(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Decrement(t *testing.T) {
	if err := quick.Check(func(v float64) bool {
		a := NewFloat64(v)
		v--
		return a.Decrement() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64Reset(t *testing.T) {
	if err := quick.Check(func(v float64) bool {
		a := NewFloat64(v)
		return a.Reset() == v && a.Load() == 0
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFloat64String(t *testing.T) {
	if err := quick.Check(func(v float64) bool {
		return NewFloat64(v).String() == fmt.Sprint(v)
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNewFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewFloat64(0)
	}
}

func BenchmarkFloat64Load(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Load()
	}
}

func BenchmarkFloat64Store(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Store(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat64Swap(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Swap(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat64CompareAndSwap(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.CompareAndSwap(0, 0)
	}
}

func BenchmarkFloat64Add(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Add(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat64Increment(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Increment()
	}
}

func BenchmarkFloat64Subtract(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Subtract(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func BenchmarkFloat64Decrement(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Decrement()
	}
}

func BenchmarkFloat64Reset(b *testing.B) {
	var v Float64

	for n := 0; n < b.N; n++ {
		v.Reset()
	}
}
