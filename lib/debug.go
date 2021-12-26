/*
	debug.go: Print program state information.
	Copyright (C) 2021  onioneffect

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
)

var LoggingEnabled = false

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
