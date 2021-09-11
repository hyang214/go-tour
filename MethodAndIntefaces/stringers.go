package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

func (p Person) String() string  {
	return p.firstName + "-" + p.lastName
}

func main() {
	a := Person{"Bob", "Alex"}
	fmt.Println(a)
}
