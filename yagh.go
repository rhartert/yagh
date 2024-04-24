// Package yagh exposes the IntMap[C] data structure, a priority map that orders
// integers from 0 to N-1 by non-decreasing cost of type `C`. In case of ties,
// the smallest integer wins.
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

func (e Entry[C]) String() string {
	return fmt.Sprintf("%d:%v", e.Elem, e.Cost)
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
	m := &IntMap[C]{
		size:      0,
		positions: make([]int, n),
		entries:   make([]Entry[C], n+1),
	}

	m.entries[0].Elem = -1 // make it explicit that entries[0] does not exist
	for i := range m.positions {
		m.positions[i] = i + 1
		m.entries[i+1].Elem = i
	}

	return m
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
func (h *IntMap[C]) Put(elem int, cost C) bool {
	pos := h.positions[elem]
	h.entries[pos].Cost = cost

	if h.size < pos { // not in the heap
		h.size++
		h.swap(h.size, h.positions[elem])
		h.bubbleUp(h.size)
		return true
	}

	// If the element is already in the heap, change its cost and reposition it
	// in the heap.
	if p := pos / 2; p >= 1 && h.less(pos, p) {
		h.bubbleUp(pos)
	} else {
		h.bubbleDown(pos)
	}
	return false
}

// Pop returns and removes the entry with the smallest cost. The second returned
// value (ok) is a bool that indicates whether a valid entry was found. If the
// map is empty, it returns false, along with a zero value for the entry.
func (h *IntMap[C]) Pop() (Entry[C], bool) {
	if h.size == 0 {
		return Entry[C]{}, false
	}
	e := h.entries[1]
	h.swap(1, h.size)
	h.size--
	if h.size > 1 {
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

// Contains returns true if elem is in the map; it returns false otherwise.
func (h *IntMap[C]) Contains(elem int) bool {
	return h.positions[elem] <= h.size
}

func (h *IntMap[C]) String() string {
	if h.Size() == 0 {
		return "IntMap[]"
	}
	sb := strings.Builder{}
	sb.WriteString("IntMap[")
	sb.WriteString(h.entries[1].String())
	for i := 2; i <= h.size; i++ {
		sb.WriteByte(' ')
		sb.WriteString(h.entries[i].String())
	}
	sb.WriteByte(']')
	return sb.String()
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
	if h.entries[i].Cost > h.entries[j].Cost {
		return false
	}
	return h.entries[i].Cost < h.entries[j].Cost || h.entries[i].Elem < h.entries[j].Elem
}

func (h *IntMap[C]) swap(i, j int) {
	h.entries[i], h.entries[j] = h.entries[j], h.entries[i]
	h.positions[h.entries[i].Elem] = i
	h.positions[h.entries[j].Elem] = j
}
