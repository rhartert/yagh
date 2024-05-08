package yagh

import "fmt"

func ExampleNew() {
	m := New[float64](5)

	fmt.Println(m)

	// Add a few elements.
	m.Put(2, 0.1)
	m.Put(1, 0.2)
	m.Put(4, 0.3)
	m.Put(3, 0.4)

	fmt.Println(m)

	// Output:
	// IntMap[]
	// IntMap[2:0.1 1:0.2 4:0.3 3:0.4]
}

func ExampleIntMap_Size() {
	m := New[float64](5)

	// Add and remove some elements.
	m.Put(2, 0.1)
	m.Put(1, 0.2)
	m.Pop()
	m.Put(4, 0.3)
	m.Put(3, 0.4)
	m.Pop()

	fmt.Println(m.Size())

	// Output:
	// 2
}

func ExampleIntMap_Pop() {
	m := New[float64](5)

	// Add a few elements to pop.
	m.Put(2, 0.1)
	m.Put(1, 0.2)
	m.Put(4, 0.3)
	m.Put(3, 0.4)

	for m.Size() > 0 {
		e, _ := m.Pop()
		fmt.Printf("%d:%v\n", e.Elem, e.Cost)
	}

	// Output:
	// 2:0.1
	// 1:0.2
	// 4:0.3
	// 3:0.4
}

func ExampleIntMap_Min() {
	m := New[float64](5)

	// Add a few elements.
	m.Put(2, 0.1)
	m.Put(1, 0.2)
	m.Put(3, 0.3)
	m.Put(4, 0.4)

	e, _ := m.Min()
	fmt.Printf("%d:%v\n", e.Elem, e.Cost)

	// Output:
	// 2:0.1
}

func ExampleIntMap_Put() {
	m := New[float64](5)

	// Add new elements.
	m.Put(2, 0.1)
	m.Put(1, 0.2)
	m.Put(3, 0.3)
	m.Put(4, 0.4)

	// Update the values of existing elements.
	m.Put(1, 0.01)
	m.Put(2, 0.02)

	// Prioritized over elem 1 due to tie-breaking on elem value.
	m.Put(0, 0.01)

	for m.Size() > 0 {
		e, _ := m.Pop()
		fmt.Printf("%d:%v\n", e.Elem, e.Cost)
	}

	// Output:
	// 0:0.01
	// 1:0.01
	// 2:0.02
	// 3:0.3
	// 4:0.4
}

func ExampleIntMap_Clear() {
	m := New[float64](5)

	// Add new elements.
	m.Put(2, 0.1)
	m.Put(1, 0.2)
	m.Put(3, 0.3)
	m.Put(4, 0.4)

	// Remove everything.
	m.Clear()

	fmt.Println(m.Size())

	// Output:
	// 0
}

func ExampleIntMap_GrowBy() {
	m := New[float64](2)

	m.Put(1, 0.1)
	m.GrowBy(3)
	m.Put(2, 0.2)
	m.Put(3, 0.3)

	fmt.Println(m)

	// Output:
	// IntMap[1:0.1 2:0.2 3:0.3]
}

func ExampleIntMap_Capa() {
	m := New[float64](0)
	fmt.Println(m.Capa())

	m.GrowBy(2)
	fmt.Println(m.Capa())

	m.GrowBy(3)
	fmt.Println(m.Capa())

	// Output:
	// 0
	// 2
	// 5
}
