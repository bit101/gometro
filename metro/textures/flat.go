package textures

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/color"
)

// Flat is a flat shaded texture.
type Flat struct {
	color color.Color
}

// NewFlatWithColor creates a new Flat texture with the given color.
func NewFlatWithColor(color color.Color) Flat {
	return Flat{color: color}
}

// NewFlatWithRGB creates a new Flat texture with the given rgb values.
func NewFlatWithRGB(r, g, b float64) Flat {
	return Flat{color: color.RGB(r, g, b)}
}

// Draw draws the texture.
func (f Flat) Draw(surface *blgo.Surface, x, y, w, h float64) {
	surface.Save()
	surface.SetSourceColor(f.color)
	surface.FillRectangle(x, y, w, h)
	surface.Restore()
}
