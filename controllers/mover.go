package controllers

import (
	"fmt"

	"lemin/models"
)

func Mover(c *models.Colony) {
	colony, maxTurns := InitializeAnts(c)

	// Track which rooms are occupied by which ants
	occupied := make(map[string]*models.Ant)
	numberOfAntsAtEnd := 0

	// Initialize all ants at start position
	for i := range colony.Ants {
		colony.Ants[i].Position = colony.StartRoom.Name
	}

	var allMoves [][]string
	turnCount := 0

	// Continue until all ants reach the end
	for numberOfAntsAtEnd < len(colony.Ants) && turnCount < maxTurns {
		currentTurnMoves := make([]string, 0)
		movedThisTurn := make(map[int]bool)

		// Keep trying moves until no more moves are possible this turn
		changed := true
		for changed {
			changed = false

			// Try to move each ant
			for i := range colony.Ants {
				if movedThisTurn[colony.Ants[i].ID] {
					continue
				}

				ant := &colony.Ants[i]

				// Skip if ant is already at end
				if ant.Position == colony.EndRoom.Name {
					continue
				}

				// Find next room in ant's path
				var nextRoom string
				for j, room := range ant.Path {
					if room == ant.Position && j+1 < len(ant.Path) {
						nextRoom = ant.Path[j+1]
						break
					}
				}

				// Move ant if next room is available
				if nextRoom != "" && occupied[nextRoom] == nil {
					// Clear current room from occupied map
					if ant.Position != colony.StartRoom.Name {
						delete(occupied, ant.Position)
					}

					// Move ant
					ant.Position = nextRoom

					// Update occupied rooms and ant count
					if nextRoom != colony.EndRoom.Name {
						occupied[nextRoom] = ant
					} else {
						numberOfAntsAtEnd++
					}

					currentTurnMoves = append(currentTurnMoves, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
					movedThisTurn[ant.ID] = true
					changed = true
				}
			}
		}

		// Add turn if any moves were made
		if len(currentTurnMoves) > 0 {
			allMoves = append(allMoves, currentTurnMoves)
		}
		turnCount++
	}

	for _, input := range c.Output {
		fmt.Println(input)
	}
	println()

	// Print all moves
	for _, turnMoves := range allMoves {
		for i, move := range turnMoves {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(move)
		}
		fmt.Println()
	}
}
