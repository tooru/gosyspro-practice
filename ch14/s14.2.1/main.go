package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputs := make(chan []byte)

	go func() {
		a, _ := ioutil.ReadFile("a.txt")

		inputs <- a
	}()
	go func() {
		b, _ := ioutil.ReadFile("b.txt")

		inputs <- b
	}()

	fmt.Printf("%v\n", <-inputs)
	fmt.Printf("%v\n", <-inputs)
}
