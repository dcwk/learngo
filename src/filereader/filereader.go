package filereader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Read() {
	file, err := os.Open("C:/Projects/goTest/src/average/data/orders.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
