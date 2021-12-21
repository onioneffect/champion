// line.go
package champlib

import (
	"errors"
	"fmt"
	"strings"
)

type Line struct {
	// HexColor is just an array with the RGB values.
	HexColor [3]int32

	// Start and end are represented by x and y
	// Two coordinates for each. Four numbers in total.
	Start, End [2]int32

	// StartInit and EndInit reflect if `Start`
	// and `End` have been properly initialized.
	// (We can't just compare each of them to 0,0)
	StartInit, EndInit bool
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

func (l Line) LineToString() (string, error) {
	var b strings.Builder
	var hColor string = l.RGBToHex()

	if !l.StartInit {
		ChampLog("No start coordinates specified!")
		return "", errors.New("cannot stringify without start coordinates")
	} else if !l.EndInit {
		ChampLog("No end in sight!")
		return "", errors.New("cannot stringify without end coordinates")
	}

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

	return b.String(), nil
}

func (l Line) Eq(cmp Line) bool {
	return (l.HexColor == cmp.HexColor)
}

func (lp *Line) SetStart(x, y int32) {
	(*lp).Start = [2]int32{x, y}
	(*lp).StartInit = true
}

func (lp *Line) SetEnd(x, y int32) {
	(*lp).End = [2]int32{x, y}
	(*lp).EndInit = true
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

	var currLine Line
	var lPtr *Line = &currLine

	var currColor, lastColor [3]int32
	var isSame, started bool

	lineSlice := make([]Line, 1024)

	msg := fmt.Sprintf("Looping through pixels: %dx%d\n", xLen, yLen)
	ChampLog(msg)

	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			currColor = (*decodedPtr)[y][x]

			if !started {
				ChampLog("No last color!")

				currLine.HexColor = currColor
				// Debugging purposes
				lPtr.SetStart(int32(x), int32(y))
				lPtr.SetEnd(111, 111)

				lineSlice[0] = currLine
				// /

				started = true
				lastColor = currColor
				continue
			}

			isSame = currColor == lastColor
			if !isSame {
				ChampLog("Color changed!")
			}

			fmt.Println(currColor, x, y)
			lastColor = currColor
		}
	}
}
