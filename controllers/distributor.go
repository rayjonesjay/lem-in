package controllers

import (
	"log"

	"lemin/models"
)

// InitializeAnts initializes all ants with optimized path distribution
func InitializeAnts(c *models.Colony) (*models.Colony, int) {
	// Get optimized paths from PathFinder
	var bestPaths = [][]string{}
	var bestAntsPerPath = []int{}
	var leastTurn = 0
	optimizedPaths1, optimizedPaths2, err := PathFinder(*c)

	if err != nil {
		log.Fatalf("Failed to find optimized paths: %v", err)
	}

	// Calculate how many ants per path using optimal distribution
	antsPerPath1 := CalculateOptimalAntDistribution(optimizedPaths1, int(c.NumberOfAnts))
	antsPerPath2 := CalculateOptimalAntDistribution(optimizedPaths2, int(c.NumberOfAnts))

	turns1 := getTotalTurns(optimizedPaths1, antsPerPath1)
	turns2 := getTotalTurns(optimizedPaths2, antsPerPath2)

	if turns1 < turns2 {
		leastTurn = turns1
		bestPaths = optimizedPaths1
		bestAntsPerPath = antsPerPath1
	} else {
		leastTurn = turns2
		bestPaths = optimizedPaths2
		bestAntsPerPath = antsPerPath2
	}
	// Initialize ants
	c.Ants = make([]models.Ant, c.NumberOfAnts)
	antIndex := 0

	// Distribute ants across the optimized paths
	for pathIndex, path := range bestPaths {
		for i := 0; i < bestAntsPerPath[pathIndex]; i++ {
			// pathRooms := convertPathToRooms(path, c) // Converts room names to Room objects
			c.Ants[antIndex] = *models.NewAnt(antIndex+1, &c.StartRoom)
			c.Ants[antIndex].SetPath(path)
			antIndex++
		}
	}

	return c, leastTurn
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

// getTotalTurns calculates the total number of turns required for the given distribution
func getTotalTurns(paths [][]string, antsPerPath []int) int {
	if len(paths) == 0 || len(antsPerPath) == 0 {
		return 0
	}
	maxTurns := 0
	for i, path := range paths {
		if antsPerPath[i] > 0 {
			turns := len(path) - 1 + antsPerPath[i] - 1
			if turns > maxTurns {
				maxTurns = turns
			}
		}
	}
	return maxTurns
}
