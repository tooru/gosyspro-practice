package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("process ID: %d\n", os.Getpid())
	fmt.Printf("parent process ID: %d\n", os.Getppid())
}
