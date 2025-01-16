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
		Neighbours: make([]*Room, 0),
	}

	// Create ant
	ant := NewAnt(1, startRoom)

	// Test initial state
	if ant.ID != 1 {
		t.Errorf("Expected ant ID to be 1, got %d", ant.ID)
	}

	if ant.Position != startRoom {
		t.Error("Expected ant to be in start room")
	}

	// Create rooms for path testing
	room1 := &Room{
		Name: "room1",
		Coordinate: Coordinate{
			X: 1.0,
			Y: 1.0,
		},
		Neighbours: make([]*Room, 0),
	}

	room2 := &Room{
		Name: "room2",
		Coordinate: Coordinate{
			X: 2.0,
			Y: 2.0,
		},
		Neighbours: make([]*Room, 0),
	}

	// Test movement along path
	path := []*Room{startRoom, room1, room2}
	ant.SetPath(path)
	
	// Test first move
	movedRoom := ant.Move()
	if movedRoom != room1 {
		t.Error("Expected ant to move to room1")
	}

	// Test second move
	movedRoom = ant.Move()
	if movedRoom != room2 {
		t.Error("Expected ant to move to room2")
	}

	// Test movement beyond path
	movedRoom = ant.Move()
	if movedRoom != nil {
		t.Error("Expected no more moves after reaching end of path")
	}

	// Test GetNextRoom
	ant.SetPath(path) // Reset path
	nextRoom := ant.GetNextRoom()
	if nextRoom != room1 {
		t.Error("GetNextRoom should return next room without moving ant")
	}
}