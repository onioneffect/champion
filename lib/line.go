// line.go
package imagelib

import "fmt"

type Line struct {
	// HexColor is just an array with the RGB values.
	HexColor [3]int32

	// Start and end are represented by x and y, so four numbers in total.
	Start, End [2]int32
}

func (l Line) RGBToHex() (string, error) {
	var fmtString string
	var retString string = "#"

	for _, i := range l.HexColor {
		fmtString = "%X"
		if i < 16 {
			fmtString = "0" + fmtString
		}

		retString += fmt.Sprintf(fmtString, uint8(i))
	}

	return retString, nil
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
