package controllers

import (
	"lemin/models"
)

type Parser struct {
	rooms   map[string]*models.Room
	tunnels []*models.Tunnel
}


func NewParser() *Parser {
	return &Parser{
		rooms:   make(map[string]*models.Room),
		tunnels: make([]*models.Tunnel, 0),
	}
}

func (p *Parser) ParseFile(filename string) (*models.Farm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("empty input file")
	}

	antCount, err := strconv.Atoi(scanner.Text())
	if err != nil || antCount <= 0 {
		return nil, fmt.Errorf("invalid ant count: %s", scanner.Text())
	}

	// Parse rooms and tunnels
	var startRoom, endRoom *models.Room
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
			err := p.parseTunnel(line)
			if err != nil {
				return nil, err
			}
		default:
			room, err := p.parseRoom(line)
			if err != nil {
				return nil, err
			}

			if expectingStart {
				startRoom = room
				room.SetAsStart()
				expectingStart = false
			} else if expectingEnd {
				endRoom = room
				room.SetAsEnd()
				expectingEnd = false
			}
		}
	}

	if startRoom == nil || endRoom == nil {
		return nil, fmt.Errorf("missing start or end room")
	}

	return &models.Farm{
		AntCount:  antCount,
		Rooms:     p.rooms,
		Tunnels:   p.tunnels,
		StartRoom: startRoom,
		EndRoom:   endRoom,
	}, nil
}

func (p *Parser) parseRoom(line string) (*models.Room, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid room format: %s", line)
	}

	name := parts[0]
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return nil, fmt.Errorf("invalid room name: %s", name)
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid x coordinate: %s", parts[1])
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid y coordinate: %s", parts[2])
	}

	if _, exists := p.rooms[name]; exists {
		return nil, fmt.Errorf("duplicate room name: %s", name)
	}

	room := models.NewRoom(name, x, y)
	p.rooms[name] = room
	return room, nil
}

func (p *Parser) parseTunnel(line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid tunnel format: %s", line)
	}

	room1, exists := p.rooms[parts[0]]
	if !exists {
		return fmt.Errorf("unknown room in tunnel: %s", parts[0])
	}

	room2, exists := p.rooms[parts[1]]
	if !exists {
		return fmt.Errorf("unknown room in tunnel: %s", parts[1])
	}

	if room1 == room2 {
		return fmt.Errorf("self-loop detected: %s", line)
	}

	// Check for duplicate tunnels
	for _, t := range p.tunnels {
		if (t.Room1 == room1 && t.Room2 == room2) ||
			(t.Room1 == room2 && t.Room2 == room1) {
			return fmt.Errorf("duplicate tunnel: %s", line)
		}
	}

	p.tunnels = append(p.tunnels, models.NewTunnel(room1, room2))
	return nil
}