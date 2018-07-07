package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("user id: %d\n", os.Getuid())
	fmt.Printf("group id: %d\n", os.Getgid())
	fmt.Printf("effective user id: %d\n", os.Geteuid())
	fmt.Printf("effective group id: %d\n", os.Getegid())
}
