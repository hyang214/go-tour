package main

import "fmt"

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	var c = 100
	d := 100
	fmt.Println(c)
	fmt.Println(d)
}

func swap(x, y string) (string, string)  {
	return y, x
}
