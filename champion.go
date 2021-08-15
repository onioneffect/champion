package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

type ImageInfo struct {
	Format        string
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

func print_img_info(imginf ImageInfo) {
	fmt.Printf("Image dimensions: %d, %d\n", imginf.Width, imginf.Height)
	fmt.Println("Image format:", imginf.Format)
	fmt.Println("First pixel:", (*imginf.Data).At(0, 0))
}

func img_processor(fp *os.File) {
	var current_img ImageInfo = read_img_info(fp)
	print_img_info(current_img)
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		img_file, err := os.Open(os.Args[i])
		if err != nil {
			panic(err)
		}

		img_processor(img_file)
		img_file.Close()
	}
}
