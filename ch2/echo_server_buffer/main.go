package main

import (
	"bufio"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		s, err := rw.ReadString('\n')
		if err != nil {
			log.Println("")
			break
		}
		log.Printf("Read %d bytes: %s\n", len(s), s)

		log.Println("Write data")
		if _, err := rw.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		rw.Flush()
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatal("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")

	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Println("Unable to accept connection")
		}

		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}
