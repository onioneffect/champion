/*
	file.go: Handle file input and output.
	Copyright (C) 2021-2022  onioneffect

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

package champlib

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type LineWriteType struct {
	HexColor   [3]uint8
	Start, End [2]int32
}

func GenerateFilename(outDir, inputName string) (string, error) {
	_, err := os.Stat("outputs")
	if os.IsNotExist(err) {
		err = os.Mkdir("outputs", 0755)
		if err != nil {
			return "", err
		}
	}

	listing, err := ioutil.ReadDir("outputs")
	if err != nil {
		ChampLog(err)
	}

	// TODO: Handle file already exists.
	for i := range listing {
		if listing[i].Name() == inputName {
			return "", errors.New("file already exists")
		}
	}

	nameBase := filepath.Base(inputName)
	nameExt := filepath.Ext(inputName)
	trimmedFileName := strings.TrimSuffix(nameBase, nameExt)
	formattedOutput := fmt.Sprintf("%s-output.txt", trimmedFileName)

	return filepath.Join(outDir, formattedOutput), nil
}

func WriteLineSlicePlain(slicePtr *[]Line, fileName string) error {
	fp, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		return err
	}

	writeFunc := func(format string, argument Line) error {
		var formattedStr string

		str, err := argument.LineToString()
		if err != nil {
			return err
		}

		formattedStr = fmt.Sprintf(format, str)
		fp.WriteString(formattedStr)

		return nil
	}

	// Write the first argument
	writeFunc("[%s,", (*slicePtr)[0])

	// Then move onto the loop
	var f string
	for i := 1; i < len(*slicePtr); i++ {
		if i == len(*slicePtr)-1 {
			f = "%s]"
		} else {
			f = "%s,"
		}

		writeFunc(f, (*slicePtr)[i])
	}

	return nil
}

func WriteLineSliceEncoded(slicePtr *[]Line, fileName string) error {
	var WriteMe LineWriteType
	var LoopedLine Line

	fp, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		return err
	}

	for i := 0; i < len(*slicePtr); i++ {
		LoopedLine = (*slicePtr)[i]

		WriteMe = LineWriteType{
			Start: LoopedLine.Start,
			End:   LoopedLine.End,
		}

		WriteMe.HexColor = [3]uint8{
			uint8(LoopedLine.HexColor[0]),
			uint8(LoopedLine.HexColor[1]),
			uint8(LoopedLine.HexColor[2]),
		}

		err = binary.Write(fp, binary.LittleEndian, WriteMe)
		if err != nil {
			return err
		}
	}

	return nil
}
