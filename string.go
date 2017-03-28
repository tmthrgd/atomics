// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import "sync/atomic"

// String provides an atomic string.
type String struct {
	noCopy noCopy
	val    atomic.Value
}

// NewString returns an atomic string with a given value.
func NewString(val string) *String {
	var s String
	s.val.Store(val)
	return &s
}

// Load returns the value of the string.
func (s *String) Load() string {
	if v, ok := s.val.Load().(string); ok {
		return v
	}

	return ""
}

// Store sets the value of the string.
func (s *String) Store(val string) {
	s.val.Store(val)
}
