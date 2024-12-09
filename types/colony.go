package types

import (
	"fmt"
	"strconv"
	"strings"

	"lem-in/xerrors"
)

const (
	MaxAntsPerColony = 1_000
)

// a room looks like A 1 2 where A is the name, and 1 and 2 is the x and y cordinate respectively
// although this cordinates will be used in visualization
type Room struct {
	Name       string
	X, Y       int      // room cordinates
	Neighbours []string // each room has neighbours
	Occupied   bool     // indicator to show room occupation status
}

// Colony is a model of the ant farm also known as colony
type Colony struct {
	StartRoom    Room
	StartFound   bool
	EndFound     bool
	EndRoom      Room
	NumberOfAnts uint64           // number of ants cannot be negative
	Rooms        map[string]*Room // rooms is a slice of rooms
}

// CheckNumAnts checks the number of ants per colony if they have exceeded MaxAntsPerColony
// and exits with status code 1
func CheckNumAnts(c Colony) error {
	if c.NumberOfAnts >= MaxAntsPerColony {
		return (xerrors.ErrMaxAntNumExceeded)
	}
	if c.NumberOfAnts == 0 {
		return xerrors.ErrZeroAnts
	}
	return nil
}

func ParseFileContentsToColony(fileContents []string) (*Colony, error) {
	colony := &Colony{
		Rooms: make(map[string]*Room),
	}

	// First line: Number of ants
	numberOfAnts, err := strconv.Atoi(strings.TrimSpace(fileContents[0]))
	if err != nil || numberOfAnts < 0 {
		return nil, xerrors.ErrInvalidNumberOfAnts
	}
	colony.NumberOfAnts = uint64(numberOfAnts)

	// Helper function to parse room name and coordinates
	parseRoom := func(s string) (name string, x, y int, err error) {
		parts := strings.Fields(strings.TrimSpace(s))
		if len(parts) != 3 {
			return "", 0, 0, xerrors.ErrInvalidRoomFormat
		}

		name = parts[0]
		x, err = strconv.Atoi(parts[1])
		if err != nil {
			return "", 0, 0, xerrors.ErrInvalidRoomCoordinates
		}
		y, err = strconv.Atoi(parts[2])
		if err != nil {
			return "", 0, 0, xerrors.ErrInvalidRoomCoordinates
		}
		return name, x, y, nil
	}

	// Check for ##start and ##end directives
	i := 1
	if strings.TrimSpace(fileContents[i]) == "##start" {
		colony.StartFound = true
		i++
		name, x, y, err := parseRoom(fileContents[i])
		if err != nil {
			return nil, err
		}
		colony.StartRoom = Room{Name: name, X: x, Y: y}
		colony.Rooms[name] = &colony.StartRoom
		i++
	} else {
		return nil, xerrors.ErrStartNotFound
	}

	endRoomSet := false
	for ; i < len(fileContents); i++ {
		line := strings.TrimSpace(fileContents[i])

		if line == "##end" && !endRoomSet {
			colony.EndFound = true
			i++
			name, x, y, err := parseRoom(fileContents[i])
			if err != nil {
				return nil, err
			}
			colony.EndRoom = Room{Name: name, X: x, Y: y}
			colony.Rooms[name] = &colony.EndRoom
			endRoomSet = true
			continue
		}

		// Parse regular rooms
		if !strings.HasPrefix(line, "#") && line != "" {
			name, x, y, err := parseRoom(line)
			if err != nil {
				return nil, err
			}
			if _, exists := colony.Rooms[name]; exists {
				return nil, fmt.Errorf("%s", xerrors.ErrDuplicateRoom, colony.Rooms[name])
			}
			colony.Rooms[name] = &Room{Name: name, X: x, Y: y}
		}
	}

	// Validate if start and end rooms were found
	if !colony.StartFound || colony.StartRoom.Name == "" {
		return nil, xerrors.ErrStartNotFound
	}
	if !colony.EndFound || colony.EndRoom.Name == "" {
		return nil, xerrors.ErrEndNotFound
	}

	return colony, nil
}
