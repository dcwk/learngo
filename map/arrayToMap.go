package main

import "fmt"

func main() {
	currentArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newMap := make(map[int]string)

	for _, item := range currentArray {
		newMap[item] = "test"
	}

	for key, value := range newMap {
		fmt.Println(key, value)
	}
}
