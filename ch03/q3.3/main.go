package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	newFile, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	zipWriter := zip.NewWriter(newFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(writer, strings.NewReader("12345"))
}
