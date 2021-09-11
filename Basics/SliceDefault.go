package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)

	defaultLow := array[:7]
	fmt.Println(defaultLow)

	defaultHigh := array[4:]
	fmt.Println(defaultHigh)
}
