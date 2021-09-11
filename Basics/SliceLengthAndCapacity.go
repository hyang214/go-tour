package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)

	slice := array[3:8]
	fmt.Println(slice)
	// from 3 to 7
	fmt.Println(len(slice))
	// from 3 to 9
	fmt.Println(cap(slice))
}
