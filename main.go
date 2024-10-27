package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

// The integers are interpreted as bluescale values, where the value 0 means full blue, and the value 255 means full white.

// dx is a length (255), dy is a length (255)
func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)

	// ss - slice slice
	ss := make([][]uint8, dy) // Creates a slice with 255 slices

	for i := 0; i < dy; i++ {
		// s - slice
		s := make([]uint8, dx) // Creates a slice with 255 0's

		for x := 0; x < dx; x++ {
			s[x] = uint8((x + i) / 2)
			//s[x] = uint8(x^i)
			//s[x] = uint8(x*i)
		}

		// Update the slice at index i within the slice ss
		ss[i] = s
	}

	fmt.Println(len(ss))
	return ss
}

func main() {
	pic.Show(Pic)
}
