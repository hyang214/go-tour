package main

import "fmt"

func main()  {
	fmt.Println("add function: ", add(13, 30))
}

func add(x int, y int) int {
	return x + y
}