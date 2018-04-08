package main

import (
	"fmt"
	"runtime"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	go sub()
	//time.Sleep(2 * time.Second)
	runtime.Goexit()
}
