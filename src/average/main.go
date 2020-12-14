package main

import (
	"filereader"
	"fmt"
	"log"
)

func main() {
	var sum float64
	numbers, err := filereader.ReadFloat("./data/orders.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range numbers {
		sum += value
	}

	arrayLength := float64(len(numbers))
	fmt.Println("Average: ", sum/arrayLength)
}
