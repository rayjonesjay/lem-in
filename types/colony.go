package types

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"lem-in/xerrors"
)

const (
	MaxAntsPerColony = 1_000
)

// Room a looks like A 1 2 where A is the name, and 1 and 2 is the x and y coordinate respectively
// although this coordinates will be used in visualization
type Room struct {
	Name       string
	X, Y       int      // room coordinates
	Neighbours []string // each room has neighbours
	Occupied   bool     // indicator to show room occupation status
	mu         sync.Mutex
}

// Colony is a model of the ant farm also known as colony
type Colony struct {
	StartRoom    Room             // the initial position where all ants are supposed to start from
	StartFound   bool             // true if start room was found in the input file
	EndFound     bool             // true if end room was found in the input file
	EndRoom      Room             // the destination where all ants are supposed to go
	NumberOfAnts uint64           // number of ants cannot be negative
	Rooms        map[string]*Room // rooms is a slice of rooms
	Ants         []Ant            // all ants in the colony
}

func (c *Colony) FindAllPaths(start, end string) [][]string {
	// Check if the start or end room doesn't exist
	if c.Rooms[start] == nil || c.Rooms[end] == nil {
		return nil
	}

	// Initialize a slice to store all paths
	var allPaths [][]string

	// Helper function to perform DFS
	var dfs func(path []string, current string)
	dfs = func(path []string, current string) {
		// Add the current room to the path
		path = append(path, current)

		// If the current room is the destination, save the path
		if current == end {
			// Make a copy of the path and add it to allPaths
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
			return
		}

		// Explore each neighbor
		for _, neighbor := range c.Rooms[current].Neighbours {
			// Avoid revisiting rooms already in the current path
			if !contains(path, neighbor) {
				dfs(path, neighbor)
			}
		}
	}

	// Start the DFS from the starting room
	dfs([]string{}, start)

	return allPaths
}

// Helper function to check if a room is already in the path
func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func (c *Colony) MoveAnts() error {
	// the outer loop represents the undefined number of turns the ants will take
	// inside the outer loop another loop which checks if an ant can move to the next room
	// depending on its current position and its path
	// an ant can move if the next room in its path is empty
	// lock the current room and destination to prevent race
	fmt.Println("******************")
	xerrors.Logger(c, "json.txt")
	for {
		numberOfAntsWhoReachedTheEnd := 0
		for i := range c.Ants {

			ant := &c.Ants[i]

			if ant.Current == c.EndRoom.Name {
				numberOfAntsWhoReachedTheEnd += 1
				continue
			}

			if len(ant.Path) < 2 {
				return fmt.Errorf("invalid path for %d", ant.Id)
			}

			xerrors.Logger(c, "ant.txt")
			// if the current ant position is not equal to the end room
			if ant.Current != c.EndRoom.Name {

				if ant.Path == nil || len(ant.Path) == 0 {
					return fmt.Errorf(errors.New("ant %d does not have a defined path to move").Error(), ant.Id)
				}

				nextRoom := ant.Path[1] // get next room in the path
				currentRoom := c.Rooms[ant.Current]

				currentRoom.mu.Lock()
				nextRoomObject := c.Rooms[nextRoom]
				nextRoomObject.mu.Lock()

				if !nextRoomObject.Occupied {
					// move the ant
					ant.Current = nextRoom
					nextRoomObject.Occupied = true
					currentRoom.Occupied = false
					ant.Path = ant.Path[1:]
					fmt.Printf("L%d-%s ", ant.Id, nextRoom)
				}
				currentRoom.mu.Unlock()
				nextRoomObject.mu.Unlock()
			}
			if numberOfAntsWhoReachedTheEnd == len(c.Ants) {
				break
			}
		}
	}
	return nil
}

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

	// Helper function to parse room name and coordinates
	parseRoom := func(s string) (name string, x, y int, err error) {
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
			// split at the occurrence of -
			arr := Split(line, "-")
			//fmt.Println(arr)
			if len(arr) != 2 {
				return nil, fmt.Errorf(xerrors.ErrInvalidLink.Error(), line)
			}

			from, to := arr[0], arr[1]
			//fmt.Println(from, to)
			// create a room
			room := colony.Rooms[from]
			// add to as its neighbour
			room.Neighbours = append(room.Neighbours, to)

			room = colony.Rooms[to]
			// for undirected graphs
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

		// Parse regular rooms
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

	// Validate if start and end rooms were found
	if !colony.StartFound || colony.StartRoom.Name == "" {
		return nil, xerrors.ErrStartNotFound
	}
	if !colony.EndFound || colony.EndRoom.Name == "" {
		return nil, xerrors.ErrEndNotFound
	}

	C := InitializeAnts(colony)
	return C, nil
}

// CheckNumAnts checks the number of ants per colony if they have exceeded MaxAntsPerColony
// and exits with status code 1
func CheckNumAnts(c *Colony) error {
	if c.NumberOfAnts >= MaxAntsPerColony {
		return xerrors.ErrMaxAntNumExceeded
	}
	if c.NumberOfAnts == 0 {
		return xerrors.ErrZeroAnts
	}
	return nil
}

// InitializeAnts initializes all ants, with their id,start room name,paths
func InitializeAnts(c *Colony) *Colony {
	// all ant id start at 1
	for i := uint64(1); i <= c.NumberOfAnts; i++ {
		ant := Ant{
			Id:          i,
			Current:     c.StartRoom.Name,
			Destination: c.EndRoom.Name,
			Path:        nil, // this will be calculated
		}
		ant.Path = c.PathFinder(c.StartRoom.Name, c.EndRoom.Name)
		xerrors.Logger(ant, "ant.err")
		c.Ants = append(c.Ants, ant)
	}
	return c
}

// isLink looks for pattern X in s where pattern X is M-N where M and N are both numbers or letters
// and returns true if a match exist
func isLink(s string) bool {
	pattern := `\d+-\d+`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

func Split(s string, sep string) []string {
	arr := strings.Split(s, sep)
	return arr
}

// ValidateColony checks if the fields in the colony struct have the correct fields
// and that the fields obey the rules of the game
func ValidateColony(c *Colony) error {
	// check if number of ants exceed 1000
	err := CheckNumAnts(c)
	if err != nil {
		return err
	}
	return nil
}
