package controllers

import (
	"reflect"
	"testing"

	"lemin/models"
)

// used helper funcions to create sample colonies for different test cases
// Test cases
// 1:colony with no start room
// 2: colony with no end room
// 3: colony with 1 path
// 4:colony with 2+ paths
// colony with no start or end rooms

// --Colony with no start room--//
func colonyNoStartRoom() models.Colony {
	colony := models.Colony{
		Rooms: make(map[string]*models.Room),
	}

	// Creating rooms
	startRoom := &models.Room{Name: "A", Neighbours: []*models.Room{}}
	endRoom := &models.Room{Name: "D", Neighbours: []*models.Room{}}
	roomB := &models.Room{Name: "B", Neighbours: []*models.Room{}}
	roomC := &models.Room{Name: "C", Neighbours: []*models.Room{}}
	// Connecting rooms manually
	startRoom.Neighbours = append(startRoom.Neighbours, roomB)
	roomB.Neighbours = append(roomB.Neighbours, startRoom, roomC)
	roomC.Neighbours = append(roomC.Neighbours, roomB, endRoom)
	endRoom.Neighbours = append(endRoom.Neighbours, roomC)

	// Adding rooms to the colony (no start room here)
	colony.Rooms["B"] = roomB
	colony.Rooms["C"] = roomC
	colony.Rooms["D"] = endRoom

	// marking start and end rooms
	colony.StartFound = false
	colony.EndFound = true

	return colony
}

// --Colony with No end Room --//
func colonyNoEndRoom() models.Colony {
	colony := models.Colony{
		Rooms: make(map[string]*models.Room),
	}

	// creating rooms
	// Creating rooms
	startRoom := &models.Room{Name: "A", Neighbours: []*models.Room{}}
	roomB := &models.Room{Name: "B", Neighbours: []*models.Room{}}
	roomC := &models.Room{Name: "C", Neighbours: []*models.Room{}}
	// Connecting rooms manually
	startRoom.Neighbours = append(startRoom.Neighbours, roomB)
	roomB.Neighbours = append(roomB.Neighbours, startRoom, roomC)

	// Adding rooms to the colony (no start room here)
	colony.Rooms["A"] = startRoom
	colony.Rooms["B"] = roomB
	colony.Rooms["C"] = roomC

	// marking start and end rooms
	colony.StartFound = true
	colony.EndFound = false

	return colony
}

// --Colony with only one path--//
func colonyOnePath() models.Colony {
	colony := models.Colony{
		Rooms: make(map[string]*models.Room),
	}

	// Creating rooms
	startRoom := &models.Room{Name: "A", Neighbours: []*models.Room{}}
	endRoom := &models.Room{Name: "D", Neighbours: []*models.Room{}}
	roomB := &models.Room{Name: "B", Neighbours: []*models.Room{}}
	roomC := &models.Room{Name: "C", Neighbours: []*models.Room{}}

	// Connecting rooms manually
	startRoom.Neighbours = append(startRoom.Neighbours, roomB)
	roomB.Neighbours = append(roomB.Neighbours, startRoom, roomC)
	roomC.Neighbours = append(roomC.Neighbours, roomB, endRoom)
	endRoom.Neighbours = append(endRoom.Neighbours, roomC)

	// Adding rooms to the colony (no start room here)
	colony.Rooms["A"] = startRoom
	colony.Rooms["B"] = roomB
	colony.Rooms["C"] = roomC
	colony.Rooms["D"] = endRoom

	// marking start and end rooms
	colony.StartRoom = *startRoom
	colony.EndRoom = *endRoom
	colony.StartFound = true
	colony.EndFound = true

	return colony
}

// --Colony with many paths --//
func colonyMultiplePaths() models.Colony {
	colony := models.Colony{
		Rooms: make(map[string]*models.Room),
	}

	// Creating rooms
	startRoom := &models.Room{Name: "A", Neighbours: []*models.Room{}}
	endRoom := &models.Room{Name: "D", Neighbours: []*models.Room{}}
	roomB := &models.Room{Name: "B", Neighbours: []*models.Room{}}
	roomC := &models.Room{Name: "C", Neighbours: []*models.Room{}}
	roomE := &models.Room{Name: "E", Neighbours: []*models.Room{}}
	roomF := &models.Room{Name: "F", Neighbours: []*models.Room{}}

	// Connecting rooms manually (defining neighbors)
	startRoom.Neighbours = append(startRoom.Neighbours, roomB, roomE)
	roomB.Neighbours = append(roomB.Neighbours, startRoom, roomC)
	roomC.Neighbours = append(roomC.Neighbours, roomB, endRoom)
	roomE.Neighbours = append(roomE.Neighbours, startRoom, roomF)
	roomF.Neighbours = append(roomF.Neighbours, roomE, endRoom)
	endRoom.Neighbours = append(endRoom.Neighbours, roomC, roomF)

	// Adding rooms to the colony
	colony.Rooms["A"] = startRoom
	colony.Rooms["B"] = roomB
	colony.Rooms["C"] = roomC
	colony.Rooms["D"] = endRoom
	colony.Rooms["E"] = roomE
	colony.Rooms["F"] = roomF

	// Marking start and end rooms
	colony.StartRoom = *startRoom
	colony.EndRoom = *endRoom
	colony.StartFound = true
	colony.EndFound = true

	return colony
}

func TestPathFinder(t *testing.T) {
	tests := []struct {
		name    string
		args    models.Colony
		want    [][]string
		wantErr bool
	}{
		{
			name:    "Test: with no start room",
			args:    colonyNoStartRoom(),
			want:    nil,
			wantErr: true,
		},

		{
			name:    "Test with no end room",
			args:    colonyNoEndRoom(),
			want:    nil,
			wantErr: true,
		},

		{
			name:    "Test with one path",
			args:    colonyOnePath(),
			want:    [][]string{{"A", "B", "C", "D"}},
			wantErr: false,
		},
		{
			name:    "Test with multiple paths",
			args:    colonyOnePath(),
			want:    [][]string{{"A", "B", "C", "D"},{"A", "E", "F", "D"}},
			wantErr: false,
		},
	}

	// Running the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathFinder(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathFinder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
