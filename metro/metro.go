package metro

import (
	"math"

	"github.com/bit101/blgg/blgg"
	"github.com/bit101/gometro/metro/textures"
	"github.com/fogleman/gg"
)

var leftMatrix = gg.Matrix{
	XX: 1,
	YX: 0.5,
	XY: 0,
	YY: 1,
	X0: 0,
	Y0: 0,
}

var rightMatrix = gg.Matrix{
	XX: 1,
	YX: -0.5,
	XY: 0,
	YY: 1,
	X0: 0,
	Y0: 0,
}

// Box is a box
type Box struct {
	X            float64
	Y            float64
	Z            float64
	W            float64
	D            float64
	H            float64
	TopTexture   textures.Texture
	LeftTexture  textures.Texture
	RightTexture textures.Texture
}

// NewBox creates a new box of the given size
func NewBox(w, d, h float64) *Box {
	return &Box{
		X:            0,
		Y:            0,
		Z:            0,
		W:            w,
		D:            d,
		H:            h,
		TopTexture:   textures.NewFlatWithRGB(1, 1, 1),
		LeftTexture:  textures.NewFlatWithRGB(0.5, 0.5, 0.5),
		RightTexture: textures.NewFlatWithRGB(0.75, 0.75, 0.75),
	}
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

// Render renders the box to a context
func (box *Box) Render(context *blgg.Context) {
	context.Push()
	context.Translate(box.X, box.Y-box.Z)

	box.drawBack(context)
	box.drawLeftWall(context)
	box.drawRightWall(context)
	box.drawTop(context)
	context.Pop()
}

func (box *Box) drawBack(context *blgg.Context) {
	// draw a triangle across all face seams to prevent background color leaking through
	context.Push()
	context.SetRGB(0.5, 0.5, 0.5)
	context.MoveTo(0, 0)
	context.LineTo(-box.D, -box.H-box.D/2)
	context.LineTo(box.W, -box.H-box.W/2)
	context.Fill()
	context.Pop()
}

func (box *Box) drawLeftWall(context *blgg.Context) {
	context.Push()
	context.Shear(leftMatrix.XY, leftMatrix.YX)
	box.LeftTexture.Draw(context, -box.D, -box.H, box.D, box.H)
	context.Pop()
}

func (box *Box) drawRightWall(context *blgg.Context) {
	context.Push()
	context.Shear(rightMatrix.XY, rightMatrix.YX)
	box.RightTexture.Draw(context, 0, -box.H, box.W, box.H)
	context.Pop()
}

func (box *Box) drawTop(context *blgg.Context) {
	context.Push()
	context.Translate(0, -box.H)
	context.Scale(1, 0.5)
	context.Scale(math.Sqrt2, math.Sqrt2)
	context.Rotate(math.Pi / 4)
	box.TopTexture.Draw(context, -box.D, -box.W, box.D, box.W)
	context.Pop()
}
