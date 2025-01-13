package models

// Room represents a single room in the "colony"
// a room consist:
// 1. name
// 2. coordinates
// 3. neighbours
type Room struct {
	Name string
	Coordinate Coordinate
	Neighbours map[string]*Room
}

// Cartesian coordinates, this type represent the X and Y coordinates of a room
type Coordinate struct {
	X float64
	Y float64
}