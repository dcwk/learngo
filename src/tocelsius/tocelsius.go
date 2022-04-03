package main

import (
	"fmt"
	"log"

	"github.com/dcwk/learngo/src/keyboard"
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
