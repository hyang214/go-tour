package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	t0 := 0
	t1 := 1
	return func() int {
		if i == 1 {
			i += 1
			return t0
		} else if i == 2 {
			i += 1
			return t1
		} else {
			sum := t0 + t1
			t0 = t1
			t1 = sum
			return t1
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}