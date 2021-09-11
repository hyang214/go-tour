package main

import "fmt"

func main() {
	// define a pointer refer to an integer value
	var p *int
	// define an integer with value 123
	var i int = 123
	fmt.Println(i)
	// assign the address of variable i to p
	p = &i
	// print the value of p, namely the the address of i
	fmt.Println(p)
	// print the value which p point to, namely the value of i
	fmt.Println(*p)
}
