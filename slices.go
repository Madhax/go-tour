package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		out[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			out[y][x] = uint8(x ^ y)
		}
	}
	return out
}

func main() {
	pic.Show(Pic)
}
