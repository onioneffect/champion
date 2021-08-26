// main.go
package main

import (
	"fmt"
	"os"

	"github.com/onioneffect/champion/imagelib"
)

func print_test_intarray(arrptr *[][][3]int32) {
	fmt.Println("357:785 (should be 1 1 1) ->", (*arrptr)[785][357])
	fmt.Println("358:785 (should be 76 76 76) ->", (*arrptr)[785][358])
}

func print_intarray_info(arrptr *[][][3]int32) {
	fmt.Println("Array len:", len(*arrptr))
	fmt.Println("Row len:", len((*arrptr)[0]))
	fmt.Println("Cell len:", len((*arrptr)[0][0]))
}

func ImgProcessorRGB(fp *os.File) {
	var currentImg imagelib.ImageInfo = imagelib.Read_img_info(fp)

	// Decodes RGBA into a 3-dimensional array
	// TODO: This probably works with grayscale images,
	// but the resulting array should have dimensions [][][1]!
	var currentDecoded [][][3]int32 = imagelib.ImageArray(currentImg)
	currentImg.Decoded = &currentDecoded

	fmt.Println("Printing image information:")
	currentImg.Print_img_info()

	fmt.Println("\nPrinting array information:")
	print_intarray_info(currentImg.Decoded)

	fmt.Println("\nPrinting test pixels:")
	print_test_intarray(currentImg.Decoded)
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		img_file, err := os.Open(os.Args[i])
		if err != nil {
			panic(err)
		}

		ImgProcessorRGB(img_file)
		img_file.Close()
	}
}
