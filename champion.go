// champion.go
package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	imagelib "github.com/onioneffect/champion/lib"
)

func logIntarrayInfo(arrptr *[][][3]int32) {
	log.Println("Array len:", len(*arrptr))
	log.Println("Row len:", len((*arrptr)[0]))
	log.Println("Cell len:", len((*arrptr)[0][0]))
}

func logImgInfo(imginf imagelib.ImageInfo) {
	log.Printf("Image dimensions: %d, %d\n", imginf.Width, imginf.Height)
	log.Println("Image format:", imginf.Format)

	log.Println("Image bounds:", (*imginf.Data).Bounds())

	// Makes ColorModel convert an empty color.
	// Returns the corresponding color model.
	// Thanks to https://stackoverflow.com/questions/45226991/
	imgColorModel := imginf.ColorModel.Convert(color.RGBA{})
	log.Printf("Image color model: %T\n", imgColorModel)

	grayColorModel := color.Gray{}
	firstPix := (*imginf.Decoded)[0][0]
	// Check if image is grayscale
	if imgColorModel == grayColorModel {
		log.Println("First pixel:", firstPix[0])
	} else {
		log.Println("First pixel:", firstPix)
	}
}

func imgProcessor(fp *os.File, debug bool) {
	var currentImg imagelib.ImageInfo = imagelib.ReadImgInfo(fp)
	var currentDecoded [][][3]int32 = imagelib.ImageArray(currentImg)
	currentImg.Decoded = &currentDecoded

	if debug {
		log.Println("We are in debuggign mode!!! :D")

		log.Println("Printing image information:")
		logImgInfo(currentImg)

		log.Println("Printing array information:")
		logIntarrayInfo(currentImg.Decoded)

		log.Println("Running TestPixLoop:")
		imagelib.TestPixLoop(currentImg, 100)
	}
}

func main() {
	var allFiles []string
	var allFilesCtr int = 0
	var useDebugging bool = false
	var logOutputFile string
	var curr string

	for i := 1; i < len(os.Args); i++ {
		curr = os.Args[i]

		if curr == "--debug" {
			useDebugging = true
		} else if curr == "--file" {
			// We increment i so it points to the argument right
			// after "--file", and so the next iteration of the
			// loop doesn't include it in the allFiles list.
			i++
			logOutputFile = os.Args[i]
			fmt.Println(logOutputFile)
		} else {
			allFiles = append(allFiles, curr)
			allFilesCtr++
		}
	}

	if logOutputFile != "" {
		logOutFp, err := os.OpenFile(
			logOutputFile,
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)

		if err != nil {
			log.Printf("ERROR (log file): %s\n", err)
			log.Println("Ignoring log file option...")
		} else {
			log.SetOutput(logOutFp)
		}
	}

	for i := 0; i < allFilesCtr; i++ {
		imgFile, err := os.Open(allFiles[i])
		if err != nil {
			log.Printf("ERROR (image file): %s\n", err)
			continue
		}

		log.Println("Successfully opened file", allFiles[i])
		imgProcessor(imgFile, useDebugging)
		imgFile.Close()
	}
}
