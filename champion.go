// champion.go
package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	imagelib "github.com/onioneffect/champion/lib"
)

func printIntarrayInfo(arrptr *[][][3]int32) {
	fmt.Println("Array len:", len(*arrptr))
	fmt.Println("Row len:", len((*arrptr)[0]))
	fmt.Println("Cell len:", len((*arrptr)[0][0]))
}

func PrintImgInfo(imginf imagelib.ImageInfo) {
	log.Printf("Image dimensions: %d, %d\n", imginf.Width, imginf.Height)
	log.Println("Image format:", imginf.Format)

	log.Println("Image bounds:", (*imginf.Data).Bounds())

	// Makes ColorModel convert an empty color.
	// Returns the corresponding color model.
	// Thanks to https://stackoverflow.com/questions/45226991/
	imgColorModel := imginf.ColorModel.Convert(color.RGBA{})
	fmt.Printf("Image color model: %T\n", imgColorModel)

	grayColorModel := color.Gray{}
	firstPix := (*imginf.Decoded)[0][0]
	// Check if image is grayscale
	if imgColorModel == grayColorModel {
		fmt.Println("First pixel:", firstPix[0])
	} else {
		fmt.Println("First pixel:", firstPix)
	}
}

func ImgProcessor(fp *os.File, debug bool) {
	var currentImg imagelib.ImageInfo = imagelib.ReadImgInfo(fp)
	var currentDecoded [][][3]int32 = imagelib.ImageArray(currentImg)
	currentImg.Decoded = &currentDecoded

	if debug {
		log.Println("We are in debuggign mode!!! :D")
	}

	fmt.Println("Printing image information:")
	PrintImgInfo(currentImg)

	fmt.Println("\nPrinting array information:")
	printIntarrayInfo(currentImg.Decoded)

	//fmt.Println("\nRunning ImagePixLoop:")
	//imagelib.ImagePixLoop(currentImg)

	//fmt.Println("\nRunning TestPixLoop:")
	//imagelib.TestPixLoop(currentImg, 100)
}

func main() {
	var allFiles []string
	var allFilesCtr int = 0
	var useDebugging bool = false

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--debug" {
			useDebugging = true
		} else {
			allFiles = append(allFiles, os.Args[i])
			allFilesCtr++
		}
	}

	for i := 0; i < allFilesCtr; i++ {
		imgFile, err := os.Open(allFiles[i])
		if err != nil {
			panic(err)
		}

		ImgProcessor(imgFile, useDebugging)
		imgFile.Close()
	}
}
