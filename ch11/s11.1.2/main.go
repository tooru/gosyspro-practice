package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "group id: %d session id: %d\n", syscall.Getpgrp(), sid)
}
