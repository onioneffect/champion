// image_test.go
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
	testLines := [3]champlib.Line{
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
		fp, err := os.Open("tests/" + testFilenames[i])
		if err != nil {
			t.Error(err)
		}
		defer fp.Close()

		testImgInfo := champlib.ReadImgInfo(fp)

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
