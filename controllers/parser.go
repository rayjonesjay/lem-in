package controllers

import (
	"bufio"
	"fmt"
	"os"
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

// ParseFile parses input from a file
func (p *Parser) ParseFile(filename string) (*models.Colony, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("empty input")
	}

	antCount, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil || antCount <= 0 {
		return nil, fmt.Errorf("invalid ant count: %s", scanner.Text())
	}
	p.colony.NumberOfAnts = antCount

	// Parse rooms and connections
	expectingStart := false
	expectingEnd := false

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case line == "##start":
			expectingStart = true
		case line == "##end":
			expectingEnd = true
		case strings.HasPrefix(line, "#"):
			continue
		case strings.Contains(line, "-"):
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

	if !p.colony.StartFound || !p.colony.EndFound {
		return nil, fmt.Errorf("missing start or end room")
	}

	return p.colony, nil
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
