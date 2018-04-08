package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)

	var buffer bytes.Buffer
	reader2 := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader2, &buffer)
	_, _ = ioutil.ReadAll(teeReader)

	fmt.Printf("%q\n", buffer.String())
}
