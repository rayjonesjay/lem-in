package parse

import (
	"bufio"
	"os"
	"strings"
)

// ReadValidateFileContents checks if the file is a .txt file - meaning it contains only ascii characters, if runes or special symbols are encountered
// a nil/empty slice of strings and an ErrInvalidFileContents will be returned. if success it will return file contents as a slice of strings and nil as the error value
func ReadValidateFileContents(filename string) ([]string, error) {

	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// scan
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {

		currentLine := scanner.Text()

		currentLine = strings.TrimSpace(currentLine)
		if currentLine == "" {
			continue
		}

		// helper function to check the characters on each line
	}

}

func helper()
