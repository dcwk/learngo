package main

import "fmt"

func main() {
	testMap := make(map[string]string)
	testMap["apple"] = "home"
	testMap["mango"] = "spring"
	testMap["orange"] = "africa"
	var keys [3]string
	var values [3]string
	counter := 0

	for key, value := range testMap {
		keys[counter] = key
		values[counter] = value
		counter++
	}

	fmt.Println(keys)
	fmt.Println(values)
}
