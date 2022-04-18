package main

import (
	"fmt"
	"log"

	"github.com/dcwk/learngo/incapsulation/calendar"
)

func main() {
	date := calendar.DateTime{}
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

	event := calendar.Event{}
	event.SetTitle("Test event")
	event.SetDay(1)
	event.SetMonth(4)
	event.SetYear(2018)

	fmt.Println("Event date is", event.Day())
}
