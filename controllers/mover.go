package controllers

import (
	"lemin/models"
)

// Movement represents a single ant movement for output
type Movement struct {
	AntID  int
	ToRoom string
}

// Mover coordinates simultaneous ant movements
type Mover struct {
	ants   []*models.Ant
	colony *models.Colony // Reference to the Colony for start and end room
}

// NewMover creates a new Mover instance
func NewMover(ants []*models.Ant, colony *models.Colony) *Mover {
	return &Mover{
		ants:   ants,
		colony: colony,
	}
}

// ExecuteMovements returns a series of steps where each step contains
// simultaneous valid movements for that turn
func (m *Mover) ExecuteMovements() [][]Movement {
	var allSteps [][]Movement

	// Continue until all ants reach the end
	for {
		currentMoves := m.executeNextStep()
		if len(currentMoves) == 0 {
			break
		}
		allSteps = append(allSteps, currentMoves)
	}

	return allSteps
}

// executeNextStep performs one turn of simultaneous ant movements
func (m *Mover) executeNextStep() []Movement {
	var moves []Movement
	occupied := make(map[string]bool)
	finished := make(map[int]bool) // Track ants that are done

	// Prepopulate occupied map with current positions (excluding start and end rooms)
	for _, ant := range m.ants {
		if ant.Position != nil &&
			ant.Position.Name != m.colony.StartRoom.Name &&
			ant.Position.Name != m.colony.EndRoom.Name {
			occupied[ant.Position.Name] = true
		}
	}

	// Try to move each ant that hasn't finished
	for _, ant := range m.ants {
		if finished[ant.ID] {
			continue
		}

		nextRoom := ant.GetNextRoom()
		if nextRoom == nil {
			// Mark the ant as finished if there's no next room
			finished[ant.ID] = true
			continue
		}

		// Skip occupancy check for start and end rooms
		if nextRoom.Name == m.colony.StartRoom.Name ||
			nextRoom.Name == m.colony.EndRoom.Name ||
			!occupied[nextRoom.Name] {
			moves = append(moves, Movement{
				AntID:  ant.ID,
				ToRoom: nextRoom.Name,
			})

			// Update occupancy for the current room (if not start or end room)
			if ant.Position != nil &&
				ant.Position.Name != m.colony.StartRoom.Name &&
				ant.Position.Name != m.colony.EndRoom.Name {
				occupied[ant.Position.Name] = false
			}

			// Mark the next room as occupied (if not start or end room)
			if nextRoom.Name != m.colony.StartRoom.Name &&
				nextRoom.Name != m.colony.EndRoom.Name {
				occupied[nextRoom.Name] = true
			}

			// Move the ant
			ant.Move()
		}
	}

	return moves
}
