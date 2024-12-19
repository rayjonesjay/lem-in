package types

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"lem-in/xerrors"
)

const (
	MaxAntsPerColony = 1_000
)

// Room represents a location in the colony with coordinates and neighbors
type Room struct {
	Name       string
	X, Y       int      // room coordinates
	Neighbours []string // each room has neighbours
	Occupied   bool     // indicator to show room occupation status
	// mu         sync.Mutex
}

// Colony represents the ant farm structure and its components
type Colony struct {
	StartRoom    Room             // the initial position where all ants are supposed to start from
	StartFound   bool             // true if start room was found in the input file
	EndFound     bool             // true if end room was found in the input file
	EndRoom      Room             // the destination where all ants are supposed to go
	NumberOfAnts uint64           // number of ants cannot be negative
	Rooms        map[string]*Room // rooms is a slice of rooms
	Ants         []Ant            // all ants in the colony
}

// FindAllPaths finds all possible paths from start to end room
func (c *Colony) FindAllPaths(start, end string) [][]string {
	if c.Rooms[start] == nil || c.Rooms[end] == nil {
		return nil
	}

	var allPaths [][]string

	var dfs func(path []string, current string)
	dfs = func(path []string, current string) {
		path = append(path, current)

		if current == end {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
			return
		}

		for _, neighbor := range c.Rooms[current].Neighbours {
			if !contains(path, neighbor) {
				dfs(path, neighbor)
			}
		}
	}

	dfs([]string{}, start)
	return allPaths
}

// Helper function to check if a room is in a path
func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

// MoveAnts coordinates ant movement through the colony
func (c *Colony) MoveAnts() error {
	numberOfAntsAtEnd := 0
	occupied := make(map[string]bool)

	// Mark start room as occupied initially
	occupied[c.StartRoom.Name] = true

	for numberOfAntsAtEnd < len(c.Ants) {
		moves := make([]string, 0)
		newOccupied := make(map[string]bool)

		// Process ants in order
		for i := range c.Ants {
			ant := &c.Ants[i]

			if ant.Current == c.EndRoom.Name {
				continue
			}

			if len(ant.Path) < 2 {
				continue
			}

			nextRoom := ant.Path[1]

			// Check if next room is available
			if !occupied[nextRoom] && !newOccupied[nextRoom] {
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.Id, nextRoom))
				newOccupied[nextRoom] = true

				// Update ant's position
				ant.Current = nextRoom
				ant.Path = ant.Path[1:]

				if ant.Current == c.EndRoom.Name {
					numberOfAntsAtEnd++
				}
			}
		}

		// Update occupied rooms for next turn
		occupied = newOccupied

		// Print moves for this turn
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		} else if numberOfAntsAtEnd < len(c.Ants) {
			return fmt.Errorf("deadlock detected: no moves possible but not all ants at end")
		}
	}

	return nil
}

// ParseFileContentsToColony parses input file contents into a Colony structure
func ParseFileContentsToColony(fileContents []string) (*Colony, error) {
	colony := &Colony{
		Rooms: make(map[string]*Room),
	}

	// First line: Number of ants
	numberOfAnts, err := strconv.Atoi(strings.TrimSpace(fileContents[0]))
	if err != nil || numberOfAnts < 0 {
		return nil, xerrors.ErrInvalidNumberOfAnts
	}
	colony.NumberOfAnts = uint64(numberOfAnts)

	// Check for ##start and ##end directives
	i := 1
	if strings.TrimSpace(fileContents[i]) == "##start" {
		colony.StartFound = true
		i++
		name, x, y, err := parseRoom(fileContents[i])
		if err != nil {
			return nil, err
		}
		colony.StartRoom = Room{Name: name, X: x, Y: y}
		colony.Rooms[name] = &colony.StartRoom
		i++
	} else {
		return nil, xerrors.ErrStartNotFound
	}

	endRoomSet := false
	for ; i < len(fileContents); i++ {
		line := strings.TrimSpace(fileContents[i])

		if isLink(line) {
			arr := Split(line, "-")
			if len(arr) != 2 {
				return nil, fmt.Errorf(xerrors.ErrInvalidLink.Error(), line)
			}

			from, to := arr[0], arr[1]
			room := colony.Rooms[from]
			room.Neighbours = append(room.Neighbours, to)

			room = colony.Rooms[to]
			room.Neighbours = append(room.Neighbours, from)
			continue
		}

		if line == "##end" && !endRoomSet {
			colony.EndFound = true
			i++
			name, x, y, err := parseRoom(fileContents[i])
			if err != nil {
				return nil, err
			}
			colony.EndRoom = Room{Name: name, X: x, Y: y}
			colony.Rooms[name] = &colony.EndRoom
			endRoomSet = true
			continue
		}

		if !strings.HasPrefix(line, "#") && line != "" {
			name, x, y, err := parseRoom(line)
			if err != nil {
				return nil, err
			}
			if _, exists := colony.Rooms[name]; exists {
				return nil, fmt.Errorf(xerrors.ErrDuplicateRoom.Error(), colony.Rooms[name])
			}
			colony.Rooms[name] = &Room{Name: name, X: x, Y: y}
		}
	}

	if !colony.StartFound || colony.StartRoom.Name == "" {
		return nil, xerrors.ErrStartNotFound
	}
	if !colony.EndFound || colony.EndRoom.Name == "" {
		return nil, xerrors.ErrEndNotFound
	}

	return InitializeAnts(colony), nil
}

// Helper function to parse room information
func parseRoom(s string) (name string, x, y int, err error) {
	parts := strings.Fields(strings.TrimSpace(s))
	if len(parts) != 3 {
		return "", 0, 0, xerrors.ErrInvalidRoomFormat
	}

	name = parts[0]
	x, err = strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, 0, xerrors.ErrInvalidRoomCoordinates
	}
	y, err = strconv.Atoi(parts[2])
	if err != nil {
		return "", 0, 0, xerrors.ErrInvalidRoomCoordinates
	}
	return name, x, y, nil
}

// InitializeAnts initializes all ants with optimized path distribution
func InitializeAnts(c *Colony) *Colony {
	allPaths := c.FindAllPaths(c.StartRoom.Name, c.EndRoom.Name)

	// Sort paths by length
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Get top N shortest paths where N is number of ants
	maxPaths := int(c.NumberOfAnts)
	if len(allPaths) > maxPaths {
		allPaths = allPaths[:maxPaths]
	}

	// Initialize ants with distributed paths
	c.Ants = make([]Ant, c.NumberOfAnts)

	// Assign paths round-robin style to ensure even distribution
	for i := uint64(0); i < c.NumberOfAnts; i++ {
		pathIndex := int(i) % len(allPaths)
		path := make([]string, len(allPaths[pathIndex]))
		copy(path, allPaths[pathIndex])

		c.Ants[i] = Ant{
			Id:          i + 1,
			Current:     c.StartRoom.Name,
			Destination: c.EndRoom.Name,
			Path:        path,
		}
	}

	return c
}

// CheckNumAnts validates the number of ants in the colony
func CheckNumAnts(c *Colony) error {
	if c.NumberOfAnts >= MaxAntsPerColony {
		return xerrors.ErrMaxAntNumExceeded
	}
	if c.NumberOfAnts == 0 {
		return xerrors.ErrZeroAnts
	}
	return nil
}

// ValidateColony validates the colony structure
func ValidateColony(c *Colony) error {
	return CheckNumAnts(c)
}

// isLink checks if a string represents a valid link between rooms
func isLink(s string) bool {
	pattern := `[0-9A-Za-z]+-[0-9A-Za-z]+`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

// Split splits a string by a separator
func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// PrintState prints the current state of all ants (for debugging)
func (c *Colony) PrintState() {
	fmt.Println("\nCurrent State:")
	for _, ant := range c.Ants {
		fmt.Printf("Ant %d: Current=%s, Path=%v\n",
			ant.Id, ant.Current, ant.Path)
	}
	fmt.Println()
}
