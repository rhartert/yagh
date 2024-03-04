package yagh

import (
	"container/heap"
	"math/rand"
	"testing"
)

const N = 10_000

type entry struct {
	elem int
	cost float64
}

// goHeap is a slice of entries that implements heap.Interface, and is used
// as baseline to evaluate the performance of IntMap.
type goHeap []entry

func (h goHeap) Len() int { return len(h) }

func (h goHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func (h goHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *goHeap) Push(x any) {
	*h = append(*h, x.(entry))
}

func (h *goHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func makeRandomEntries(n int) []entry {
	rng := rand.New(rand.NewSource(42))
	entries := make([]entry, 0, n)
	for i := 0; i < N; i++ {
		entries = append(entries, entry{
			elem: i,
			cost: rng.Float64(),
		})
	}
	return entries
}

func BenchmarkIntMapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		entries := makeRandomEntries(N)
		m := New[float64](N)
		b.StartTimer()

		for _, e := range entries {
			m.Put(e.elem, e.cost)
		}
		for m.Size() != 0 {
			m.Pop()
		}
	}
}

func BenchmarkGoHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		entries := makeRandomEntries(N)
		h := goHeap(make([]entry, 0, N))
		heap.Init(&h)
		b.StartTimer()

		for _, e := range entries {
			heap.Push(&h, e)
		}
		for h.Len() != 0 {
			heap.Pop(&h)
		}
	}
}
