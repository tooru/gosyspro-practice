package main

import (
	"fmt"
	"net"
	"time"
)

const interval = 10 * time.Second

func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	start := time.Now()
	wait := start.Round(interval).Add(interval).Sub(start)
	time.Sleep(wait)
	ticker := time.Tick(interval)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
	fmt.Println("Receive from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Receive: %s\n", string(buffer[:length]))
}
