package models

import (
	"testing"
)

func TestAnt(t *testing.T) {
	// Create start room
	startRoom := &Room{
		Name: "start",
		Coordinate: Coordinate{
			X: 0.0,
			Y: 0.0,
		},
		IsStart:    true,
		Neighbours: []string{}, // Neighbours should be a slice of room names (strings)
	}

	// Create ant
	ant := NewAnt(1, startRoom)

	// Test initial state
	if ant.ID != 1 {
		t.Errorf("Expected ant ID to be 1, got %d", ant.ID)
	}

	// Test that the ant's position matches the name of the start room
	if ant.Position == startRoom.Name { // Compare with the room's name
		t.Error("Expected ant to be in start room")
	}

	// Create rooms for path testing
	room1 := &Room{
		Name: "room1",
		Coordinate: Coordinate{
			X: 1.0,
			Y: 1.0,
		},
		// Neighbours should be names of rooms (strings)
	}

	room2 := &Room{
		Name: "room2",
		Coordinate: Coordinate{
			X: 2.0,
			Y: 2.0,
		},
		// Neighbours should be names of rooms (strings)
	}

	// Test movement along path
	path := []string{startRoom.Name, room1.Name, room2.Name} // Use names of rooms
	ant.SetPath(path)

	// Test first move
	movedRoom := ant.Move()
	if movedRoom != room1.Name { // Compare with room names
		t.Error("Expected ant to move to room1")
	}

	// Test second move
	movedRoom = ant.Move()
	if movedRoom != room2.Name { // Compare with room names
		t.Error("Expected ant to move to room2")
	}

	// Test movement beyond path
	movedRoom = ant.Move()
	if movedRoom != "" { // Check for empty string when path is exhausted
		t.Error("Expected no more moves after reaching end of path")
	}

	// Test GetNextRoom
	ant.SetPath(path) // Reset path
	nextRoom := ant.GetNextRoom()
	if nextRoom != room1.Name { // Check using room names
		t.Error("GetNextRoom should return next room without moving ant")
	}
}
