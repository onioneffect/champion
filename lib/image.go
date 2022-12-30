/*
	image.go: Read image metadata.
	Copyright (C) 2021-2023  onioneffect

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
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

type ImageInfo struct {
	Bounds        image.Rectangle
	Format        string
	ColorModel    color.Model
	Width, Height int
	Data          *image.Image
	Decoded       *[][][3]int32
}

// Thanks to https://stackoverflow.com/questions/33186783/
func ImageArray(im ImageInfo) [][][3]int32 {
	width, height := im.Width, im.Height
	iaa := make([][][3]int32, height)
	src_rgba := image.NewRGBA(im.Bounds)
	draw.Copy(src_rgba, image.Point{}, (*im.Data), im.Bounds, draw.Src, nil)

	for y := 0; y < height; y++ {
		row := make([][3]int32, width)
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			row[x] = [3]int32{int32(pix[0]), int32(pix[1]), int32(pix[2])}
		}

		iaa[y] = row
	}

	return iaa
}

func ReadImgInfo(imgReader *os.File) (ImageInfo, error) {
	var returnInfo ImageInfo

	config, format, err := image.DecodeConfig(imgReader)
	if err != nil {
		return ImageInfo{}, err
	}

	returnInfo.Height = config.Height
	returnInfo.Width = config.Width
	returnInfo.Format = format
	returnInfo.ColorModel = config.ColorModel

	_, err = imgReader.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	imgData, _, err := image.Decode(imgReader)
	if err != nil {
		panic(err)
	}
	returnInfo.Bounds = imgData.Bounds()
	returnInfo.Data = &imgData

	return returnInfo, nil
}
