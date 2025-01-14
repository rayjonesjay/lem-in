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
