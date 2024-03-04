package yagh

import (
	"math/rand"
	"testing"
)

func TestIntMap_Min_empty(t *testing.T) {
	m := New[string](5)
	if got := m.Min(); got != nil {
		t.Errorf("Min(): want nil, got %v", got)
	}
}

func TestIntMap_Pop_empty(t *testing.T) {
	m := New[string](5)
	if got := m.Pop(); got != nil {
		t.Errorf("Pop(): want nil, got %v", got)
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
		if got := m.Pop(); got.Elem != want {
			t.Errorf("Pop(): want %d, got %d", want, got.Elem)
		}
	}
}
