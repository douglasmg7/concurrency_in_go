package main

import (
	"io"
	"log"
	"os"
	"net"
)

func main() {
	// TODO: connect to server on localhost port 8000
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	mustCopy(os.Stdout, c)

}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
