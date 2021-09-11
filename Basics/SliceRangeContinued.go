package main

import "fmt"

func main() {
	s := []int{0,1,4,8,16,32,64,128,256,512}
	for i, _ := range s {
		fmt.Printf("%d\n", i)
	}
	for _, e := range s {
		fmt.Printf("%d\n", e)
	}
	for i := range s {
		fmt.Printf("%d\n", i)
	}
}
