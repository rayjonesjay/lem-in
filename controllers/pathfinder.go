package controllers

import (
	"fmt"
	"math"
	"sort"

	"lemin/models"
)

func PathFinder(colony models.Colony) ([][]string, error) {
	var paths [][]string
	// check if startroom or end room is missing

	// check if both start room and end room are present
	// dfs function to find paths
	// visted to mark visited paths
	// sort the final paths to return a slice from the shortest to the longest(optimization)
	visited := make(map[string]bool)
	var path []string

	var dfs func(r string)
	dfs = func(r string) {
		// return if room is already visited
		if visited[r] {
			return
		}

		// fmt.Println(r)
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
	fmt.Println("one", optimizedPath1)
	fmt.Println("two", optimizedPath2)
	if len(optimizedPath1) > len(optimizedPath2) {
		return optimizedPath1, nil
	} else {
		return optimizedPath2, nil
	}
}

// optimizating paths 1
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

// helper function
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
