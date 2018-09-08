package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(path string) *StringFuture {
	promise, future := NewStringFuture()
	go func() {
		content, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
			promise.Close()
		} else {
			future <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource *StringFuture) chan []string {
	promise := make(chan []string)

	go func() {
		var result []string
		for _, line := range strings.Split(futureSource.Get(), "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func countLines(futureSource *StringFuture) chan int {
	promise := make(chan int)
	go func() {
		promise <- len(strings.Split(futureSource.Get(), "\n"))
	}()
	return promise
}

func main() {
	futureSource := readFile("main.go")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
	fmt.Println(<-countLines(futureSource))
}
