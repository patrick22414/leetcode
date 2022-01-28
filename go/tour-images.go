package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w     int
	h     int
	bytes [][]byte
}

func NewImage() Image {
	w, h := 400, 200
	bytes := make([][]byte, w)
	for row := range bytes {
		bytes[row] = make([]byte, h)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			bytes[i][j] = byte(i * j)
		}
	}

	return Image{w, h, bytes}
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{
		R: img.bytes[x][y],
		G: img.bytes[x][y],
		B: 255,
		A: 255,
	}
}

func main() {
	m := NewImage()
	pic.ShowImage(m)
}
