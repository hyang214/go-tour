package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64) float64  {
	return fn(3, 4)
}

func main() {
	f := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(f(3, 4))

	fmt.Println(compute(f))
}
