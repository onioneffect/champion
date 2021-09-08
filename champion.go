// champion.go
package main

import (
	"fmt"
	"os"

	imagelib "github.com/onioneffect/champion/lib"
)

func printTestIntArray(arrptr *[][][3]int32) {
	fmt.Println("357:785 (should be 1 1 1) ->", (*arrptr)[785][357])
	fmt.Println("358:785 (should be 76 76 76) ->", (*arrptr)[785][358])
}

func printIntarrayInfo(arrptr *[][][3]int32) {
	fmt.Println("Array len:", len(*arrptr))
	fmt.Println("Row len:", len((*arrptr)[0]))
	fmt.Println("Cell len:", len((*arrptr)[0][0]))
}

func ImgProcessor(fp *os.File) {
	var currentImg imagelib.ImageInfo = imagelib.ReadImgInfo(fp)
	var currentDecoded [][][3]int32 = imagelib.ImageArray(currentImg)
	currentImg.Decoded = &currentDecoded

	fmt.Println("Printing image information:")
	currentImg.PrintImgInfo()

	fmt.Println("\nPrinting array information:")
	printIntarrayInfo(currentImg.Decoded)

	fmt.Println("\nPrinting test pixels:")
	printTestIntArray(currentImg.Decoded)

	//fmt.Println("\nRunning ImagePixLoop:")
	//imagelib.ImagePixLoop(currentImg)

	fmt.Println("\nRunning TestPixLoop:")
	imagelib.TestPixLoop(currentImg, 100)

	fmt.Println("\nTesting color compare:")
	var firstLine, secondLine imagelib.Line
	firstLine.HexColor = [3]int32{255, 255, 255}
	secondLine.HexColor = [3]int32{10, 20, 30}

	firstResult := firstLine.Eq(firstLine)
	secResult := firstLine.Eq(secondLine)
	fmt.Println(firstResult, secResult)
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		imgFile, err := os.Open(os.Args[i])
		if err != nil {
			panic(err)
		}

		ImgProcessor(imgFile)
		imgFile.Close()
	}
}
