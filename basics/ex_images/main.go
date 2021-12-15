package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (*Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.w, m.h)
}

func (m *Image) At(x, y int) color.Color {
	fx := (float64(x) + .5) / float64(m.w)
	fy := (float64(y) + .5) / float64(m.h)
	r := uint8(math.Round(255 * fx))
	g := uint8(math.Round(255 * fy))
	return color.RGBA{r, g, 255, 255}
}

func main() {
	m := &Image{w: 400, h: 400}
	pic.ShowImage(m)
	
}
