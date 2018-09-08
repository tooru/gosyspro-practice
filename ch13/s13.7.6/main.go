package main

import (
	"fmt"
	"sync"
)

func showSmap(smap *sync.Map) {
	fmt.Printf("show sync.Map\n")

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}

func main() {
	smap := &sync.Map{}

	smap.Store("hello", "world")
	smap.Store(1, 2)

	smap.Delete("test")

	value, ok := smap.Load("hello")
	fmt.Printf("key=%v value=%v exists?=%v\n", "hello", value, ok)

	showSmap(smap)
	smap.LoadOrStore(1, 3)
	showSmap(smap)
	smap.LoadOrStore(2, 4)
	showSmap(smap)
}
