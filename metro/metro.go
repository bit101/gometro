package metro

import (
	"math"

	"github.com/bit101/blg"
	cairo "github.com/ungerik/go-cairo"
)

var leftMatrix = cairo.Matrix{
	Xx: 1,
	Yx: 0.5,
	Xy: 0,
	Yy: 1,
	X0: 0,
	Y0: 0,
}

var rightMatrix = cairo.Matrix{
	Xx: 1,
	Yx: -0.5,
	Xy: 0,
	Yy: 1,
	X0: 0,
	Y0: 0,
}

// Box is a box
type Box struct {
	X float64
	Y float64
	Z float64
	W float64
	D float64
	H float64
}

// NewBox creates a new box of the given size
func NewBox(w, d, h float64) *Box {
	return &Box{0, 0, 0, w, d, h}
}

// Position positions the box
func (box *Box) Position(x, y, z float64) {
	box.X = x
	box.Y = y
	box.Z = z
}

// Size sizes the box
func (box *Box) Size(w, d, h float64) {
	box.W = w
	box.D = d
	box.H = h
}

// Render renders the box to a surface
func (box *Box) Render(surface *blg.Surface) {
	surface.Save()
	surface.Translate(box.X, box.Y-box.Z)

	box.drawLeftWall(surface)
	box.drawRightWall(surface)
	box.drawTop(surface)
	surface.Restore()
}

func (box *Box) drawLeftWall(surface *blg.Surface) {
	surface.Save()
	surface.SetSourceRGB(0.5, 0.5, 0.5)
	surface.Transform(leftMatrix)
	surface.FillRectangle(0, 0, -box.D, -box.H)
	surface.Restore()
}

func (box *Box) drawRightWall(surface *blg.Surface) {
	surface.Save()
	surface.SetSourceRGB(0.75, 0.75, 0.75)
	surface.Transform(rightMatrix)
	surface.FillRectangle(0, 0, box.W, -box.H)
	surface.Restore()
}

func (box *Box) drawTop(surface *blg.Surface) {
	surface.Save()
	surface.SetSourceRGB(1, 1, 1)
	surface.Translate(0, -box.H)
	surface.Scale(1, 0.5)
	surface.Scale(math.Sqrt2, math.Sqrt2)
	surface.Rotate(math.Pi / 4)
	surface.FillRectangle(0, 0, -box.D, -box.W)
	surface.Restore()
}
