/*
	file.go: Handle file input and output.
	Copyright (C) 2021  onioneffect

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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFilename(outDir string, inputName string) (string, error) {
	_, err := os.Stat("outputs")
	if os.IsNotExist(err) {
		err = os.Mkdir("outputs", 0755)
		if err != nil {
			ChampLog(err)
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

func WriteLineSlice(slicePtr *[]Line, fileName string) error {
	var formattedStr string

	fp, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		return err
	}

	for i, j := range *slicePtr {
		str, err := j.LineToString()
		if err != nil {
			return err
		}

		if i == 0 {
			formattedStr = fmt.Sprintf("[%s,", str)
		} else if i == len(*slicePtr)-1 {
			formattedStr = fmt.Sprintf("%s]", str)
		} else {
			formattedStr = fmt.Sprintf("%s, ", str)
		}
		fp.WriteString(formattedStr)
	}

	return nil
}
