package main

import (
	"fmt"
	"net"
	"time"
)

func worker(ports <-chan int) {
	for port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			conn.Close()
			fmt.Println(port, " open")
		} else {
			// fmt.Printf("port: %d, err: %s\n", port, err)
		}
	}
}

func main() {
	start := time.Now()

	ports := make(chan int, 100)
	for i := 0; i < cap(ports); i++ {
		go worker(ports)
	}
	defer close(ports)
	for i := 1; i <= 2048; i++ {
		ports <- i
	}

	fmt.Printf("elapsed: %s\n", time.Now().Sub(start))
}
