// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

// +build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	for _, typ := range []struct{ Type, Name, MathName string }{
		{"int32", "Int32", ""},
		{"int64", "Int64", ""},
		{"uint32", "Uint32", ""},
		{"uint64", "Uint64", ""},
		{"float32", "Float32", "Float32"},
		{"float64", "Float64", "Float64"},
	} {
		f, err := os.Create(typ.Type + "_test.go")
		if err != nil {
			log.Fatal(err)
		}

		if err = tmpl.Execute(f, typ); err != nil {
			log.Fatal(err)
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
	"fmt"
{{- if .MathName}}
	"math"
{{- end}}
	"testing"
	"testing/quick"
)

func Test{{.Name}}Default(t *testing.T) {
	var v {{.Name}}

	if v.Load() != 0 {
		t.Fatal("invalid default value")
	}
}

func TestNew{{.Name}}(t *testing.T) {
	if New{{.Name}}(0) == nil {
		t.Fatal("New{{.Name}} returned nil")
	}
}

func Test{{.Name}}Raw(t *testing.T) {
	var v {{.Name}}

	if v.Raw() == nil {
		t.Fatal("Raw returned nil")
	}

	if err := quick.Check(func(v {{.Type}}) bool {
{{- if .MathName}}
		return *New{{.Name}}(v).Raw() == math.{{.MathName}}bits(v)
{{- else}}
		return *New{{.Name}}(v).Raw() == v
{{- end}}
	}, nil); err != nil {
		t.Fatal(err)
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
		var a {{.Name}}
		a.Store(v)
		return a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Swap(t *testing.T) {
	if err := quick.Check(func(old, new {{.Type}}) bool {
		a := New{{.Name}}(old)
		return a.Swap(new) == old && a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}CompareAndSwap(t *testing.T) {
	if err := quick.Check(func(old, new {{.Type}}) bool {
		a := New{{.Name}}(old)
		return !a.CompareAndSwap(-old, new) &&
			a.Load() == old &&
			a.CompareAndSwap(old, new) &&
			a.Load() == new
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Add(t *testing.T) {
	if err := quick.Check(func(v, delta {{.Type}}) bool {
		a := New{{.Name}}(v)
		v += delta
		return a.Add(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Increment(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		a := New{{.Name}}(v)
		v++
		return a.Increment() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Subtract(t *testing.T) {
	if err := quick.Check(func(v, delta {{.Type}}) bool {
		a := New{{.Name}}(v)
		v -= delta
		return a.Subtract(delta) == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Decrement(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		a := New{{.Name}}(v)
		v--
		return a.Decrement() == v && a.Load() == v
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}Reset(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		a := New{{.Name}}(v)
		return a.Reset() == v && a.Load() == 0
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func Test{{.Name}}String(t *testing.T) {
	if err := quick.Check(func(v {{.Type}}) bool {
		return New{{.Name}}(v).String() == fmt.Sprint(v)
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNew{{.Name}}(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New{{.Name}}(0)
	}
}

func Benchmark{{.Name}}Load(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Load()
	}
}

func Benchmark{{.Name}}Store(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Store(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func Benchmark{{.Name}}Swap(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Swap(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func Benchmark{{.Name}}CompareAndSwap(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.CompareAndSwap(0, 0)
	}
}

func Benchmark{{.Name}}Add(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Add(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func Benchmark{{.Name}}Increment(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Increment()
	}
}

func Benchmark{{.Name}}Subtract(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Subtract(4) // RFC 1149.5 specifies 4 as the standard IEEE-vetted random number.
	}
}

func Benchmark{{.Name}}Decrement(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Decrement()
	}
}

func Benchmark{{.Name}}Reset(b *testing.B) {
	var v {{.Name}}

	for n := 0; n < b.N; n++ {
		v.Reset()
	}
}
`))
