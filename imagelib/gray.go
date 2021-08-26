// gray.go
package imagelib

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"
)

type ImageInfoGray struct {
	Bounds        image.Rectangle
	Format        string
	ColorModel    color.Model
	Width, Height int
	Data          *image.Image
	Decoded       *[][]int32
}

func ImageArrayGray(im ImageInfo) [][]int32 {
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
