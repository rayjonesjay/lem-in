package controllers

import (
	"fmt"
	"lemin/models"
)

func Mover(c *models.Colony) {
	// Track positions of ants in a map for each turn
	for turn := 1; turn <= len(c.Ants[0].Path); turn++ {
		fmt.Printf("Turn %d: ", turn)

		// Track which rooms are occupied (map of room -> ant ID)
		occupied := make(map[string]int)

		// Store the movement results
		movementResults := make(map[int]string)

		// First pass: check which ants can move and ensure no room is occupied
		for _, ant := range c.Ants {
			// Get the next room the ant should move to
			nextRoom := ant.GetNextRoom()
			if nextRoom == "" {
				continue // Skip if no next room (end of path)
			}

			// Check if the room is already occupied by another ant
			if _, exists := occupied[nextRoom]; exists {
				// If occupied, the ant doesn't move (it stays in the same room)
				movementResults[ant.ID] = ant.Position
			} else {
				// Mark the room as occupied
				occupied[nextRoom] = ant.ID
				// Move the ant and record the new position
				ant.Move()
				movementResults[ant.ID] = ant.Position
			}
		}

		// Print the results for this turn in the required format
		for _, ant := range c.Ants {
			if movementResults[ant.ID] != "" {
				// Print the ant ID and its position
				fmt.Printf("L%d-%s ", ant.ID, movementResults[ant.ID])
			}
		}
		fmt.Println() // Newline after each turn
	}
}
