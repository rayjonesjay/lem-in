package controllers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"lemin/models"
)

type Parser struct {
	colony *models.Colony
}

func NewParser() *Parser {
	return &Parser{
		colony: models.NewColony(),
	}
}

// ParseFile accepts a file name and internally calls ReadValidateInputFile to get the contents of the file
func (p *Parser) ParseFile(filename string) (*models.Colony, error) {

	fileContents, err := ReadValidateInputFile(filename)
	if err != nil {
		return nil, err
	}

	// store the read input to the Output field, in order to print it out later
	p.colony.Output = fileContents

	// First Line: Number of ants are at index 0 of the slice since we assume it was read first
	numberOfAnts, err := strconv.Atoi(string(fileContents[0]))

	if err != nil || numberOfAnts < 1 {
		return nil, errors.New("invalid number of ants")
	}
	p.colony.NumberOfAnts = uint64(numberOfAnts)

	// Parse rooms and connections
	expectingStart := false
	expectingEnd := false

	// starting off at index 1 since we have already found the number of ants no need to check it
	for _, line := range fileContents[1:] {
		switch {
		case line == "##start":
			expectingStart = true

		case line == "##end":
			expectingEnd = true

			//if the line starts with # and we have already found the end and start
		case strings.HasPrefix(line, "#"):
			if expectingStart && expectingEnd {
				continue
			}
			continue

		case isLink(line):
			err := p.parseConnection(line)
			if err != nil {
				return nil, err
			}
		default:
			err := p.parseRoom(line, expectingStart, expectingEnd)
			if err != nil {
				return nil, err
			}
			expectingStart = false
			expectingEnd = false
		}
	}

	if !p.colony.StartFound || !p.colony.EndFound || !expectingStart || !expectingEnd {
		return nil, fmt.Errorf("missing start or end room")
	}

	return p.colony, nil
}

// isLink checks if the s is separated by a hyphen and returns true else false
func isLink(s string) bool {
	pattern := `^.+\-.+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

func (p *Parser) parseRoom(line string, isStart, isEnd bool) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("invalid room format: %s", line)
	}

	name := parts[0]
	x, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return fmt.Errorf("invalid x coordinate: %s", parts[1])
	}

	y, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return fmt.Errorf("invalid y coordinate: %s", parts[2])
	}

	room := &models.Room{
		Name: name,
		Coordinate: models.Coordinate{
			X: x,
			Y: y,
		},
		Neighbours: make([]*models.Room, 0),
	}

	if isStart {
		room.IsStart = true
		p.colony.StartRoom = *room
		p.colony.StartFound = true
	} else if isEnd {
		room.IsEnd = true
		p.colony.EndRoom = *room
		p.colony.EndFound = true
	}

	p.colony.Rooms[name] = room
	return nil
}

func (p *Parser) parseConnection(line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid connection format: %s", line)
	}

	return p.colony.ConnectRooms(parts[0], parts[1])
}
