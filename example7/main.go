package main

import (
	"fmt"

	"k8s.io/utils/buffer"
)

func main() {
	// Create a store

	fmt.Println("------------TEST 1 ---------------")
	ring := *buffer.NewRingGrowing(3)
	val, ok := ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)

	ring.WriteOne(1)
	val, ok = ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)

	val, ok = ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)

	fmt.Println("------------TEST 2 ---------------")

	ring.WriteOne(1)
	ring.WriteOne(2)
	ring.WriteOne(3)
	ring.WriteOne(4)

	val, ok = ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)

	val, ok = ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)

	val, ok = ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)

	val, ok = ring.ReadOne()
	fmt.Println(val)
	fmt.Println(ok)
}
