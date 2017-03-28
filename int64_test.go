// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package counters

import "testing"

func TestSignedDecrement(t *testing.T) {
	var c Int64
	c.Store(0, 2)
	c.Decrement(0)

	if c.Reset(0) != 1 {
		t.Fatal("Decrement failed")
	}

	c.Decrement(0)

	if c.Load(0) != -1 {
		t.Fatal("Decrement failed")
	}
}
