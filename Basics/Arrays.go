package main

import "fmt"

func main() {
	var arrays [2]string
	arrays[0] = "hello"
	arrays[1] = "world"
	fmt.Println(arrays)
	fmt.Println(arrays[0])
	fmt.Println(arrays[1])

	iArrays := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(iArrays)
}
