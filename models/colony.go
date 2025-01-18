package models

import "fmt"

type Colony struct {
	StartRoom    Room
	StartFound   bool
	EndFound     bool
	EndRoom      Room
	NumberOfAnts uint64
	Rooms        map[string]*Room
	Ants         []Ant
	Output       []string
}


// NewColony initializes and returns a new Colony instance
func NewColony() *Colony {
	return &Colony{
		StartRoom:    Room{},
		StartFound:   false,
		EndFound:     false,
		EndRoom:      Room{},
		NumberOfAnts: 0,
		Rooms:        make(map[string]*Room),
		Ants:         make([]Ant, 0),
	}
}

// ConnectRooms connects two rooms in the colony by adding each room to the other's neighbours.
func (c *Colony) ConnectRooms(room1Name, room2Name string) error {
	room1, exists := c.Rooms[room1Name]
	if !exists {
		return fmt.Errorf("room %s not found", room1Name)
	}

	room2, exists := c.Rooms[room2Name]
	if !exists {
		return fmt.Errorf("room %s not found", room2Name)
	}

	room1.Neighbours = append(room1.Neighbours, room2)
	room2.Neighbours = append(room2.Neighbours, room1)
	return nil
}

// GetRoomByName retrieves a room by its name from the colony's Rooms map.
func (c *Colony) GetRoomByName(name string) *Room {
	if room, exists := c.Rooms[name]; exists {
		return room
	}
	return nil // Return nil if the room doesn't exist
}
