// main_test.go
package main

import (
	"testing"

	imagelib "github.com/onioneffect/champion/lib"
)

func TestLineEquality(t *testing.T) {
	white := imagelib.Line{HexColor: [3]int32{255, 255, 255}}
	testLines := [3]imagelib.Line{
		{HexColor: [3]int32{255, 255, 255}},
		{HexColor: [3]int32{128, 128, 128}},
		{HexColor: [3]int32{0, 0, 0}},
	}
	expected := [3]bool{true, false, false}

	for i, j := range testLines {
		if white.Eq(j) != expected[i] {
			t.Errorf("#FFF did not equal %v\n", j)
		}
	}
}

// TODO: Include color compare in this function, maybe
func TestLineFormat(t *testing.T) {
	testLines := [3]imagelib.Line{
		{Start: [2]int32{0, 0}, End: [2]int32{0, 10}},
		{Start: [2]int32{256, 256}, End: [2]int32{128, 1000}},
		{Start: [2]int32{1000, 1000}, End: [2]int32{0, 0}},
	}

	expected := [3]string{
		`{"thickness": 0.1,"color": "#000000","points": "0,0|0,10"}`,
		`{"thickness": 0.1,"color": "#000000","points": "256,256|128,1000"}`,
		`{"thickness": 0.1,"color": "#000000","points": "1000,1000|0,0"}`,
	}

	for i, j := range testLines {
		if actual := j.LineToString(); actual != expected[i] {
			t.Error(actual, expected[i])
		}
	}
}
