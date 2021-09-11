package main

import "fmt"

func main() {
	// typed assigment
	var i int
	j := i
	fmt.Println(i, j)

	// untyped assigment with type inference
	a := 43
	b := 1.4
	c := "abc"
	d := 129 + 29i
	fmt.Println(a, b, c, d)
}