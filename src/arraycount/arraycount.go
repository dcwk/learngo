package main

import (
	"fmt"
)

func main() {
	var testArray [8]int
	counter := 0

	for i := 0; i < len(testArray); i++ {
		if i%2 == 0 {
			continue
		}

		testArray[i] = i * 5
		counter++
	}

	fmt.Printf("%#v\n", testArray)
	fmt.Printf("Array elements count: %d\n", counter)

	fmt.Println("For ... range cycle test")

	for index, value := range testArray {
		fmt.Printf("Index %d and value %d\n", index, value)
	}
}
