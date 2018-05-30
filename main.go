package main

import (
	"github.com/bit101/blg"
	"github.com/bit101/gometro/metro"
)

const width = 400.0
const height = 400.0

func main() {
	surface := blg.NewSurface(400, 400)
	surface.ClearRGB(0.9, 0.9, 0.9)
	box := metro.NewBox(100, 200, 100)
	box.Position(200, 300, 0)
	box.Render(surface)

	surface.WriteToPNG("out.png")
}
