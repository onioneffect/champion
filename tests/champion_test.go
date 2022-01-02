/*
	champion_test.go: Test the functionalities of the program.
	Copyright (C) 2021-2022  onioneffect

	This file is part of Champion.

	Champion is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Champion is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with Champion.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"os"
	"testing"

	"image"

	champlib "github.com/onioneffect/champion/lib"
)

func TestLineEquality(t *testing.T) {
	white := champlib.Line{HexColor: [3]int32{255, 255, 255}}
	testLines := [3]champlib.Line{
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
	testLines := [3]champlib.Line{}

	testLines[0].SetStart(0, 0)
	testLines[0].SetEnd(0, 10)

	testLines[1].SetStart(256, 256)
	testLines[1].SetEnd(128, 1000)

	testLines[2].SetStart(1000, 1000)
	testLines[2].SetEnd(0, 0)

	expected := [3]string{
		`{"thickness":0.1,"color":"#000000","points":"0,0|0,10"}`,
		`{"thickness":0.1,"color":"#000000","points":"256,256|128,1000"}`,
		`{"thickness":0.1,"color":"#000000","points":"1000,1000|0,0"}`,
	}

	for i, j := range testLines {
		actual, err := j.LineToString()

		if actual != expected[i] {
			t.Error(actual, expected[i])
		} else if err != nil {
			t.Error(err)
		}
	}
}

func boundGenerator(twoDimensions [3][2]int) (ret []image.Rectangle) {
	for i := 0; i < 3; i++ {
		currObj := image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: twoDimensions[i][0],
				Y: twoDimensions[i][1],
			},
		}

		ret = append(ret, currObj)
	}

	return
}

func TestReadImgInfo(t *testing.T) {
	testFilenames := [3]string{
		"bird.png",
		"burger.png",
		"donda.jpg",
	}

	testDimensions := [3][2]int{
		{450, 450},
		{5000, 5000},
		{512, 512},
	}
	expectedBounds := boundGenerator(testDimensions)

	expected := [3]champlib.ImageInfo{
		{
			Bounds: expectedBounds[0],
			Format: "png",
			Width:  450,
			Height: 450,
		},

		{
			Bounds: expectedBounds[1],
			Format: "png",
			Width:  5000,
			Height: 5000,
		},

		{
			Bounds: expectedBounds[2],
			Format: "jpeg",
			Width:  512,
			Height: 512,
		},
	}

	for i := 0; i < 3; i++ {
		fp, err := os.Open("test_img/" + testFilenames[i])
		if err != nil {
			t.Error(err)
		}
		defer fp.Close()

		testImgInfo, err := champlib.ReadImgInfo(fp)

		if err != nil {
			t.Error(err)
		}

		if testImgInfo.Bounds != expected[i].Bounds {
			t.Errorf("Mismatched bounds. Index %d", i)
		}

		if testImgInfo.Width != expected[i].Width {
			t.Errorf("Mismatched dimensions. Index %d", i)
		}

		if testImgInfo.Format != expected[i].Format {
			t.Errorf("Mismatched format. Index %d", i)
		}
	}
}
