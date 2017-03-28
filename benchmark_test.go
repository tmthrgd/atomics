// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import "testing"

func BenchmarkUnsafeLoadNew(b *testing.B) {
	var c Uint64

	for n := 0; n < b.N; n++ {
		c.UnsafeLoad(n)
	}
}

func BenchmarkUnsafeLoadExisting(b *testing.B) {
	var c Uint64
	c.Store(0, 0)

	for n := 0; n < b.N; n++ {
		c.UnsafeLoad(0)
	}
}
