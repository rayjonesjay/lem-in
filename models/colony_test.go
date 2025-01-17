package models

import (
	"testing"
)

func TestColony(t *testing.T) {
	// Test colony creation
	colony := NewColony()

	// Test initial state
	if len(colony.Rooms) != 0 {
		t.Error("Expected empty rooms map in new colony")
	}

	if len(colony.Ants) != 0 {
		t.Error("Expected empty ants slice in new colony")
	}

	if colony.StartFound {
		t.Error("Expected StartFound to be false in new colony")
	}

	if colony.EndFound {
		t.Error("Expected EndFound to be false in new colony")
	}

	if colony.NumberOfAnts != 0 {
		t.Error("Expected NumberOfAnts to be 0 in new colony")
	}

	// Test room connections
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

	colony.Rooms["room1"] = room1
	colony.Rooms["room2"] = room2

	// Test connecting rooms
	err := colony.ConnectRooms("room1", "room2")
	if err != nil {
		t.Errorf("Failed to connect rooms: %v", err)
	}

	if len(room1.Neighbours) != 1 || room1.Neighbours[0] != room2 {
		t.Error("Room1 not properly connected to room2")
	}

	if len(room2.Neighbours) != 1 || room2.Neighbours[0] != room1 {
		t.Error("Room2 not properly connected to room1")
	}

	// Test connecting non-existent rooms
	err = colony.ConnectRooms("room1", "nonexistent")
	if err == nil {
		t.Error("Expected error when connecting to nonexistent room")
	}

	// Test start and end room setup
	startRoom := &Room{
		Name: "start",
		Coordinate: Coordinate{
			X: 0.0,
			Y: 0.0,
		},
		IsStart:    true,
		Neighbours: make([]*Room, 0),
	}

	endRoom := &Room{
		Name: "end",
		Coordinate: Coordinate{
			X: 5.0,
			Y: 5.0,
		},
		IsEnd:      true,
		Neighbours: make([]*Room, 0),
	}

	colony.StartRoom = *startRoom
	colony.EndRoom = *endRoom
	colony.StartFound = true
	colony.EndFound = true

	if !colony.StartFound || !colony.EndFound {
		t.Error("Start or end room not properly set")
	}

	if colony.StartRoom.Name != "start" || colony.EndRoom.Name != "end" {
		t.Error("Start or end room names not properly set")
	}
}