package textures

import (
	"github.com/bit101/blg"
)

type Texture interface {
	Draw(surface *blg.Surface, x, y, w, h float64)
}
