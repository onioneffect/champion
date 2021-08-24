// gray.go
package imagelib

import (
	"image"

	"golang.org/x/image/draw"
)

func image_array_gray(im ImageInfo) [][]int32 {
	width, height := im.Width, im.Height
	iaa := make([][]int32, height)
	src_rgba := image.NewRGBA(im.Bounds)
	draw.Copy(src_rgba, image.Point{}, (*im.Data), im.Bounds, draw.Src, nil)

	for y := 0; y < height; y++ {
		row := make([]int32, width)
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			row[x] = int32(pix[0])
		}

		iaa[y] = row
	}

	return iaa
}
