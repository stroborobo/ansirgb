package ansirgb

import (
	"fmt"
	"image/color"
)

type Color struct {
	color.Color
	Code	int
}

func (c *Color) String() string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%3d: \033[38;5;%dm%04X %04X %04X\033[0m", c.Code, c.Code, r, g, b)
}

var (
	Palette = make(color.Palette, 0, 255)
)

func init() {
	var r uint8 = 0
	var g uint8 = 0
	var b uint8 = 0

	// 8 Bit colors
	var min8 uint8 = 0
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
	var start uint8 = 95
	var step uint8 = 40
	r = 0
	g = 0
	b = 0
	for i := 16; i <= 231; i++ {
		Palette = append(Palette, &Color{&color.RGBA{r, g, b, 255}, i})
		b = next(b, start, step)
		if i > 16 {
			if (i - 16) % 6 == 0 {
				g = next(g, start, step)
			}
			if (i - 16) % 36 == 0 {
				r = next(r, start, step)
			}
		}
	}

	// grayscale
	r = 238
	g = 238
	b = 238
	for i := 232; i <= 255; i++ {
		Palette = append(Palette, &Color{&color.RGBA{r, g, b, 255}, i})
		r -= 10
		g -= 10
		b -= 10
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

func Convert(c color.Color) *Color {
	return Palette.Convert(c).(*Color)
}
