package main

import (
	"github.com/bit101/blgg/blgg"
	"github.com/bit101/blgg/util"
	"github.com/bit101/gometro/metro"
	"github.com/bit101/gometro/metro/textures"
)

const width = 800.0
const height = 800.0

func main() {
	context := blgg.NewContext(width, height)
	context.ClearRGB(0.9, 0.9, 0.9)
	box := metro.NewBox(100, 100, 1)
	// box.TopTexture = textures.NewNoise(0.2)
	// box.LeftTexture = textures.NewNoise(0.75)
	// box.RightTexture = textures.NewNoise(0.4)
	// box.TopTexture = textures.NewGrid(10, 10)
	// box.LeftTexture = textures.NewGrid(10, 20)
	// box.RightTexture = textures.NewGrid(10, 20)
	box.TopTexture = textures.NewFlatWithRGB(1, 1, 1)
	box.LeftTexture = textures.NewFlatWithRGB(0.5, 0.5, 0.5)
	box.RightTexture = textures.NewFlatWithRGB(0.75, 0.75, 0.75)

	context.Translate(width/2, height/2)
	context.StrokeCircle(0, 0, 20)

	x := 0.0
	y := 0.0
	box.H = 100
	box.Position((x - y), (x+y)/2, 0)
	box.Render(context)

	x = 110.0
	box.H = 50
	box.Position((x - y), (x+y)/2, 0)
	box.Render(context)

	x = 0.0
	y = 110.0
	box.H = 25
	box.Position((x - y), (x+y)/2, 0)
	box.Render(context)

	context.SavePNG("out.png")
	util.ViewImage("out.png")
}
