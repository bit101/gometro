package textures

import (
	"github.com/bit101/blgo"
)

// Texture is the interface for any drawable texture.
type Texture interface {
	Draw(surface *blgo.Surface, x, y, w, h float64)
}
