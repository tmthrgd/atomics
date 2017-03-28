// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"sync/atomic"
	"unsafe"
)

// String provides an atomic string.
type String struct {
	noCopy noCopy
	val    unsafe.Pointer
}

// NewString returns an atomic string with a given value.
func NewString(val string) *String {
	return &String{
		val: unsafe.Pointer(&val),
	}
}

// Load returns the value of the string.
func (s *String) Load() string {
	if val := atomic.LoadPointer(&s.val); val != nil {
		return *(*string)(val)
	}

	return ""
}

// Store sets the value of the string.
func (s *String) Store(val string) {
	atomic.StorePointer(&s.val, unsafe.Pointer(&val))
}
