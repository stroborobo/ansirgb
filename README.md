ansirgb
=======

A package for converting `color.Color`s to ANSI color codes, 256 colors supported.

Example
-------

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

=> colored output (dark blue):

	 26: 0000 5F5F D7D7
