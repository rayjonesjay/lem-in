package models

type Ant struct {
	ID       int
	Position *Room
	Path     []*Room
	PathIdx  int
}

// NewAnt creates a new ant with the given ID and starting room
func NewAnt(id int, startRoom *Room) *Ant {
	return &Ant{
		ID:       id,
		Position: startRoom,
		PathIdx:  0,
	}
}

// SetPath move ant to the next room in the path
func (a *Ant) SetPath(path []*Room) {
	a.Path = path
	a.PathIdx = 0
}

// GetNextRoom returns the next room in the path without moving the ant
func (a *Ant) GetNextRoom() *Room {
	if a.PathIdx >= len(a.Path)-1 {
		return nil
	}
	return a.Path[a.PathIdx+1]
}

// Move ant to the next room in the path and return the new room
func (a *Ant) Move() *Room {
	if a.PathIdx >= len(a.Path)-1 {
		return nil
	}

	a.PathIdx++
	a.Position = a.Path[a.PathIdx]
	return a.Position
}
