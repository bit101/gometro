package textures

import (
	"github.com/bit101/blgg/blcolor"
	"github.com/bit101/blgg/blgg"
)

// Flat is a flat shaded texture.
type Flat struct {
	color blcolor.Color
}

// NewFlatWithColor creates a new Flat texture with the given color.
func NewFlatWithColor(color blcolor.Color) Flat {
	return Flat{color: color}
}

// NewFlatWithRGB creates a new Flat texture with the given rgb values.
func NewFlatWithRGB(r, g, b float64) Flat {
	return Flat{color: blcolor.RGB(r, g, b)}
}

// Draw draws the texture.
func (f Flat) Draw(context *blgg.Context, x, y, w, h float64) {
	context.Push()
	context.SetColor(f.color)
	context.FillRectangle(x, y, w, h)
	context.Pop()
}
