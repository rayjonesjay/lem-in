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