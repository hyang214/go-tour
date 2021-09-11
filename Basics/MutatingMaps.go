package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["A"] = 1
	fmt.Println(m)
	m["A"] = 2
	fmt.Println(m)
	e, ok := m["A"]
	fmt.Println("e: ", e, " ok?", ok)

	delete(m, "A")
	fmt.Println(m)
	e, ok = m["A"]
	fmt.Println("e: ", e, " ok?", ok)
}
