# Yet Another Go Heap (YAGH)

YAGH is a small Go package that provides its clients with the `IntMap[C]` data 
structure, a priority map that orders integers from 0 to N-1 by non-decreasing 
cost of type `C`. 

Data structure `IntMap[C]` is tailored for use cases where:

- The elements to be inserted in the map are known in advance and can be
  identified from 0 to N-1. 
- The map is meant to experience an arbitrarily large number of mutations 
  (including random access updates of its elements) over its lifetime.

It offers time complexity comparable to traditional heaps.

## O(1) Mallocs Complexity 

Data structure `IntMap[C]` is designed for use cases requiring numerous updates 
and mutations over the heap's lifetime. In particular, the map's constructor 
takes care of pre-allocating all the memory necessary to store up to N elements 
so that no further mallocs are required during the lifetime of the map. 

## Benchmark

Below are the benchmark results obtained by comparing YAGH's `IntMap` 
implementation with the standard [`container/heap`](https://pkg.go.dev/container/heap) 
implementation on a heapsort of 10000 random entries. Note that zero allocs are 
made in `BenchmarkIntMapSort`. 

```
BenchmarkIntMapSort-8   	     812	   1310105 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoHeapSort-8   	     706	   1680950 ns/op	  320000 B/op	   20000 allocs/op
```

See `yagh_benchmark_test.go` for more details. 