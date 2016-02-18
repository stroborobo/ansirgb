ansirgb
=======

A package for converting `color.Color`s to ANSI color codes, 256 colors
supported.

Example
-------

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

=> colored output (dark blue)

Installation
------------

	> go get github.com/stroborobo/ansirgb

Have fun :)

Documentation
-------------

Documentation is available at
[GoDoc.org](http://godoc.org/github.com/stroborobo/ansirgb).
