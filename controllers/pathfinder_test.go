package controllers

import (
	"reflect"
	"testing"

	"lemin/models"
)

func TestPathFinder(t *testing.T) {
	// Creating a sample colony
	colony := models.Colony{
		Rooms: make(map[string]*models.Room),
	}

	// create rooms
	// Manually creating rooms
	startRoom := &models.Room{Name: "A", Neighbours: []*models.Room{}}
	endRoom := &models.Room{Name: "D", Neighbours: []*models.Room{}}
	roomB := &models.Room{Name: "B", Neighbours: []*models.Room{}}
	roomC := &models.Room{Name: "C", Neighbours: []*models.Room{}}

	// Connecting rooms manually (defining neighbors)
	startRoom.Neighbours = append(startRoom.Neighbours, roomB)
	roomB.Neighbours = append(roomB.Neighbours, startRoom, roomC)
	roomC.Neighbours = append(roomC.Neighbours, roomB, endRoom)
	endRoom.Neighbours = append(endRoom.Neighbours, roomC)

	// Adding rooms to the colony
	colony.Rooms["A"] = startRoom
	colony.Rooms["B"] = roomB
	colony.Rooms["C"] = roomC
	colony.Rooms["D"] = endRoom

	// Marking start and end rooms in the colony
	colony.StartRoom = *startRoom
	colony.EndRoom = *endRoom
	colony.StartFound = true
	colony.EndFound = true

	tests := []struct {
		name    string
		args    models.Colony
		want    [][]string
		wantErr bool
	}{
		{
			name: "Simple test case with one path",
			args: colony,
			want: [][]string{
				{"A", "B", "C", "D"},
			},
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
