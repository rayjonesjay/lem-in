package types

import "sync"

// Ant ANT RULES
/*
1. each room can only contain one ant at a time ( except at ##start and ##end which contain as many ants as needed)
2. to be the first ant to arrive, ants will need to take the shortest paths.
	also manage traffic and also prevent them waling over their fellow ants
3. only display ants that moved on each turn, you can only move each ant only once and through a tunnel.
	the room where the ant is going should be empty. - check in advance
4.
*/
type Ant struct {
	Id          uint64
	mu          sync.Mutex
	Current     string
	Destination string
	Path        []string
}
