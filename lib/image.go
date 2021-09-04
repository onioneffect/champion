// image.go
package imagelib

import (
	"fmt"
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

func ReadImgInfo(img_reader *os.File) ImageInfo {
	var return_info ImageInfo

	config, format, err := image.DecodeConfig(img_reader)
	if err != nil {
		panic(err)
	}

	return_info.Height = config.Height
	return_info.Width = config.Width
	return_info.Format = format
	return_info.ColorModel = config.ColorModel

	_, err = img_reader.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	img_data, _, err := image.Decode(img_reader)
	if err != nil {
		panic(err)
	}
	return_info.Bounds = img_data.Bounds()
	return_info.Data = &img_data

	return return_info
}

func (imginf ImageInfo) PrintImgInfo() {
	fmt.Printf("Image dimensions: %d, %d\n", imginf.Width, imginf.Height)
	fmt.Println("Image format:", imginf.Format)

	fmt.Println("Image bounds:", (*imginf.Data).Bounds())

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
