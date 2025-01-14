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