// package xerrors contains error types and error handling functions for efficient error handling
package xerrors

import (
	"errors"
	"fmt"
	"os"
)

// all error variables must have the prefix Err
// ErrStartNotFound and ErrEndNotFound will have exit status of 3
// ErrNoArgsPassed exits with status 1
// ErrInvalidFileContents exits with status 4
// ErrMaxAntNumExceeded exits with status 1
var (
	// ErrStartNotFound returned when ##start instruction is missing in input file, exit with status 3
	ErrStartNotFound = errors.New("StartNotFound: %s does not contain ##start instruction")

	// ErrEndNotFound returned when ##end instruction is missing in input file, exit with status 3
	ErrEndNotFound = errors.New("EndNotFound: %s does not contain ##end instruction")

	// ErrNoArgsPassed returned when no arguments have been passed through the commandline, exit status 1
	ErrNoArgsPassed = errors.New("NoArgsPassed: no arguments passed")

	// ErrInvalidFileContents returned when a file contains invalid characters such as emojis,characters whose ascii value > 127, values which are not letters , # or numbers.
	ErrInvalidFileContents = errors.New("InvalidFileContents: %s contains invalid file contents, valid contents include numbers[0-9],letters[A-Z], #, and - ")

	// ErrMaxAntNumExceeded returned when ant numbers surpass a given value, default is 1000 ants
	ErrMaxAntNumExceeded = errors.New("MaxAntNumExceeded: number of ants exceeded try lower than 1000")

	// ErrZeroAnts returned when the number of ants is 0
	ErrZeroAnts = errors.New("ZeroAnts: no ants in the colony")
)

func ErrorWriter(err error, exitCode int, shouldExit bool) {
	fmt.Println(err)
	if shouldExit {
		os.Exit(exitCode)
	}
}
