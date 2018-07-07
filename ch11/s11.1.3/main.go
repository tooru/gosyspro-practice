package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("user id: %d\n", os.Getuid())
	fmt.Printf("group id: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("sub group id: %v\n", groups)
}
