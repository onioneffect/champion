// line.go
package imagelib

type Line struct {
	// Try to store the 24-bit color representation in the lower bits of int32.
	HexColor int32

	// Start and end are represented by x and y, so four numbers in total.
	Start, End [2]int32
}
