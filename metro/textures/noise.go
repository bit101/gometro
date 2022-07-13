package textures

import (
	"math"

	"github.com/bit101/blgg/blcolor"
	"github.com/bit101/blgg/blgg"
	"github.com/bit101/blgg/random"
)

// Noise is a noise texture.
type Noise struct {
	backgroundColor blcolor.Color
	foregroundColor blcolor.Color
	density         float64
}

// NewNoise creates a new Noise texture with given density.
func NewNoise(density float64) Noise {
	return Noise{
		backgroundColor: blcolor.White(),
		foregroundColor: blcolor.Black(),
		density:         density,
	}
}

// NewNoiseWithColors creates a new Noise texture with colors and density
func NewNoiseWithColors(backgroundColor, foregroundColor blcolor.Color, density float64) Noise {
	return Noise{
		backgroundColor: backgroundColor,
		foregroundColor: foregroundColor,
		density:         density,
	}
}

// Draw draws the Noise texture to a context.
func (f Noise) Draw(context *blgg.Context, x, y, w, h float64) {
	context.Push()
	context.SetColor(f.backgroundColor)
	context.FillRectangle(x, y, w, h)
	context.SetColor(f.foregroundColor)
	count := math.Abs(w * h * f.density)
	for i := 0.0; i < count; i++ {
		x := random.FloatRange(x, x+w-1)
		y := random.FloatRange(y, y+h-1)
		context.FillRectangle(x, y, 1, 1)
	}

	context.Pop()
}
