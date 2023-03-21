package main

import "fmt"

type Person struct {
	Name       string
	SecondName string
}

func main() {
	var b Person
	var c int = 19
	b = Person{Name: "Bob", SecondName: "White"}
	b.SecondName = "Green"
	fmt.Println(&b)
	fmt.Printf("Pointer from main:%p\n", &b)
	reInit(&c)
	rename(&b)
	fmt.Println(b.Name)
}

func reInit(c *int) {
	fmt.Println("Value from reInit", c)
}

func rename(a *Person) {
	fmt.Println("Name from rename:", *a)
	fmt.Printf("Pointer from rename:%p\n", a)
	*a = Person{Name: "Alex", SecondName: "Black"}
	fmt.Printf("Pointer from rename:%p\n", a)
}
