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
	val    *string
}

type stringPtr struct {
	val unsafe.Pointer
}

// NewString returns an atomic string with a given value.
func NewString(val string) *String {
	return &String{val: &val}
}

// Load returns the value of the string.
func (s *String) Load() string {
	p := (*stringPtr)(unsafe.Pointer(s))
	return pointerToString(atomic.LoadPointer(&p.val))
}

func (s *String) store(val string) {
	p := (*stringPtr)(unsafe.Pointer(s))
	atomic.StorePointer(&p.val, unsafe.Pointer(&val))
}

// Store sets the value of the string.
func (s *String) Store(val string) {
	if val == "" {
		p := (*stringPtr)(unsafe.Pointer(s))
		atomic.StorePointer(&p.val, nil)
	} else {
		s.store(val)
	}
}

func (s *String) swap(new string) (old string) {
	p := (*stringPtr)(unsafe.Pointer(s))
	return pointerToString(atomic.SwapPointer(&p.val, unsafe.Pointer(&new)))
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
	p := (*stringPtr)(unsafe.Pointer(s))
	return pointerToString(atomic.SwapPointer(&p.val, nil))
}
