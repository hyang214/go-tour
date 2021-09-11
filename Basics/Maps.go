package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var sm1 = make(map[uint8]string)
	sm1[0] = "0"
	sm1[1] = "1"
	fmt.Println(sm1)
}
