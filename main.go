package main

import (
	"github.com/bit101/blgo"
	"github.com/bit101/gometro/metro"
	"github.com/bit101/gometro/metro/textures"
)

const width = 400.0
const height = 400.0

func main() {
	surface := blgo.NewSurface(400, 400)
	surface.ClearRGB(0.9, 0.9, 0.9)
	box := metro.NewBox(100, 100, 250)
	box.TopTexture = textures.NewNoise(0.2)
	box.LeftTexture = textures.NewNoise(0.75)
	box.RightTexture = textures.NewNoise(0.4)
	// box.TopTexture = textures.NewGrid(10, 10)
	// box.LeftTexture = textures.NewGrid(10, 20)
	// box.RightTexture = textures.NewGrid(10, 20)
	box.Position(200, 375, 0)
	box.Render(surface)

	surface.WriteToPNG("out.png")
}
