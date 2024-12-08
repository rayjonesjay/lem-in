package parse

import (
	"bufio"
	"os"
	"strings"

	"lem-in/xerrors"
)

// ReadValidateFileContents checks if the file is a .txt file - meaning it contains only ascii characters, if runes or special symbols are encountered
// a nil/empty slice of strings and an ErrInvalidFileContents will be returned. if success it will return file contents as a slice of strings and nil as the error value
func ReadValidateFileContents(filename string) (fileContents []string, err error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// create a scanner
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {

		currentLine := scanner.Text()

		currentLine = strings.TrimSpace(currentLine)

		// ignore empty lines
		if currentLine == "" {
			continue
		}

		fileContents = append(fileContents, currentLine)

	}

	if len(fileContents) == 0 {
		return nil, xerrors.ErrEndNotFound
	}

	return
}
