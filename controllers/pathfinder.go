package controllers

import (
	"fmt"
	"sort"

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
	// sort the final paths to return a slice from the shortest to the longest(optimization)
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

	// sort slice
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	return paths, nil
}
