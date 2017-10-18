package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (m *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (m *Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := &Image{}
	pic.ShowImage(m)
}
