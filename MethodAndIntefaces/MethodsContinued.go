package main

import "fmt"

// MyFloat set an alias for float64
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main() {
	f0 := MyFloat(-100)
	f1 := MyFloat(100)
	fmt.Println(f0.Abs())
	fmt.Println(f1.Abs())
}