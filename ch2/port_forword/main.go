package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	defer src.Close()

	dst, err := net.Dial("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer src.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal("Unable to bind to port")
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Println("Unable to accept connection")
			continue
		}

		go handle(conn)
	}
}
