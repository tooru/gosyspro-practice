package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	randReader := rand.Reader
	newFile, err := os.Create("rand.dat")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, io.LimitReader(randReader, 1024))
}
