package models

import (
	"testing"
)

func TestColony(t *testing.T) {
	// Create a new Colony
	colony := NewColony()

	// Test that the Colony is initialized correctly
	if colony.StartFound != false {
		t.Errorf("Expected StartFound to be false, got %v", colony.StartFound)
	}
	if colony.EndFound != false {
		t.Errorf("Expected EndFound to be false, got %v", colony.EndFound)
	}
	if colony.NumberOfAnts != 0 {
		t.Errorf("Expected NumberOfAnts to be 0, got %d", colony.NumberOfAnts)
	}
	if len(colony.Rooms) != 0 {
		t.Errorf("Expected Rooms map to be empty, got %d rooms", len(colony.Rooms))
	}
	if len(colony.Ants) != 0 {
		t.Errorf("Expected Ants slice to be empty, got %d ants", len(colony.Ants))
	}

	// Create and add rooms to the Colony
	room1 := &Room{Name: "room1", Neighbours: []string{}}
	room2 := &Room{Name: "room2", Neighbours: []string{}}
	colony.Rooms[room1.Name] = room1
	colony.Rooms[room2.Name] = room2

	// Test that rooms have been added to the Rooms map
	if len(colony.Rooms) != 2 {
		t.Errorf("Expected 2 rooms in the colony, got %d", len(colony.Rooms))
	}

	// Test ConnectRooms function
	err := colony.ConnectRooms("room1", "room2")
	if err != nil {
		t.Errorf("Error connecting rooms: %v", err)
	}

	// Verify that rooms are connected
	if !contains(room1.Neighbours, "room2") {
		t.Errorf("Expected room1 to have room2 as a neighbour, got %v", room1.Neighbours)
	}
	if !contains(room2.Neighbours, "room1") {
		t.Errorf("Expected room2 to have room1 as a neighbour, got %v", room2.Neighbours)
	}

	// Test error handling when trying to connect non-existing rooms
	err = colony.ConnectRooms("room1", "room3")
	if err == nil {
		t.Error("Expected error when connecting non-existing room, got nil")
	}

	// Test GetRoomByName function
	retrievedRoom := colony.GetRoomByName("room1")
	if retrievedRoom == nil {
		t.Error("Expected to retrieve room1, but got nil")
	}
	if retrievedRoom.Name != "room1" {
		t.Errorf("Expected room1, got %v", retrievedRoom)
	}

	// Test GetRoomByName for a non-existing room
	retrievedRoom = colony.GetRoomByName("room3")
	if retrievedRoom != nil {
		t.Error("Expected nil for non-existing room, got a room")
	}
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
