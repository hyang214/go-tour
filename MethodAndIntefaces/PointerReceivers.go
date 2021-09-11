package main

import "fmt"

type M3Pair struct {
	X, Y float64
}

func (m M3Pair) f1() float64  {
	m.X += 1
	m.Y += 1
	return m.X + m.Y
}

func (m *M3Pair) f2() float64  {
	m.X += 1
	m.Y += 1
	return m.X + m.Y
}

func main() {
	m := M3Pair{3, 4}
	fmt.Println(m)
	// does not change the value of m
	fmt.Println(m.f1())
	fmt.Println(m)
	// do change the value of m
	fmt.Println(m.f2())
	fmt.Println(m)
}