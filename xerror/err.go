package xerrors

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrNoDataFound = errors.New("ERROR: file contains no data")

	ErrInvalidNumberOfAnts = errors.New("ERROR: number of ants invalid")

	ErrInvalidRoomCoordinates = errors.New("ERROR: wrong coordinates detected")

	ErrInvalidDataFormat = errors.New("ERROR: invalid data format")

	ErrInvalidLink = errors.New("ERROR: wrong link format: %s")

	ErrDuplicateRoom = errors.New("ERROR: duplicat rooms found")

	ErrWrongXCoord = errors.New("ERROR: wrong x coordinates detected")

	ErrWrongYCoord = errors.New("ERROR: wrong y coordinates detected")
)

func ErrorWriter(err error, exitCode int, shouldExit bool) {
	fmt.Println(err)
	if shouldExit {
		os.Exit(exitCode)
	}
}
