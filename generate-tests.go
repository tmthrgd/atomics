// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

// +build ignore

package main

import (
	"os"
	"text/template"
)

func main() {
	for _, typ := range []struct{ Type, Name string }{
		{"int32", "Int32"},
		{"int64", "Int64"},
		{"uint32", "Uint32"},
		{"uint64", "Uint64"},
		{"float32", "Float32"},
		{"float64", "Float64"},
	} {
		f, err := os.Create(typ.Type + "_test.go")
		if err != nil {
			panic(err)
		}

		if err = tmpl.Execute(f, typ); err != nil {
			panic(err)
		}

		f.Close()
	}
}

var tmpl = template.Must(template.New("test").Parse(`// Code generated by go run generate-tests.go.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

package atomics

import (
	"testing"
	"testing/quick"
)

func TestNew{{.Name}}(t *testing.T) {
	if New{{.Name}}(0) == nil {
		t.Fatal("New{{.Name}} returned nil")
	}
}

func Test{{.Name}}UnsafeRaw(t *testing.T) {
	var c {{.Name}}
	if c.UnsafeRaw() == nil {
		t.Fatal("UnsafeRaw returned nil")
	}
}

func Test{{.Name}}Load(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		return New{{.Name}}(v).Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Store(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		var c {{.Name}}
		c.Store(v)
		return c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Swap(t *testing.T) {
	if err := quick.Check(func(old, new {{.Type}}) bool {
		c := New{{.Name}}(old)
		return c.Swap(new) == old && c.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}CompareAndSwap(t *testing.T) {
	if err := quick.Check(func(old, new {{.Type}}) bool {
		c := New{{.Name}}(old)
		return !c.CompareAndSwap(-old, new) &&
			c.Load() == old &&
			c.CompareAndSwap(old, new) &&
			c.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Add(t *testing.T) {
	if err := quick.Check(func(v, delta {{.Type}}) bool {
		c := New{{.Name}}(v)
		v += delta
		return c.Add(delta) == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Increment(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		c := New{{.Name}}(v)
		v++
		return c.Increment() == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Subtract(t *testing.T) {
	if err := quick.Check(func(v, delta {{.Type}}) bool {
		c := New{{.Name}}(v)
		v -= delta
		return c.Subtract(delta) == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Decrement(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		c := New{{.Name}}(v)
		v--
		return c.Decrement() == v && c.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Reset(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		c := New{{.Name}}(v)
		return c.Reset() == v && c.Load() == 0
	}, nil); err != nil {
		t.Fatal(err)
	}
}
`))
