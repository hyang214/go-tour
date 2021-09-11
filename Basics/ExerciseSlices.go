package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dx; i ++ {
		row := make([]uint8, dx)
		for j := 0; j < dy; j ++ {
			row[j] = uint8(i ^ j)
		}
		pic[i] = row
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
