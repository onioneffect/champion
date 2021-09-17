// champion.go
package main

import (
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
	var allFiles []string
	var allFilesCtr int = 0
	var useDebugging bool = false
	var logOutputStr string
	var curr string

	for i := 1; i < len(os.Args); i++ {
		curr = os.Args[i]

		if curr == "--debug" {
			useDebugging = true
		} else if curr == "--file" {
			// We increment i so it points to the argument right
			// after "--file", and so the next iteration of the
			// loop doesn't include it in the allFiles list.
			i++
			logOutputStr = os.Args[i]
		} else {
			allFiles = append(allFiles, curr)
			allFilesCtr++
		}
	}

	// Run this function outside of the loop, so it only runs once.
	// It's not like anyone is going to use multiple log files anyway.
	tryLogOutputStr(logOutputStr)

	for i := 0; i < allFilesCtr; i++ {
		imgFile, err := os.Open(allFiles[i])
		if err != nil {
			log.Printf("ERROR (image file): %s\n", err)
			continue
		}

		log.Println("Successfully opened file", allFiles[i])
		imgProcessor(imgFile, useDebugging)
		imgFile.Close()
	}
}
