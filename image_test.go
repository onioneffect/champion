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
