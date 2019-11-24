# timeseries

threadsafe timeseries package for gophers



[![GoDoc](https://godoc.org/github.com/axamon/timeseries?status.svg)](https://godoc.org/github.com/axamon/timeseries)
[![Build Status](https://travis-ci.org/axamon/timeseries.svg?branch=master)](https://travis-ci.org/axamon/timeseries)
[![codecov](https://codecov.io/gh/axamon/timeseries/branch/master/graph/badge.svg)](https://codecov.io/gh/axamon/timeseries)

Benchmark:


    goos: windows
    goarch: amd64
    pkg: github.com/axamon/timeseries
    BenchmarkAddNewPoint-4   	 5000000	       364 ns/op	      70 B/op	       0 allocs/op
    PASS
    ok  	github.com/axamon/timeseries	2.443s
    Success: Benchmarks passed.