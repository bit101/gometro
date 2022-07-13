package textures

import (
	"github.com/bit101/blgg/blcolor"
	"github.com/bit101/blgg/blgg"
)

// Grid is a grid texture.
type Grid struct {
	BackgroundColor blcolor.Color
	ForegroundColor blcolor.Color
	Rows            int
	Cols            int
	LineWidth       float64
}

// NewGrid creates a new Grid texture.
func NewGrid(cols, rows int) Grid {
	return Grid{
		BackgroundColor: blcolor.White(),
		ForegroundColor: blcolor.Black(),
		Rows:            rows,
		Cols:            cols,
		LineWidth:       0.5,
	}
}

// NewGridWithColors creates a new Grid texture with given colors.
func NewGridWithColors(backgroundColor, foregroundColor blcolor.Color, cols, rows int) Grid {
	return Grid{
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		Rows:            rows,
		Cols:            cols,
	}
}

// Draw draws the texture.
func (g Grid) Draw(context *blgg.Context, x, y, w, h float64) {
	context.Push()
	// clip
	context.DrawRectangle(x, y, w, h)
	context.Clip()
	// background
	context.SetColor(g.BackgroundColor)
	context.FillRectangle(x, y, w, h)
	// grid properties
	context.SetColor(g.ForegroundColor)
	context.SetLineWidth(g.LineWidth)
	// horiz lines
	hh := (h - g.LineWidth)
	for i := 0; i <= g.Rows; i++ {
		yy := y + g.LineWidth/2 + float64(i)/float64(g.Rows)*hh
		context.MoveTo(x, yy)
		context.LineTo(x+w, yy)
	}
	// vert lines
	ww := (w - g.LineWidth)
	for i := 0; i <= g.Cols; i++ {
		xx := x + g.LineWidth/2 + float64(i)/float64(g.Cols)*ww
		context.MoveTo(xx, y)
		context.LineTo(xx, y+h)
	}
	context.Stroke()
	context.ResetClip()
	context.Pop()
}
