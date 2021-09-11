package main

import "fmt"

type Abser interface {
	Abs() float64
}

// MyFloat2 alias for float64
type MyFloat2 float64

// Abs implement interface Abser
func (m MyFloat2) Abs() float64  {
	return float64(m)
}

type Vertex2 struct {
	X, Y float64
}

func (p *Vertex2) Abs() float64  {
	return p.X + p.Y
}

func main() {
	var a Abser
	a = MyFloat2(1.1)
	fmt.Println(a.Abs())

	// error since *Vertex2 implement Abs rather than Vertex2
	// a = Vertex2{1, 3}
	a = &Vertex2{1, 3}
	fmt.Println(a.Abs())
}

