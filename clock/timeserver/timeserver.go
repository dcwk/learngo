package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args[1:]
	listener, err := net.Listen("tcp", "localhost:"+arguments[0])
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
