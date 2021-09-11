package main

import "fmt"

type M4Pair struct {
	X, Y float64
}

func f1(m M4Pair) float64  {
	m.X += 1
	m.Y += 1
	return m.X + m.Y
}

func f2(m *M4Pair) float64  {
	m.X += 1
	m.Y += 1
	return m.X + m.Y
}

func main() {
	m := M4Pair{3, 4}
	fmt.Println(m)
	// does not change the value of m
	fmt.Println(f1(m))
	fmt.Println(m)
	// do change the value of m
	fmt.Println(f2(&m))
	fmt.Println(m)
}
