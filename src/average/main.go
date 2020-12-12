package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64

	for _, value := range numbers {
		sum += value
	}

	arrayLength := float64(len(numbers))
	fmt.Println("Average: ", sum/arrayLength)
}
