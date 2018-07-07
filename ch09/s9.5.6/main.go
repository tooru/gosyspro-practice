package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Match("image-*.png", "image-100.png"))

	files, err := filepath.Glob("./*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
