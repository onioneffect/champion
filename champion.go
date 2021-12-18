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

	champlib.ChampLog("START")
	champlib.ChampLog("We are in debuggign mode!!! :D")

	champlib.ChampLog("Printing image information:")
	champlib.LogImgInfo(currentImg)

	champlib.ChampLog("Printing array information:")
	champlib.LogIntarrayInfo(currentImg.Decoded)

	champlib.ChampLog("Running TestPixLoop:")
	champlib.TestPixLoop(currentImg, 100)

	champlib.ChampLog("DONE")
}

func main() {
	var logOutputStr string
	var sliceLength int

	flag.BoolVar(&champlib.LoggingEnabled, "debug", false, champlib.HelpDebug)
	flag.StringVar(&logOutputStr, "file", "", champlib.HelpFile)
	flag.IntVar(&sliceLength, "slice", 1024, champlib.HelpSlice)
	flag.Parse()

	// Run this function outside of the loop, so it only runs once.
	// It's not like anyone is going to use multiple log files anyway.
	if champlib.LoggingEnabled {
		log.Println("Calling tryLogOutputStr...")
		tryLogOutputStr(logOutputStr)
	}

	for i := 0; i < flag.NArg(); i++ {
		imgFile, err := os.Open(flag.Args()[i])
		if err != nil {
			log.Printf("ERROR (image file): %s\n", err)
			continue
		}

		champlib.ChampLog("Successfully opened file ", flag.Args()[i])
		imgProcessor(imgFile, champlib.LoggingEnabled)
		imgFile.Close()
	}
}
