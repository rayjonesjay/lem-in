package models

type Ant struct {
	ID       int
	Position string
	Path     []string
	PathIdx  int
}

// NewAnt creates a new ant with the given ID and starting room
func NewAnt(id int, startRoom *Room) *Ant {
	return &Ant{
		ID:       id,
		Position: "",
		Path:     []string{},
		PathIdx:  0,
	}
}

// SetPath sets the path
func (a *Ant) SetPath(path []string) {
	a.Path = path
	a.PathIdx = 0
}

// GetNextRoom returns the next room in the path without moving the ant
func (a *Ant) GetNextRoom() string {
	if a.PathIdx >= len(a.Path)-1 {
		return ""
	}
	return a.Path[a.PathIdx+1]
}

// Move ant to the next room in the path and return the new room
func (a *Ant) Move() string {
	if a.PathIdx >= len(a.Path)-1 {
		return ""
	}

	a.PathIdx++
	a.Position = a.Path[a.PathIdx]
	return a.Position
}
