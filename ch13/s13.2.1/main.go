package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

func main() {
	go sub1(10)
	c := 20
	go func() {
		fmt.Println("shared b capture", c*c)
	}()
	time.Sleep(time.Second)
}
