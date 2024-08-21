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
	fmt.Println("Creating store")
	// Create a store
	keyFunc := func(obj interface{}) (string, error) {
		book, ok := obj.(Book)
		if !ok {
			return "", errors.New("Not a book")
		}
		return book.Author + "/" + book.Title, nil
	}
	myStore := cache.NewStore(keyFunc)

	fmt.Println("Adding two books to store")
	jane := Book{Author: "Jane", Title: "Persuasion", Random: "A"}
	agatha := Book{Author: "Agatha", Title: "Orient", Random: "A"}

	myStore.Add(jane)
	myStore.Add(agatha)

	fmt.Println(myStore.List())

	// Get a book
	fmt.Println("Get book again")
	item, exists, err := myStore.Get(jane)
	if !exists || (err != nil) {
		fmt.Println("Hmm...this should exist!")
		panic("Could not get book")
	}
	fmt.Printf("Got book: %s\n", item.(Book))

	// Get book by key
	key, _ := keyFunc(agatha)
	fmt.Printf("Getting book by key, key is %s\n", key)
	item, exists, err = myStore.GetByKey(key)
	if !exists || (err != nil) {
		fmt.Println("Hmm...this should exist!")
		panic("Could not get book")
	}
	fmt.Printf("Got book: %s\n", item.(Book))

	fmt.Println("Update jane")
	jane2 := Book{Author: "Jane", Title: "Persuasion", Random: "B"}
	myStore.Update(jane2)
	fmt.Println(myStore.List())

	fmt.Println(myStore.ListKeys())

	// INDEXED INFORMER

	//keyFunc2 := func(obj interface{}) string {
	//	book, ok := obj.(Book)
	//	if !ok {
	//		return ""
	//	}
	//	return book.Author + "/" + book.Title
	//}
	//
	//authorIndex := func(obj interface{}) ([]string, error) {
	//	book, _ := obj.(Book)
	//	return []string{book.Author}, nil
	//}
	//
	//titleIndex := func(obj interface{}) ([]string, error) {
	//	book, _ := obj.(Book)
	//	return []string{book.Title}, nil
	//}
	//
	//indexedStore := cache.NewThreadSafeStore(
	//	map[string]cache.IndexFunc{
	//		"author": authorIndex,
	//		"title":  titleIndex,
	//	},
	//	map[string]cache.Index{})
	//
	//jane := Book{Author: "Jane", Title: "Persuasion", Random: "A"}
	//agatha := Book{Author: "Agatha", Title: "Orient", Random: "A"}
	//
	//indexedStore.Add(keyFunc2(jane), jane)
	//indexedStore.Add(keyFunc2(agatha), agatha)
	//fmt.Println(indexedStore.ListIndexFuncValues("author"))
}
