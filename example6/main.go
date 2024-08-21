package main

import (
	"errors"
	"fmt"

	"k8s.io/client-go/tools/cache"
)

type Book struct {
	Author string
	Title  string
	Random string
}

func main() {
	fmt.Println("Creating FIFO")
	// Create a store
	keyFunc := func(obj interface{}) (string, error) {
		book, ok := obj.(Book)
		if !ok {
			return "", errors.New("Not a book")
		}
		return book.Author + "/" + book.Title, nil
	}
	myFIFO := cache.NewDeltaFIFOWithOptions(
		cache.DeltaFIFOOptions{
			KeyFunction:           keyFunc,
			KnownObjects:          nil,
			EmitDeltaTypeReplaced: false,
			Transformer:           nil,
		},
	)
	fmt.Println("Adding two books to fifo")

	jane := Book{Author: "Jane", Title: "Persuasion", Random: "A"}
	agatha := Book{Author: "Agatha", Title: "Orient", Random: "A"}

	myFIFO.Add(jane)
	myFIFO.Add(agatha)

	fmt.Println(myFIFO.List())

	jane2 := Book{Author: "Jane", Title: "Persuasion", Random: "B"}
	fmt.Println("performing get on jane")
	myFIFO.Update(jane2)

	fmt.Println("performing list")
	fmt.Println(myFIFO.List())

	fmt.Println("performing get on jane")
	fmt.Println(myFIFO.Get(jane))

	jane3 := Book{Author: "Jane", Title: "Persuasion", Random: "C"}
	fmt.Println("performing get on jane")
	myFIFO.Update(jane3)

	fmt.Println("performing list")
	fmt.Println(myFIFO.List())

	fmt.Println("performing get on jane")
	fmt.Println(myFIFO.Get(jane))

	fmt.Println("performing delete on jane")
	myFIFO.Delete(jane3)

	fmt.Println("performing list")
	fmt.Println(myFIFO.List())

	fmt.Println("performing get on jane")
	fmt.Println(myFIFO.Get(jane))

	myFIFO.Pop(jane2)

}
