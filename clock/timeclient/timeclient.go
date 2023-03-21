package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	arguments := os.Args[1:]
	address := "localhost:" + arguments[0]
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Client run on", address)

	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdout)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
