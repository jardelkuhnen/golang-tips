package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// store values
	m.Store(1, "value number 1")
	m.Store(2, "value number 2")

	// load values
	value, ok := m.Load(1)
	if !ok {
		println("Value not found")
	}
	println(value.(string))

	// load or store values
	value, loaded := m.LoadOrStore(3, "value number 3")
	if !loaded {
		println("Value not found")
	}
	fmt.Println("Loaded:", loaded, "Value:", value) // Output: Loaded: false Value: newValue

	// update value stored
	value, loaded = m.LoadOrStore(1, "new value number 1")
	if !loaded {
		println("Value not found")
	}
	fmt.Println("Loaded:", loaded, "Value:", value)

	// delete value
	value, deleted := m.LoadAndDelete(1)
	if !deleted {
		println("Value not found")
	}
	fmt.Println("Deleted:", deleted, "Value:", value)
}
