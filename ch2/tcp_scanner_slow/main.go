package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	start := time.Now()

	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or finished
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}

	fmt.Printf("elapsed: %s\n", time.Now().Sub(start))
}
