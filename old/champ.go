package main

import (
	"fmt"
	"os"
	"image"
	"image/png"
	"errors"
	"image/color"
)

var lineFMT string = `{"thickness": 0.1,"color": "%s","points": "%s"},`
var coordFMT string = "%d,%d|"

func color_compare(first, second color.Color) bool {
	r, g, b, _ := first.RGBA()
	t, h, n, _ := second.RGBA()

	return (uint8(r) == uint8(t) && uint8(g) == uint8(h) && uint8(b) == uint8(n))
}

func hexRGB(args ...uint32) (string, error) {
	if len(args) != 3 {
		return "", errors.New("hexRGB was called with an incorrect number of arguments")
	}

	var coolFormat string
	var coolReturn string = "#"

	for _, i := range args {
		coolFormat = "%X"
		if i < 16 {
			coolFormat = "0" + coolFormat
		}

		coolReturn += fmt.Sprintf(coolFormat, uint8(i))
	}

	return coolReturn, nil
}

func image_loop(ptr *image.Image) (finalLines string) {
	var x int = (*ptr).Bounds().Max.X
	var y int = (*ptr).Bounds().Max.Y
	var current_color, last_color color.Color
	var coords string
	var same bool
	var r, g, b uint32

	for i := 0; i < x; i++ {
		fmt.Println(same, i)

		for j := 0; j < y; j++ {
			current_color = (*ptr).At(i, j)

			if last_color != nil {
				same = color_compare(current_color, last_color)
				r, g, b, _ = last_color.RGBA()
			} else {
				r, g, b, _ = current_color.RGBA()
			}

			if same && j != 0 {
				continue
			} else if coords != "" {
				h, _ := hexRGB(r, g, b)

				coords = coords[:len(coords) - 1] // Remove last char

				finalLines += fmt.Sprintf(lineFMT, h, coords)

				coords = ""
				coords += fmt.Sprintf(coordFMT, i, j)
			}

			fmt.Println(finalLines)
			last_color = current_color
		}
	}

	return
}

func main() {
	//imgfile, err := os.Open(path + "\\320.png")
	imgfile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer imgfile.Close()

	img, err := png.Decode(imgfile)
	if err != nil {
		panic(err.Error())
	}

	var bingus string = image_loop(&img)
	fmt.Println(len(bingus))

	payload, err := os.Create(os.Args[2])
	if err != nil {
		panic(err.Error)
	}
	defer payload.Close()

	payload.WriteString(bingus)
	return
}

