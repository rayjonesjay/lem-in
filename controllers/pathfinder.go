package controllers

import (
	"math"
	"sort"

	"lemin/models"
)

func PathFinder(colony models.Colony) (paths1 [][]string, paths2 [][]string, error error) {
	var paths [][]string
	
	// recursve dfs function to find paths
	// visited to mark visited paths
	//sort the paths by length
	//two optimization functions to get optimized paths
	// return both paths to be used in the movement and distribution functions

	visited := make(map[string]bool)
	var path []string

	var dfs func(r string)
	dfs = func(r string) {
		// return if room is already visited
		if visited[r] {
			return
		}

		// mark room as visited
		visited[r] = true

		// add room to current path
		path = append(path, r)

		// add path to paths if end room is reached
		if r == colony.EndRoom.Name {
			paths = append(paths, append([]string(nil), path...))
		} else {
			// explore the neighbors(connected rooms)
			for _, neighbor := range colony.Rooms[r].Neighbours {
				dfs(neighbor)
			}
		}

		// unmark the current room as visited
		visited[r] = false
		// remove room from current path
		path = path[:len(path)-1]
	}

	// start dfs from start room
	dfs(colony.StartRoom.Name)

	// sort slice
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	optimizedPath1 := optimize(paths, colony)
	optimizedPath2 := optimize2(paths)
	return optimizedPath1, optimizedPath2, nil
}

// optimizating paths 1
//adds path to the slice of optimized paths only if the following conditions are met:
//the length of the path is less than or equal to half the number of ants, the length of the room is not equal to the length of the first room in the optimized path and  if none of the rooms in the path is in the optimized paths
func optimize(paths [][]string, Num models.Colony) [][]string {
	optimizedPaths := [][]string{}
	optimizedPaths = append(optimizedPaths, paths[0])
	for i := 1; i < len(paths); i++ {
		firstPath := optimizedPaths[0]
		firstPathRooms := firstPath[1 : len(firstPath)-1]
		currPathRooms := paths[i][1 : len(paths[i])-1]
		if float64(len(currPathRooms)) <= math.Round(float64(Num.NumberOfAnts)) && float64(len(currPathRooms)) != float64(len(firstPathRooms)) && !contains(optimizedPaths, currPathRooms) {
			optimizedPaths = optimizedPaths[1:]
			optimizedPaths = append(optimizedPaths, paths[i])
		} else {
			if !contains(optimizedPaths, currPathRooms) {
				optimizedPaths = append([][]string{paths[i]}, optimizedPaths...)
			}
		}
	}
	return optimizedPaths
}

// optimizing paths 2
//returns unique paths only
func optimize2(paths [][]string) [][]string {
	optimizedPaths := [][]string{}
	optimizedPaths = append(optimizedPaths, paths[0])
	for i := 1; i < len(paths); i++ {
		if !contains(optimizedPaths, paths[i][1:len(paths[i])-1]) {
			optimizedPaths = append(optimizedPaths, paths[i])
		}
	}
	return optimizedPaths
}

// helper function to check if rooms in slice b are found in any rooms that are already in the slice of paths
func contains(a [][]string, b []string) bool {
	for _, slice := range a {
		for item := range slice {
			for i := range b {
				if slice[item] == b[i] {
					return true
				}
			}
		}
	}
	return false
}
