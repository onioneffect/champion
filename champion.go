package main

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
}

// Thanks to https://stackoverflow.com/questions/33186783/
func image_array(im ImageInfo) [][][3]float32 {
	width, height := im.Width, im.Height
	iaa := make([][][3]float32, height)
	src_rgba := image.NewRGBA(im.Bounds)
	draw.Copy(src_rgba, image.Point{}, (*im.Data), im.Bounds, draw.Src, nil)

	for y := 0; y < height; y++ {
		row := make([][3]float32, width)
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			row[x] = [3]float32{float32(pix[0]), float32(pix[1]), float32(pix[2])}
		}
		iaa[y] = row
	}

	return iaa
}

func read_img_info(img_reader *os.File) ImageInfo {
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

func (imginf ImageInfo) print_img_info() {
	fmt.Printf("Image dimensions: %d, %d\n", imginf.Width, imginf.Height)
	fmt.Println("Image format:", imginf.Format)

	// Makes ColorModel convert an empty color.
	// Returns the corresponding color model.
	// Thanks to https://stackoverflow.com/questions/45226991/
	fmt.Printf("Image color model: %T\n", imginf.ColorModel.Convert(color.RGBA{}))

	fmt.Println("Image bounds:", (*imginf.Data).Bounds())
	fmt.Println("First pixel:", (*imginf.Data).At(0, 0))
}

func img_processor(fp *os.File) {
	var currentImg ImageInfo = read_img_info(fp)
	var asdasd [][][3]float32 = image_array(currentImg)
	currentImg.print_img_info()
	fmt.Println(">>>", asdasd[0][0])
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		img_file, err := os.Open(os.Args[i])
		if err != nil {
			panic(err)
		}
		defer img_file.Close()

		img_processor(img_file)
	}
}
