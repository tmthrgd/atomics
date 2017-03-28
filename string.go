// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"sync/atomic"
	"unsafe"
)

func pointerToString(val unsafe.Pointer) string {
	if val != nil {
		return *(*string)(val)
	}

	return ""
}

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
	return pointerToString(atomic.LoadPointer(&s.val))
}

func (s *String) store(val string) {
	atomic.StorePointer(&s.val, unsafe.Pointer(&val))
}

// Store sets the value of the string.
func (s *String) Store(val string) {
	if val == "" {
		atomic.StorePointer(&s.val, nil)
	} else {
		s.store(val)
	}
}

func (s *String) swap(new string) (old string) {
	return pointerToString(atomic.SwapPointer(&s.val, unsafe.Pointer(&new)))
}

// Swap sets the value of the string and returns the old value.
func (s *String) Swap(new string) (old string) {
	if new == "" {
		return s.Reset()
	}

	return s.swap(new)
}

// Reset sets the value of the string to "" and returns the old
// value.
func (s *String) Reset() (old string) {
	return pointerToString(atomic.SwapPointer(&s.val, nil))
}
