package main

import (
	"fmt"
	"time"
)

const delay = 1 * time.Second

func main() {
	<-time.After(delay)
	fmt.Printf("%v passed.\n", delay)
}
