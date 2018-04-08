package main

import (
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachement; filename=ascii_sample.zip")

	file, err := os.Open("new.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(w, file)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
