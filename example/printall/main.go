package main

import (
	"fmt"
	"github.com/Knorkebrot/ansirgb"
)

func main() {
	for _, c := range ansirgb.Palette {
		fmt.Println(c)
	}
}
