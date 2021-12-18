// debug.go
package champlib

import (
	"image/color"
	"log"
)

var LoggingEnabled = false

func LogIntarrayInfo(arrptr *[][][3]int32) {
	log.Println("Array len:", len(*arrptr))
	log.Println("Row len:", len((*arrptr)[0]))
	log.Println("Cell len:", len((*arrptr)[0][0]))
}

func LogImgInfo(imginf ImageInfo) {
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

func ChampLog(v ...interface{}) bool {
	if LoggingEnabled {
		log.Print(v...)
		return true
	} else {
		return false
	}
}
