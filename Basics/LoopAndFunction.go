package main

import (
	"fmt"
	"math"
)

const t = 0.000000000001

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(x - z * z) > t {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	for i := 0.0; i < 10; i ++ {
		fmt.Println(math.Sqrt(i), Sqrt(i))
	}
}

