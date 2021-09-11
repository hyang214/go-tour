package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			// send value x to channel c
			case c <- x:
				// do fibonacci calc
				x, y = y, x+y
			// quit receive a value
			case <-quit:
				fmt.Println("quit")
				return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
