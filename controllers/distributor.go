package controllers

import (
	"fmt"
	"log"

	"lemin/models"
)

// InitializeAnts initializes all ants with optimized path distribution
func InitializeAnts(c *models.Colony) *models.Colony {
	// Get optimized paths from PathFinder
	optimizedPaths, err := PathFinder(*c)
	// fmt.Println(optimizedPaths)
	// os.Exit(1)
	if err != nil {
		log.Fatalf("Failed to find optimized paths: %v", err)
	}

	// Calculate how many ants per path using optimal distribution
	antsPerPath := CalculateOptimalAntDistribution(optimizedPaths, int(c.NumberOfAnts))
	fmt.Println(antsPerPath)
	// Initialize ants
	c.Ants = make([]models.Ant, c.NumberOfAnts)
	antIndex := 0

	// Distribute ants across the optimized paths
	for pathIndex, path := range optimizedPaths {
		for i := 0; i < antsPerPath[pathIndex]; i++ {
			// pathRooms := convertPathToRooms(path, c) // Converts room names to Room objects
			c.Ants[antIndex] = *models.NewAnt(antIndex+1, &c.StartRoom)
			c.Ants[antIndex].SetPath(path)
			antIndex++
		}
	}
	fmt.Println("c.Ants", c.Ants)
	return c
}

// CalculateOptimalAntDistribution determines the optimal number of ants per path
func CalculateOptimalAntDistribution(paths [][]string, totalAnts int) []int {
	numPaths := len(paths)
	antsPerPath := make([]int, numPaths)
	pathLengths := make([]int, numPaths)

	// Calculate the length of each path
	for i, path := range paths {
		pathLengths[i] = len(path)
	}

	// Distribute ants based on path lengths and current load
	for i := 0; i < totalAnts; i++ {
		minIndex := -1
		minLoad := int(^uint(0) >> 1) // Initialize with maximum integer value

		// Find the path with the minimum load
		for j := 0; j < numPaths; j++ {
			currentLoad := antsPerPath[j] + pathLengths[j]
			if currentLoad < minLoad {
				minLoad = currentLoad
				minIndex = j
			}
		}

		// Assign an ant to the path with the minimum load
		if minIndex != -1 {
			antsPerPath[minIndex]++
		}
	}

	return antsPerPath
}

// // convertPathToRooms converts a path of room names to Room objects
// func convertPathToRooms(path []string, c *models.Colony) []*models.Room {
// 	rooms := make([]*models.Room, len(path))
// 	for i, roomName := range path {
// 		room := c.GetRoomByName(roomName)
// 		if room == nil {
// 			log.Fatalf("Room %s not found in colony", roomName)
// 		}
// 		rooms[i] = room
// 	}
// 	return rooms
// }
