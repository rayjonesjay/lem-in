package controllers

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"unicode"
)

// ReadValidateInputFile reads the data from the file passed as a command line argument
// It checks if the file contains only ascii characters
// It checks if the file passed exists
// It returns err if any of the checks fail
func ReadValidateInputFile(filename string) (fileContents []string, err error) {

	fd, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		currentLine := scanner.Text()

		// check if the read line contains a character that is not ASCII and ignores it
		if !ContainsASCII(currentLine) {
			continue
		}

		// remove any trailing or leading white spaces
		currentLine = strings.TrimSpace(currentLine)

		// ignore empty line
		if currentLine == "" {
			continue
		}

		fileContents = append(fileContents, currentLine)
	}

	if len(fileContents) == 0 {
		return nil, errors.New("no valid input file found")
	}

	return
}

func ContainsASCII(s string) bool {
	for _, c := range s {
		// checks if c is greater than 127
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}
