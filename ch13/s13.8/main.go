package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var id int64

func generateId(mutex *sync.Mutex) int64 {
	return atomic.AddInt64(&id, 1)
}

func main() {
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
