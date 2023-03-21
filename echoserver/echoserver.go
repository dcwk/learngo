package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	arguments := os.Args[1:]
	address := "localhost:" + arguments[0]
	var wg sync.WaitGroup
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server run on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		wg.Add(1)

		go handleConn(conn, wg)
	}

	wg.Wait()
	fmt.Println("No clients")

}

func echo(c net.Conn, shout string, delay time.Duration) {

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn, wg sync.WaitGroup) {
	fmt.Println("New client")
	input := bufio.NewScanner(c)
	for input.Scan() {
		if input.Text() == "Exit" {
			break
		}

		go echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
	fmt.Println("Client disconnected")
	wg.Done()
}
