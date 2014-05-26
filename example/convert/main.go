package main

import (
	"fmt"
	"image/color"
	"github.com/Knorkebrot/ansirgb"
)

func main() {
	rgbColor := &color.RGBA{20, 80, 200, 255}
	ansiColor := ansirgb.Convert(rgbColor)
	fmt.Println(ansiColor)
}
