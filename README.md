# Yet Another Go Heap (YAGH)

[![Go Reference](https://pkg.go.dev/badge/github.com/rhartert/yagh.svg)](https://pkg.go.dev/github.com/rhartert/yagh)

YAGH is a small Go package that provides its clients with the `IntMap[C]` data 
structure, a priority map that orders integers from 0 to N-1 by non-decreasing 
cost of type `C`. 

Data structure `IntMap[C]` is tailored for use cases where:

- The elements to be inserted in the map are known in advance and can be
  identified from 0 to N-1. 
- The map is meant to experience an arbitrarily large number of mutations 
  (including random access updates of its elements) over its lifetime.

Its operations have comparable time complexities to traditional heaps.

## Garbage Friendly 🌱

`IntMap[C]` is designed for scenarios where frequent updates and mutations are
expected. The map's constructor pre-allocates all necessary memory to store up
to `N` elements, thus eliminating the need for further memory allocations during 
its lifetime. This design ensures consistent speed performance, as it minimizes
variations caused by memory allocation or garbage collection.

## Benchmark

We've compared the performance of YAGH's `IntMap` to Go's standard 
[`container/heap`](https://pkg.go.dev/container/heap) on multiple heapsorts of 
10000 random entries (see benchmark results below). 

To run the benchmark, simply run the following command from the root of this
repository:

```bash
go test -benchmem -bench . 
```

This should output something similar to this:

```
goos: darwin
goarch: arm64
pkg: github.com/rhartert/yagh
BenchmarkIntMapSort-8  1063  1092887 ns/op  0 B/op       0 allocs/op
BenchmarkGoHeapSort-8   672  1802387 ns/op  320001 B/op  20000 allocs/op
```

On average, YAGH achieves a 1.64x [speed-up] compared to Go's standard heap. 
Also, note that zero allocs are made in `BenchmarkIntMapSort` (i.e. the heapsort 
using YAGH). 

[speed-up]: https://en.wikipedia.org/wiki/Speedup