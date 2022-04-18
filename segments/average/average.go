package main

import (
	"fmt"
	"log"

	"github.com/dcwk/learngo/filereader"
)

func main() {
	var sum float64
	numbers, err := filereader.ReadFloat("./src/average/data/orders.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range numbers {
		sum += value
	}

	arrayLength := float64(len(numbers))
	average := sum / arrayLength
	fmt.Printf("Average: %.2f\n", average)
}
