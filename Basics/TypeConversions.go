package main

import "fmt"

func main() {
	var i int = 42
	f := float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
}