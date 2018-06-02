package textures

import (
	"github.com/bit101/blg"
	"github.com/bit101/blg/color"
)

type Flat struct {
	color color.Color
}

func NewFlatWithColor(color color.Color) Flat {
	return Flat{color: color}
}

func NewFlatWithRGB(r, g, b float64) Flat {
	return Flat{color: color.RGB(r, g, b)}
}

func (f Flat) Draw(surface *blg.Surface, x, y, w, h float64) {
	surface.Save()
	surface.SetSourceColor(f.color)
	surface.FillRectangle(x, y, w, h)
	surface.Restore()
}
