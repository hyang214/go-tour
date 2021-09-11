package main

import "fmt"

type MPair struct {
	X, Y float64
}

func (m MPair) sth() float64 {
	return m.X * m.Y
}

func (m MPair) other(z float64) float64  {
	return m.sth() + z
}

func main() {
	m := MPair{3, 4}
	fmt.Println(m.sth())
	fmt.Println(m.other(5))
}


