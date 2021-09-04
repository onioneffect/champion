// line.go
package imagelib

import (
	"fmt"
	"strings"
)

type Line struct {
	// HexColor is just an array with the RGB values.
	HexColor [3]int32

	// Start and end are represented by x and y, so four numbers in total.
	Start, End [2]int32
}

var chars = [16]rune{
	'0', '1', '2', '3',
	'4', '5', '6', '7',
	'8', '9', 'A', 'B',
	'C', 'D', 'E', 'F',
}

func (l Line) RGBToHexNew() string {
	var sb strings.Builder

	for _, val := range l.HexColor {
		bigEnd := val / 16
		sb.WriteRune(chars[bigEnd])

		lilEnd := val % 16
		sb.WriteRune(chars[lilEnd])
	}

	return sb.String()
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
