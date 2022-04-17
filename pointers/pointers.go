package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(myIntPointer)

	myInt = 18
	fmt.Printf("int value is %d \n", myInt)
	*myIntPointer = 19
	fmt.Printf("int value is %d \n", myInt)
	testPointer(&myInt)
	fmt.Printf("int value is %d \n", myInt)

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
}

func testPointer(myIntPointer *int) {
	*myIntPointer = 20
}
