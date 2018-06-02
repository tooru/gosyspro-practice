package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("  name: %v\n", info.Name())
	fmt.Printf("  size: %v\n", info.Size())
	fmt.Printf("  modTime: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  directory?: %v\n", info.Mode().IsDir())
	fmt.Printf("  regular?: %v\n", info.Mode().IsRegular())
	fmt.Printf("  permission: %v\n", info.Mode().Perm())
	fmt.Printf("  mode text: %v\n", info.Mode().String())
}
