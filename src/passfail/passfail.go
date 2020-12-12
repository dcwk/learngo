//passfail комментарий как во многих других языках
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	status := "Passing!"

	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	if grade <= 60 {
		status = "Failing!"
	}

	fmt.Println(status)
}
