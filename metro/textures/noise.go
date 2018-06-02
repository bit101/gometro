package textures

import (
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/color"
	"github.com/bit101/blg/random"
)

type Noise struct {
	backgroundColor color.Color
	foregroundColor color.Color
	density         float64
}

func NewNoise(density float64) Noise {
	return Noise{
		backgroundColor: color.White(),
		foregroundColor: color.Black(),
		density:         density,
	}
}

func NewNoiseWithColors(backgroundColor, foregroundColor color.Color, density float64) Noise {
	return Noise{
		backgroundColor: backgroundColor,
		foregroundColor: foregroundColor,
		density:         density,
	}
}

func (f Noise) Draw(surface *blg.Surface, x, y, w, h float64) {
	surface.Save()
	surface.SetSourceColor(f.backgroundColor)
	surface.FillRectangle(x, y, w, h)
	surface.SetSourceColor(f.foregroundColor)
	count := math.Abs(w * h * f.density)
	for i := 0.0; i < count; i++ {
		x := random.FloatRange(x, x+w-1)
		y := random.FloatRange(y, y+h-1)
		surface.FillRectangle(x, y, 1, 1)
	}

	surface.Restore()
}
