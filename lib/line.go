// line.go
package imagelib

import "fmt"

type Line struct {
	// Try to store the 24-bit color representation in the lower bits of int32.
	HexColor int32

	// Start and end are represented by x and y, so four numbers in total.
	Start, End [2]int32
}

func ImagePixLoop(im ImageInfo) {
	decodedPtr := im.Decoded
	xLen, yLen := im.Width, im.Height

	fmt.Printf("Looping through pixels: %dx%d\n", xLen, yLen)
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			fmt.Printf("[%d, %d]\t%d\n", x, y, (*decodedPtr)[y][x])
		}
	}
}
