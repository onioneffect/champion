package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type ImageInfo struct {
	Format        string
	CModel        color.Model
	Width, Height int
	Data          *image.Image
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
	return_info.CModel = config.ColorModel

	_, err = img_reader.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	img_data, _, err := image.Decode(img_reader)
	if err != nil {
		panic(err)
	}
	return_info.Data = &img_data

	return return_info
}

func (imginf ImageInfo) print_img_info() {
	fmt.Printf("Image dimensions: %d, %d\n", imginf.Width, imginf.Height)
	fmt.Println("Image format:", imginf.Format)

	// Makes CModel convert an empty color.
	// Returns the corresponding color model.
	// Thanks to https://stackoverflow.com/questions/45226991/
	fmt.Printf("Image color mode: %T\n", imginf.CModel.Convert(color.RGBA{}))

	fmt.Println("First pixel:", (*imginf.Data).At(0, 0))
}

func img_processor(fp *os.File) {
	var currentImg ImageInfo = read_img_info(fp)
	currentImg.print_img_info()
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
