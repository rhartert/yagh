package yagh

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestIntMap_Min_empty(t *testing.T) {
	m := New[string](5)
	wantEntry := Entry[string]{}
	wantOK := false

	gotEntry, gotOK := m.Min()
	if gotOK != wantOK || gotEntry != wantEntry {
		t.Errorf("Min(): want (%v, %v), got (%v, %v)", wantEntry, wantOK, gotEntry, gotOK)
	}
}

func TestIntMap_Pop_empty(t *testing.T) {
	m := New[string](5)
	wantEntry := Entry[string]{}
	wantOK := false

	gotEntry, gotOK := m.Pop()
	if gotOK != wantOK || gotEntry != wantEntry {
		t.Errorf("Pop(): want (%v, %v), got (%v, %v)", wantEntry, wantOK, gotEntry, gotOK)
	}
}

func TestIntMap_Pop(t *testing.T) {
	n := 100
	rng := rand.New(rand.NewSource(42))
	elems := rng.Perm(n)

	m := New[int](n)
	for _, e := range elems {
		m.Put(e, e)
	}

	for want := 0; want < n; want++ {
		gotEntry, gotOK := m.Pop()
		if !gotOK || gotEntry.Elem != want {
			t.Errorf("Pop(): want (%d, true), got (%d, %v)", want, gotEntry.Elem, gotOK)
		}
	}
}

func TestIntMap_Contains(t *testing.T) {
	nOps := 10_000
	nElems := 100
	rng := rand.New(rand.NewSource(42))

	set := map[int]bool{}
	m := New[int](nElems)

	// Apply a random sequence of put and pop.
	for i := 0; i < nOps; i++ {
		switch rng.Intn(3) { // 33% chance to pop
		case 0: // pop
			e, ok := m.Pop()
			if ok {
				set[e.Elem] = false
			}
		default: // put
			elem := rng.Intn(nElems)
			m.Put(elem, elem)
			set[elem] = true
		}
	}

	// Verify that the map and the set contain the same elements.
	for elem := 0; elem < nElems; elem++ {
		want := set[elem]
		got := m.Contains(elem)
		if want != got {
			t.Errorf("Contains(%d): want %v, got %v", elem, want, got)
		}
	}
}

func TestIntMap_mallocs(t *testing.T) {
	m := New[string](5)

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	m.Put(1, "a")
	m.Min()
	m.Pop()

	runtime.ReadMemStats(&m2)
	if got := m2.Mallocs - m1.Mallocs; got != 0 {
		t.Errorf("Mallocs: want 0, got %d", got)
	}
}
