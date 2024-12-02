// package xerrors contains error types and error handling functions for efficient error handling
package xerrors

import "errors"

// all error variables must have the perfix Err
// ErrStartNotFound and ErrEndNotFound will have exit status of 3
// ErrNoArgsPassed exits with status 1
// ErrInvalidFileContents exits with status 4
var (
	// returned when ##start instruction is missing in input file, exit with status 3
	ErrStartNotFound = errors.New("StartNotFound: %s does not contain ##start instruction")

	// returned when ##end istruction is missing in input file, exit with status 3
	ErrEndNotFound = errors.New("EndNotFound: %s does not contain ##end instruction")

	// returned when no arguments have been passed through the commandline, exit status 1
	ErrNoArgsPassed = errors.New("NoArgsPassed: no arguments passed")

	// returned when a file contains invalid characters such as emojis,characters whose ascii value > 127, values which are not letters , # or numbers.
	ErrInvalidFileContents = errors.New("ERROR: InvalidFileContents: %s contains invalid file contents, valid contents include numbers[0-9],letters[A-Z], #, and -")
)
