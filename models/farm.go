package models

type Farm struct {
	AntCount  int
	Rooms     map[string]*Room
	Tunnels   []*Tunnel
	StartRoom *Room
	EndRoom   *Room
}