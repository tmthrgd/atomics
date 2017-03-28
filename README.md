# counters

[![GoDoc](https://godoc.org/github.com/tmthrgd/counters?status.svg)](https://godoc.org/github.com/tmthrgd/counters)
[![Build Status](https://travis-ci.org/tmthrgd/counters.svg?branch=master)](https://travis-ci.org/tmthrgd/counters)

Package counters implements efficient atomic counters for statistics collection.

They rely on the [sync/atomic](https://golang.org/pkg/sync/atomic/) package and the experimental
[github.com/golang/sync/syncmap](https://godoc.org/github.com/golang/sync/syncmap) package *which
may change without warning*. The syncmap package is a proposed addition to the sync package in the
standard library. ([golang/go#18177](https://golang.org/issue/18177))

See the [documentation](https://godoc.org/github.com/tmthrgd/counters?status.svg) for use. This
package should be self-explanatory.