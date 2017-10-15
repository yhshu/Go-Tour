package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var ret = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			ret[i][j] = uint8(dx ^ dy + (dx+dy)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
