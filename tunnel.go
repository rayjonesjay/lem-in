package models

// Tunnel represents a tunnel between two rooms
type Tunnel struct {
	Room1 *Room
	Room2 *Room
}

// NewTunnel creates a new tunnel connecting two rooms
func NewTunnel(room1, room2 *Room) *Tunnel {
	tunnel := &Tunnel{
		Room1: room1,
		Room2: room2,
	}

	tunnel.ConnectRooms()
	return tunnel
}

// ConnectRooms connects the two rooms with the tunnel
func (t *Tunnel) ConnectRooms() {
	t.Room1.Neighbours = append(t.Room1.Neighbours, t.Room2)
	t.Room2.Neighbours = append(t.Room2.Neighbours, t.Room1)
}
