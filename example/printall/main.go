package main

import (
	"fmt"
	"github.com/stroborobo/ansirgb"
)

func main() {
	for _, c := range ansirgb.Palette {
		fmt.Println(c)
	}
}
