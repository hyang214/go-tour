package main

import "fmt"

func main() {
	defer fmt.Println("fist defer")
	fmt.Println("start")
	defer fmt.Println("second defer")
	fmt.Println("end")
	defer fmt.Println("third defer")
}
