package main

import "fmt"

func sum(array []int, c chan int) {
	sum := 0
	for _, e := range array {
		sum += e
	}
	c <- sum
}

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := make(chan int)
	go sum(array1, c)
	go sum(array2, c)
	x, y := <- c, <-c
	fmt.Println(x, y)
}
