package main

import "fmt"

func main() {
	a := make([]int, 10)
	printSlice(a)
	b := make([]int, 0, 10)
	printSlice(b)
	c := b[:5]
	printSlice(c)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n",
		len(x), cap(x), x)
}