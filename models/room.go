package models

// Room represents a single room in the "colony"
// a room consist:
// 1. name
// 2. coordinates
// 3. neighbours
type Room struct {
	Name       string
	Coordinate Coordinate
	IsStart    bool
	IsEnd      bool
	Neighbours []*Room // a slice of rooms/neighbours
	IsStart bool
	IsEnd bool
}

// Cartesian coordinates, this type represent the X and Y coordinates of a room
type Coordinate struct {
	X float64
	Y float64
}
