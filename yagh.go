// Package yagh exposes the IntMap[C] data structure, a priority map that orders
// integers from 0 to N-1 by non-decreasing cost of type `C`.
//
// Data structure IntMap is tailored for use cases where:
//   - The elements to be inserted in the map are known in advance and can be
//     identified from 0 to N-1.
//   - The map is meant to experience an arbitrarily large number of mutations
//     (including random access updates of its elements) over its lifetime.
//
// In particular, this implementation aims to minimize memory allocations and
// avoid creating objects that would ultimately have to garbage collected.
package yagh

import (
	"cmp"
	"fmt"
	"strings"
)

type Entry[C cmp.Ordered] struct {
	Elem int
	Cost C
}

type IntMap[C cmp.Ordered] struct {
	size      int
	positions []int
	entries   []Entry[C]
}

// New initializes and returns a new instance of IntMap to handle elements
// ranging from 0 to n-1, where n is the capacity of the map. This constructor
// pre-allocates the necessary memory to accomodate up to n elements, thus
// ensuring that no further memory allocation will be required during the
// instance's lifetime.
func New[C cmp.Ordered](n int) *IntMap[C] {
	return &IntMap[C]{
		size:      0,
		positions: make([]int, n),
		entries:   make([]Entry[C], n+1),
	}
}

// Size returns the number of elements in the IntMap.
func (h *IntMap[C]) Size() int {
	return h.size
}

// Min returns the entry with the smallest cost. The second returned value (ok)
// is a bool that indicates whether a valid entry was found. If the map is
// empty, it returns false, along with a zero value for the entry.
func (h *IntMap[C]) Min() (Entry[C], bool) {
	if h.size == 0 {
		return Entry[C]{}, false
	}
	return h.entries[1], true
}

// Put inserts a new element into the map or updates its cost (and position) if
// it already exists. It returns true if the element was not previously in the
// map; otherwise, it returns false.
func (h *IntMap[C]) Put(elem int, Cost C) bool {
	if pos := h.positions[elem]; pos != 0 { // already in the heap
		h.entries[pos].Cost = Cost
		if p := pos / 2; p >= 1 && h.entries[p].Cost > Cost {
			h.bubbleUp(pos)
		} else {
			h.bubbleDown(pos)
		}
		return false
	}

	h.size++
	h.positions[elem] = h.size
	h.entries[h.size] = Entry[C]{elem, Cost}
	h.bubbleUp(h.size)
	return true
}

// Pop returns and removes the entry with the smallest cost. The second returned
// value (ok) is a bool that indicates whether a valid entry was found. If the
// map is empty, it returns false, along with a zero value for the entry.
func (h *IntMap[C]) Pop() (Entry[C], bool) {
	if h.size == 0 {
		return Entry[C]{}, false
	}
	e := h.entries[1]
	l := h.entries[h.size]
	h.size--
	if h.size > 0 {
		h.entries[1] = l
		h.bubbleDown(1)
	}
	return e, true
}

// Clear removes all the elements contained in the IntMap in O(Size). It is more
// efficient to call Clear than Pop repeatedly.
func (h *IntMap[C]) Clear() {
	for ; h.size > 0; h.size -= 1 {
		h.positions[h.entries[h.size].Elem] = 0
	}
}

func (h *IntMap[C]) String() string {
	bf := strings.Builder{}
	bf.WriteString("IntMap[")
	for i := 1; i <= h.size; i++ {
		bf.WriteString(fmt.Sprintf("%d:%v", h.entries[i].Elem, h.entries[i].Cost))
		if i != h.size {
			bf.WriteByte(' ')
		}
	}
	bf.WriteByte(']')
	return bf.String()
}

func (h *IntMap[C]) bubbleUp(i int) {
	for 1 < i && h.less(i, i/2) {
		h.swap(i, i/2)
		i = i / 2
	}
}

func (h *IntMap[C]) bubbleDown(i int) {
	for i*2 <= h.size {
		n := i * 2 // left
		if n < h.size && h.less(n+1, n) {
			n++ // right child exists and is smaller than the left child
		}
		if h.less(i, n) {
			return
		}
		h.swap(i, n)
		i = n
	}
}

func (h *IntMap[C]) less(i, j int) bool {
	return h.entries[i].Cost < h.entries[j].Cost
}

func (h *IntMap[C]) swap(i, j int) {
	h.entries[i], h.entries[j] = h.entries[j], h.entries[i]
	h.positions[h.entries[i].Elem] = i
	h.positions[h.entries[j].Elem] = j
}
