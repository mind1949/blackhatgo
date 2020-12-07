package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"time"
)

func worker(ports <-chan int, result chan<- int) {
	for port := range ports {
		address := "scanme.nmap.org:" + strconv.Itoa(port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		result <- port
	}
}

func main() {
	start := time.Now()

	ports := make(chan int, 100)
	result := make(chan int)
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	go func() {
		defer func() {
			close(ports)
			close(result)
		}()
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	var openports []int
	for port := range result {
		openports = append(openports, port)
	}
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

	fmt.Printf("elapsed: %s\n", time.Now().Sub(start))
}
