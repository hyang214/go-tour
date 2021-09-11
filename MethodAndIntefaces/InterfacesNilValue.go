package main

import "fmt"

type INVF interface {
	M()
}

type INV struct {
	S string
}

func (e *INV) M()  {
	if e == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(e.S)
	}
}

func main() {
	var invf INVF
	// will throw error since invf not assign value
	// invf.M()

	// assign a nil pointer
	var inv *INV
	invf = inv
	invf.M()

	// assign a real pointer
	invf = &INV{"abc"}
	invf.M()
}
