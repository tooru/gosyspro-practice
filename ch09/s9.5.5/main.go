package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func Clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	fmt.Println(Clean2("./path/filepath/../path.go"))
}
