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

const lineFMT string = `{"thickness": 0.1,"color": "%s","points": "%d,%d|%d,%d"}`

func (l Line) LineToString() string {
	var b strings.Builder
	var hColor string = l.RGBToHex()

	fmt.Fprintf(&b, lineFMT, hColor, l.Start[0], l.Start[1], l.End[0], l.End[1])
	fmt.Println(b.String())
	return b.String()
}

func (l Line) Eq(cmp Line) bool {
	return (l.HexColor == cmp.HexColor)
}

func (l Line) RGBToHex() string {
	var sb strings.Builder
	sb.WriteRune('#')

	for _, val := range l.HexColor {
		bigEnd := val / 16
		sb.WriteRune(chars[bigEnd])

		lilEnd := val % 16
		sb.WriteRune(chars[lilEnd])
	}

	return sb.String()
}

func TestPixLoop(im ImageInfo, pixels int) {
	ImagePixLoop(im, pixels%im.Width, pixels/im.Width+1)
}

func ImagePixLoop(im ImageInfo, xLen int, yLen int) {
	decodedPtr := im.Decoded

	fmt.Printf("Looping through pixels: %dx%d\n", xLen, yLen)
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			fmt.Printf("[%d, %d]\t%d\n", x, y, (*decodedPtr)[y][x])
		}
	}
}
