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
	n := 100
	rng := rand.New(rand.NewSource(42))

	// Verify that elements are not in the heap before being added.
	m := New[float64](n)
	for elem := 0; elem < 100; elem++ {
		if m.Contains(elem) {
			t.Errorf("Contains(%d): want false, got true", elem)
		}
		m.Put(elem, rng.Float64())
		if !m.Contains(elem) {
			t.Errorf("Contains(%d): want true, got false", elem)
		}
	}

	// Verify that elements are not in the heap after being removed.
	for i := 0; i < 100; i++ {
		next, _ := m.Min()
		elem := next.Elem

		if !m.Contains(elem) {
			t.Errorf("Contains(%d): want true, got false", elem)
		}
		m.Pop()
		if m.Contains(elem) {
			t.Errorf("Contains(%d): want false, got true", elem)
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
