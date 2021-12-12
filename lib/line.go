// line.go
package champlib

import (
	"fmt"
	"log"
	"strings"
)

type Line struct {
	// HexColor is just an array with the RGB values.
	HexColor [3]int32

	// Start and end are represented by x and y
	// Two coordinates for each. Four numbers in total.
	Start, End [2]int32
}

// Characters used to encode numbers in hexadecimal
var chars = [16]rune{
	'0', '1', '2', '3',
	'4', '5', '6', '7',
	'8', '9', 'A', 'B',
	'C', 'D', 'E', 'F',
}

// JSON line format used by Champ'd Up.

// A `thickness` of 0.1 can closely represent
// a single pixel, but its color may "bleed"
// into the pixels around it, only lighter.

// `color` is encoded as "#000000" to "#FFFFFF".

// `points` are "x,y|x,y" for a straight line,
// described by only two coordinates/points.
const simpleLineFMT string = `{` +
	`"thickness": 0.1,` +
	`"color": "%s",` +
	`"points": "%d,%d|%d,%d"` +
	`}`

// TODO: Using more than two points allows for
// drawing complex paths. Exploiting this fact
// is the most important thing for this project.

func (l Line) LineToString() string {
	var b strings.Builder
	var hColor string = l.RGBToHex()

	fmt.Fprintf(
		// string builder and format string
		&b, simpleLineFMT,

		// Formatting arguments below:
		// RGB color encoded in hexadecimal
		hColor,
		// Starting points
		l.Start[0], l.Start[1],
		// End points
		l.End[0], l.End[1],
	)

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
	// var currLine Line
	var empty [3]int32
	var currColor, lastColor [3]int32
	var isSame bool

	log.Printf("Looping through pixels: %dx%d\n", xLen, yLen)
	for y := 0; y < yLen; y++ {
		// fmt.Println(isSame, y)

		for x := 0; x < xLen; x++ {
			if lastColor == empty {
				log.Print("No last color!")
			}

			currColor = (*decodedPtr)[y][x]
			isSame = currColor == lastColor

			if !isSame {
				log.Print("Color changed!")
			}

			fmt.Println(currColor, x, y)
			lastColor = currColor
		}
	}
}
