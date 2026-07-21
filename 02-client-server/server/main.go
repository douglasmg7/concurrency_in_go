package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(c)
	}			
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	num := 0
	for {
		num++
		_, err := io.WriteString(c, fmt.Sprintf("response from server %d\n", num))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
