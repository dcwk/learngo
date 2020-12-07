package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileinfo.Size())
}
