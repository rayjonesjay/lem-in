package Colony

type Room struct {
	RoomNo string
	X, Y   int
}

type ColonyGraph struct {
	Connections Connections
	Start       string
	End         string
}

type Connections map[Room][]Room

func (con *ColonyGraph) AddConnection(room1, room2 Room) {
	con.Connections[room1] = append(con.Connections[room1], room2)
	con.Connections[room2] = append(con.Connections[room2], room1)
}
