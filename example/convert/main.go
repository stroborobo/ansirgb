package main

import (
	"fmt"
	"image/color"

	"github.com/stroborobo/ansirgb"
)

func main() {
	rgbColor := &color.RGBA{20, 80, 200, 255}
	ansiColor := ansirgb.Convert(rgbColor)
	fmt.Printf("%s%s\n", ansiColor.Fg(), "Foobar")
}
