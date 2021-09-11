package main

import "fmt"

func main() {
	fmt.Println(isOdd(12))
	fmt.Println(isOdd(13))
}

func isOdd(num int) bool  {
	if num % 2 == 0 {
		return false
	} else {
		return true
	}
}