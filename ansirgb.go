// Package ansirgb converts color.Color's to ANSI colors
package ansirgb

import (
	"fmt"
	"image/color"
)

// Color represents an ANSI color
type Color struct {
	color.Color
	Code	int
}

// String returns a small description of the color
// containing ANSI code and a hex representation of
// the RGB value, e.g. " 26: 0000 4B4B D2D2"
func (c *Color) String() string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%3d: \033[38;5;%dm%04X %04X %04X\033[0m", c.Code, c.Code, r, g, b)
}

// Fg returns the escape code for foreground color
func (c *Color) Fg() string {
	return fmt.Sprintf("\033[38;5;%dm", c.Code)
}

// Bg returns the escape code for background color
func (c *Color) Bg() string {
	return fmt.Sprintf("\033[48;5;%dm", c.Code)
}

var (
	// Palette contains all 256 ANSI colors
	Palette = make(color.Palette, 0, 255)
)

func init() {
	var r uint8
	var g uint8
	var b uint8

	// 8 Bit colors
	var min8 uint8
	var max8 uint8 = 170
	for i := 0; i <= 15; i++ {
		if i == 8 {
			min8 = 85
			max8 = 255
		}
		if i & 1 == 1 {
			r = max8
		} else {
			r = min8
		}
		if i & 2 == 2 {
			g = max8
		} else {
			g = min8
		}
		if i & 4 == 4 {
			b = max8
		} else {
			b = min8
		}
		Palette = append(Palette, &Color{&color.RGBA{r, g, b, 255}, i})
	}

	// 6 * 6 * 6 = 216 colors
	var start uint8 = 75
	var step uint8 = 45
	r = 0
	g = 0
	b = 0
	for i := 16; i <= 231; i++ {
		Palette = append(Palette, &Color{&color.RGBA{r, g, b, 255}, i})
		b = next(b, start, step)
		if i > 16 {
			if (i - 15) % 6 == 0 {
				g = next(g, start, step)
			}
			if (i - 15) % (6 * 6) == 0 {
				r = next(r, start, step)
			}
		}
	}

	// grayscale
	r = 8
	g = 8
	b = 8
	for i := 232; i <= 255; i++ {
		Palette = append(Palette, &Color{&color.RGBA{r, g, b, 255}, i})
		r += 10
		g += 10
		b += 10
	}
}

func next(n, start, step uint8) uint8 {
	if n == 0 {
		return start
	} else if n == 255 {
		return 0
	}
	return n + step
}

// Convert returns the ANSI color closest to c.
func Convert(c color.Color) *Color {
	return Palette.Convert(c).(*Color)
}
