package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 2048; i++ {
		wg.Add(1)
		go func(port int) error {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or finished
				return err
			}
			conn.Close()
			fmt.Printf("%d open\n", port)
			return nil
		}(i)
	}
	wg.Wait()

	fmt.Printf("elapsed: %s\n", time.Now().Sub(start))
}
