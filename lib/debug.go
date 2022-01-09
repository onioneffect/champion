/*
	debug.go: Print program state information.
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

package champlib

import (
	"fmt"
	"image/color"
	"log"
	"os"
)

var LoggingEnabled = false

func TryLogOutputStr(path string) {
	filePtr, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Printf("ERROR (log file): %s\n", err)
		log.Println("Ignoring log file option...")
	} else {
		log.SetOutput(filePtr)
	}
}

func LogIntarrayInfo(arrptr *[][][3]int32) {
	ChampLog("Array len: ", len(*arrptr))
	ChampLog("Row len: ", len((*arrptr)[0]))
	ChampLog("Cell len: ", len((*arrptr)[0][0]))
}

func LogImgInfo(imginf ImageInfo) {
	ChampLog("Image dimensions: ", imginf.Width, imginf.Height)
	ChampLog("Image format: ", imginf.Format)

	ChampLog("Image bounds: ", (*imginf.Data).Bounds())

	// Makes ColorModel convert an empty color.
	// Returns the corresponding color model.
	// Thanks to https://stackoverflow.com/questions/45226991/
	imgColorModel := imginf.ColorModel.Convert(color.RGBA{})
	msg := fmt.Sprintf("Image color model: %T", imgColorModel)
	ChampLog(msg)

	grayColorModel := color.Gray{}
	firstPix := (*imginf.Decoded)[0][0]
	// Check if image is grayscale
	if imgColorModel == grayColorModel {
		ChampLog("First pixel: ", firstPix[0])
	} else {
		ChampLog("First pixel: ", firstPix)
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

func DebugLineSlice(DebugMe []Line, PrintMe bool) {
	i := 0
	for ; i < len(DebugMe); i++ {
		str, err := DebugMe[i].LineToString()
		if err != nil {
			break
		} else if PrintMe {
			ChampLog(str)
		}
	}

	ChampLog("Looped through ", i, " elements.")
}
