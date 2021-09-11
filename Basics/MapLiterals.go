package main

import "fmt"

type MLPair struct {
	X, Y int64
}

func main() {
	var m = map[string]MLPair {
		"A": MLPair{1,2},
		"B": MLPair{4, 6},
		"C": {7, 8},
	}
	fmt.Println(m)

	var mim = make(map[string]map[int64]MLPair)
	var im = make(map[int64]MLPair)
	im[0] = MLPair{1, 3}
	im[1] = MLPair{5, 3}
	mim["A"] = im
	fmt.Println(mim)
}


