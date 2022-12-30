/*
	champion.go: Command line parsing.
	Copyright (C) 2021-2023  onioneffect

	This file is part of Champion.

	Champion is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Champion is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with Champion.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"os"

	champlib "github.com/onioneffect/champion/lib"
)

type Settings struct {
	SliceLen int
}

func imgProcessor(fp *os.File) {
	currentImg, err := champlib.ReadImgInfo(fp)
	if errors.Is(err, image.ErrFormat) {
		msg := fmt.Sprintf("Unknown format on file `%s`! Skipping...", fp.Name())
		fmt.Print(msg)
		champlib.ChampLog(msg)
		return
	}

	var currentDecoded [][][3]int32 = champlib.ImageArray(currentImg)
	currentImg.Decoded = &currentDecoded

	champlib.ChampLog("START")

	champlib.ChampLog("Printing image information:")
	champlib.LogImgInfo(currentImg)

	champlib.ChampLog("Printing array information:")
	champlib.LogIntarrayInfo(currentImg.Decoded)

	champlib.ChampLog("Running main loop!")
	mySlice := champlib.ImagePixLoop(currentImg)

	champlib.ChampLog("Calling DebugLineSlice:")
	champlib.DebugLineSlice(mySlice, true)

	// TODO: Custom output dir
	champlib.ChampLog("Generating output filename:")
	s, err := champlib.GenerateFilename("outputs", fp.Name())
	if err != nil {
		panic(err)
	}
	champlib.ChampLog(s)

	champlib.ChampLog("Calling WriteLineSlice:")
	err = champlib.WriteLineSlicePlain(&mySlice, s)
	if err != nil {
		panic(err)
	}

	s2 := fmt.Sprintf("%s.bin", s)
	champlib.ChampLog("Calling WriteLineSliceEncoded:")
	err = champlib.WriteLineSliceEncoded(&mySlice, s2)
	if err != nil {
		panic(err)
	}

	champlib.ChampLog("DONE")
}

func main() {
	var logOutputStr string

	flag.BoolVar(&champlib.LoggingEnabled, "debug", false, champlib.HelpDebug)
	flag.StringVar(&logOutputStr, "file", "", champlib.HelpFile)
	flag.Parse()

	// Run this function outside of the loop, so it only runs once.
	// It's not like anyone is going to use multiple log files anyway.
	if champlib.LoggingEnabled {
		champlib.ChampLog("Calling TryLogOutputStr...")
		champlib.TryLogOutputStr(logOutputStr)
	}

	for i := 0; i < flag.NArg(); i++ {
		imgFile, err := os.Open(flag.Args()[i])
		if err != nil {
			msg := fmt.Sprintf("ERROR (image file): %s\n", err)
			champlib.ChampLog(msg)
			continue
		}

		champlib.ChampLog("Successfully opened file ", flag.Args()[i])
		imgProcessor(imgFile)
		imgFile.Close()
	}
}
