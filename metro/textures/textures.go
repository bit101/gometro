package textures

import "github.com/bit101/blgg/blgg"

// Texture is the interface for any drawable texture.
type Texture interface {
	Draw(context *blgg.Context, x, y, w, h float64)
}
