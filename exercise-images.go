package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (image Image) At(x, y int) color.Color {
	return color.RGBA{(uint8(x) + uint8(y)) / 2, (uint8(x) + uint8(y)) / 2, uint8(255), uint8(255)}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
