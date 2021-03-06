package textures

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/color"
)

// Grid is a grid texture.
type Grid struct {
	BackgroundColor color.Color
	ForegroundColor color.Color
	Rows            int
	Cols            int
	LineWidth       float64
}

// NewGrid creates a new Grid texture.
func NewGrid(cols, rows int) Grid {
	return Grid{
		BackgroundColor: color.White(),
		ForegroundColor: color.Black(),
		Rows:            rows,
		Cols:            cols,
		LineWidth:       0.5,
	}
}

// NewGridWithColors creates a new Grid texture with given colors.
func NewGridWithColors(backgroundColor, foregroundColor color.Color, cols, rows int) Grid {
	return Grid{
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		Rows:            rows,
		Cols:            cols,
	}
}

// Draw draws the texture.
func (g Grid) Draw(surface *blgo.Surface, x, y, w, h float64) {
	surface.Save()
	// clip
	surface.Rectangle(x, y, w, h)
	surface.Clip()
	// background
	surface.SetSourceColor(g.BackgroundColor)
	surface.FillRectangle(x, y, w, h)
	// grid properties
	surface.SetSourceColor(g.ForegroundColor)
	surface.SetLineWidth(g.LineWidth)
	// horiz lines
	hh := (h - g.LineWidth)
	for i := 0; i <= g.Rows; i++ {
		yy := y + g.LineWidth/2 + float64(i)/float64(g.Rows)*hh
		surface.MoveTo(x, yy)
		surface.LineTo(x+w, yy)
	}
	// vert lines
	ww := (w - g.LineWidth)
	for i := 0; i <= g.Cols; i++ {
		xx := x + g.LineWidth/2 + float64(i)/float64(g.Cols)*ww
		surface.MoveTo(xx, y)
		surface.LineTo(xx, y+h)
	}
	surface.Stroke()
	surface.Restore()
}
