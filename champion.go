// champion.go
package main

import (
	"flag"
	"log"
	"os"

	champlib "github.com/onioneffect/champion/lib"
)

func tryLogOutputStr(path string) {
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

func imgProcessor(fp *os.File, debug bool) {
	var currentImg champlib.ImageInfo = champlib.ReadImgInfo(fp)
	var currentDecoded [][][3]int32 = champlib.ImageArray(currentImg)
	currentImg.Decoded = &currentDecoded

	if debug {
		log.Println("START")
		log.Println("We are in debuggign mode!!! :D")

		log.Println("Printing image information:")
		champlib.LogImgInfo(currentImg)

		log.Println("Printing array information:")
		champlib.LogIntarrayInfo(currentImg.Decoded)

		log.Println("Running TestPixLoop:")
		champlib.TestPixLoop(currentImg, 100)

		log.Println("DONE")
	}
}

func main() {
	var useDebugging bool
	var logOutputStr string

	flag.BoolVar(&useDebugging, "debug", false, "enable logging")
	flag.StringVar(&logOutputStr, "file", "", "specify where to write log output")
	flag.Parse()

	// Run this function outside of the loop, so it only runs once.
	// It's not like anyone is going to use multiple log files anyway.
	if useDebugging {
		log.Println("Calling tryLogOutputStr...")
		tryLogOutputStr(logOutputStr)
	}

	for i := 0; i < flag.NArg(); i++ {
		imgFile, err := os.Open(flag.Args()[i])
		if err != nil {
			log.Printf("ERROR (image file): %s\n", err)
			continue
		}

		if useDebugging {
			log.Println("Successfully opened file", flag.Args()[i])
		}
		imgProcessor(imgFile, useDebugging)
		imgFile.Close()
	}
}
