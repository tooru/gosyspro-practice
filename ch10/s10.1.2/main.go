package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

type FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		panic(err)
	}
	return &FileLock{fd: fd}
}

func (m *FileLock) Lock() {
	m.l.Lock()
	if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
		panic(err)
	}
}

func (m *FileLock) Unlock() {
	if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
		panic(err)
	}
	m.l.Unlock()
}

func main() {
	N := 2
	lock := NewFileLock("lock")
	ch := make(chan bool)

	for i := 0; i < N; i++ {
		n := i
		go func() {
			fmt.Printf("%d: trylock\n", n)
			lock.Lock()
			fmt.Printf("%d: lock\n", n)
			time.Sleep(500 * time.Millisecond)
			lock.Unlock()
			fmt.Printf("%d: unlock\n", n)
			ch <- true
		}()
	}
	for i := 0; i < N; i++ {
		<-ch
	}
}
