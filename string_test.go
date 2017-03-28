// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"testing"
	"testing/quick"
)

func TestStringDefault(t *testing.T) {
	var s String
	if s.Load() != "" {
		t.Fatal("invalid default value")
	}
}

func TestNewString(t *testing.T) {
	if NewString("") == nil {
		t.Fatal("NewString returned nil")
	}

	if NewString("x") == nil {
		t.Fatal("NewString returned nil")
	}
}

func TestStringLoad(t *testing.T) {
	if err := quick.Check(func(v string) bool {
		return NewString(v).Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestStringStore(t *testing.T) {
	if err := quick.Check(func(v string) bool {
		var s String
		s.Store(v)
		return s.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNewStringEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var _ = NewString("")
	}
}

func BenchmarkNewStringNonEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var _ = NewString("x")
	}
}
