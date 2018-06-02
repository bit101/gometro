package metro

import (
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/gometro/metro/textures"
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

// Render renders the box to a surface
func (box *Box) Render(surface *blg.Surface) {
	surface.Save()
	surface.Translate(box.X, box.Y-box.Z)

	box.drawBack(surface)
	box.drawLeftWall(surface)
	box.drawRightWall(surface)
	box.drawTop(surface)
	surface.Restore()
}

func (box *Box) drawBack(surface *blg.Surface) {
	// draw a triangle across all face seams to prevent background color leaking through
	surface.Save()
	surface.SetSourceRGB(0.5, 0.5, 0.5)
	surface.MoveTo(0, 0)
	surface.LineTo(-box.D, -box.H-box.D/2)
	surface.LineTo(box.W, -box.H-box.W/2)
	surface.Fill()
	surface.Restore()
}

func (box *Box) drawLeftWall(surface *blg.Surface) {
	surface.Save()
	surface.Transform(leftMatrix)
	box.LeftTexture.Draw(surface, -box.D, -box.H, box.D, box.H)
	surface.Restore()
}

func (box *Box) drawRightWall(surface *blg.Surface) {
	surface.Save()
	surface.Transform(rightMatrix)
	box.RightTexture.Draw(surface, 0, -box.H, box.W, box.H)
	surface.Restore()
}

func (box *Box) drawTop(surface *blg.Surface) {
	surface.Save()
	surface.Translate(0, -box.H)
	surface.Scale(1, 0.5)
	surface.Scale(math.Sqrt2, math.Sqrt2)
	surface.Rotate(math.Pi / 4)
	box.TopTexture.Draw(surface, -box.D, -box.W, box.D, box.W)
	surface.Restore()
}
