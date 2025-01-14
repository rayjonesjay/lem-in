package models

type Ant struct {
	ID       int
	Position *Room
	Path     []*Room
	PathIdx  int
}

func NewAnt(id int, startRoom *Room) *Ant {
	return &Ant{
		ID:       id,
		Position: startRoom,
		PathIdx:  0,
	}
}
