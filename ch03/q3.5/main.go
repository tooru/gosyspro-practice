package main

import (
	"bytes"
	"io"
	"os"
)

func copyN(w io.Writer, r io.Reader, length int64) (written int64, err error) {
	return io.Copy(w, io.LimitReader(r, length))
}
func main() {
	reader := bytes.NewBufferString("1234567890")
	_, err := copyN(os.Stdout, reader, 5)
	if err != nil {
		panic(err)
	}
}
