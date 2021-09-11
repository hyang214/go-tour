package main

import "fmt"

type M5Pair struct {
	X, Y float64
}

func (m *M5Pair) m5f1() float64  {
	m.X += 1
	m.Y += 1
	return m.X + m.Y
}

func m5f2(m *M5Pair) float64  {
	m.X += 1
	m.Y += 1
	return m.X + m.Y
}

func main() {
	m := M5Pair{1, 4}
	fmt.Println(m)
	// the real correct way to call the pointer receiver
	fmt.Println((&m).m5f1())
	// go automatically help us replace m with &m
	fmt.Println(m.m5f1())
	// call method with correct argument
	fmt.Println(m5f2(&m))
	// compiler error since method m5f2 need a pointer, but we give it a value
	// error info: cannot use m (type M5Pair) as type *M5Pair in argument to m5f2
	//fmt.Println(m5f2(m))
}
