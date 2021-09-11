package main

import "fmt"

func main() {
	var s []int
	printSlice1(s)
	for i := 0; i < 100; i ++ {
		s = append(s, i)
		printSlice1(s)
	}
}

func printSlice1(s []int)  {
	fmt.Printf("len: %d cap: %d %v \n", len(s), cap(s), s)
}
