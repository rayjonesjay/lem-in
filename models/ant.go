package models

type Ant struct {
	ID       int
	Position *Room
	Path     []*Room
	PathIdx  int
}

// Create a new ant with the given ID and starting room
func NewAnt(id int, startRoom *Room) *Ant {
	return &Ant{
		ID:       id,
		Position: startRoom,
		PathIdx:  0,
	}
}

// Move ant to the next room in the path
func (a *Ant) SetPath(path []*Room) {
	a.Path = path
	a.PathIdx = 0
}
