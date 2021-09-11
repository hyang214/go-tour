package main

import "fmt"

type M2Pair struct {
	X, Y float64
}

func (m M2Pair) sth() float64 {
	return m.X * m.Y
}

func sth(m M2Pair) float64 {
	return m.X * m.Y
}

func main() {
	m := M2Pair{3, 4}
	fmt.Println(m.sth())
	fmt.Println(sth(m))
}
