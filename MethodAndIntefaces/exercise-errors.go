package main

import (
	"fmt"
	"math"
	"strconv"
)

type ErrNegativeSqrt float64

func (f ErrNegativeSqrt) Error() string  {
	return "cannot Sqrt negative number: " + strconv.FormatFloat(float64(f),'g',1,64)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

