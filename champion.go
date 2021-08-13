package main

import (
	"fmt"
	//"image/png"
	"os"
)

func main() {
	fmt.Println("Hello!")

	for i := 0; i < len(os.Args); i++ {
		imgFile, err := os.Open(os.Args[i])
		if err != nil {
			panic(err)
		}
		defer imgFile.Close()
	}

	return
}
