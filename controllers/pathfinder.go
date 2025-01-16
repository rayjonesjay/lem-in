package controllers

import (
	"fmt"

	"lemin/models"
)

func PathFinder(colony models.Colony) ([][]string, error) {
	var paths [][]string
	// check if startroom or end room is missing
	if !colony.StartFound || !colony.EndFound {
		return nil, fmt.Errorf("Missing start room or end room")
	}

	// check if both start room and end room are present
	// dfs function to find paths
	// visted to mark visited paths
	visited := make(map[string]bool)
	var path []string

	var dfs func(room *models.Room)
	dfs = func(room *models.Room) {
		// return if room is already visited
		if visited[room.Name] {
			return
		}
		// mark room as visited
		visited[room.Name] = true

		// add room to current path
		path = append(path, room.Name)
		fmt.Println(path)

		// add path to paths if end room is reached
		if room.Name == colony.EndRoom.Name {
			paths = append(paths, append([]string(nil), path...))
		} else {
			// explore the neighbors(connected rooms)
			for _, neighbor := range room.Neighbours {
				dfs(neighbor)
			}
		}

		// unmark the current room as visited
		visited[room.Name] = false
		// remove room from current path
		path = path[:len(path)-1]
	}

	// start dfs from start room
	dfs(&colony.StartRoom)

	fmt.Println(paths)
	return paths, nil
}

// func main() {
// 	// Manually creating a colony for testing
// 	colony := models.Colony{
// 		Rooms: make(map[string]*models.Room),
// 	}

// 	// Manually creating rooms
// 	startRoom := &models.Room{Name: "A", Neighbours: []*models.Room{}}
// 	endRoom := &models.Room{Name: "D", Neighbours: []*models.Room{}}
// 	roomB := &models.Room{Name: "B", Neighbours: []*models.Room{}}
// 	roomC := &models.Room{Name: "C", Neighbours: []*models.Room{}}

// 	// Connecting rooms manually (defining neighbors)
// 	startRoom.Neighbours = append(startRoom.Neighbours, roomB)
// 	roomB.Neighbours = append(roomB.Neighbours, startRoom, roomC)
// 	roomC.Neighbours = append(roomC.Neighbours, roomB, endRoom)
// 	endRoom.Neighbours = append(endRoom.Neighbours, roomC)

// 	// Adding rooms to the colony
// 	colony.Rooms["A"] = startRoom
// 	colony.Rooms["B"] = roomB
// 	colony.Rooms["C"] = roomC
// 	colony.Rooms["D"] = endRoom

// 	// Marking start and end rooms in the colony
// 	colony.StartRoom = *startRoom
// 	colony.EndRoom = *endRoom
// 	colony.StartFound = true
// 	colony.EndFound = true

// 	// Call PathFinder
// 	paths, err := PathFinder(colony)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Print the found paths
// 	fmt.Println("Paths found:", paths)
// }
