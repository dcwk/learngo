package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal()
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degree Celsius\n", celsius)
}
