package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slices := make([][]uint8, dy, dy)
	for i := range slices {
		slice := make([]uint8, dx, dx)
		slices[i] = slice
		for j := range slice {
			slice[j] = uint8(i) ^ uint8(j)
		}
	}

	return slices
}

func main() {
	pic.Show(Pic)
}
