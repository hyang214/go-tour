package main

import "fmt"

func main() {
	s := []int{0,1,4,8,16,32,64,128,256,512}
	for i, e := range s {
		fmt.Printf("%d %d\n", i, e)
	}
}
