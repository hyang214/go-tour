package main

import "fmt"

func count(n int, c chan int)  {
	for i := 0; i < n; i ++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go count(16, c)
	for i := range c {
		fmt.Println(i)
	}
}
