package main

import (
	"ferma/internal/core"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
)

func main() {
	pixelgl.Run(core.Run)
}
