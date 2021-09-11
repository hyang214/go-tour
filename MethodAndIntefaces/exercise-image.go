package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model  {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 16, 16)
}

func (i Image) At(x, y int) color.Color  {
	r := uint8(x)
	g := uint8(y)
	b := uint8((x + y) % 256)
	a := uint8((x * y) % 256)
	return color.RGBA{r, g, b, a}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
