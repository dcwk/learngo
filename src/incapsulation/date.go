package main

import (
	"fmt"
	"log"

	"github.com/dcwk/learngo/src/incapsulation/date"
)

func main() {
	date := date.DateTime{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(14)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(date.Day())
	fmt.Println(date.Month())
	fmt.Println(date.Year())
}
