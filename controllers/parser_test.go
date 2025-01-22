package controllers

import (
	"os"
	"strings"
	"testing"
)

// Tests for the parser function
// create a function to create files for use and delete them after tests are completed
func createTestFile(content string) (string, error) {
	tempFile, err := os.CreateTemp("", "parser_test_*.txt")
	if err != nil {
		return "", err
	}

	defer tempFile.Close()

	_, err = tempFile.WriteString(content)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

// Valid File: Tests a correct file with ants, start room, end room, and connections.
func TestParser_ParseFile_ValidFile(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	colony, err := parser.ParseFile(fileName)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}

	// check if number of ants is found
	if colony.NumberOfAnts != 10 {
		t.Errorf("Expected ant count of 10, got %d", colony.NumberOfAnts)
	}

	// check if start room exists or is the correct name
	if colony.StartRoom.Name != "1" || !colony.StartRoom.IsStart {
		t.Errorf("Expected start room 1, got %s", colony.StartRoom.Name)
	}

	// check if end room exists and its the right name
	if colony.EndRoom.Name != "0" || !colony.EndRoom.IsEnd {
		t.Errorf("Expected end  room 0, got %s", colony.EndRoom.Name)
	}

	// checkfor the correct number of rooms
	if len(colony.Rooms) != 8 {
		t.Errorf("Expected 8 rooms, got %d", len(colony.Rooms))
	}

	// check for the number of neighbours in  a room
	if len(colony.Rooms["0"].Neighbours) != 2 {
		t.Errorf("Expected 2 neighbours for room '0', got %d", len(colony.Rooms["0"].Neighbours))
	}
}

// Missing Start Room: Ensures that a missing start room is flagged.
func TestParser_ParseFile_MissingStartRoom(t *testing.T) {
	content := `10
##end
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
`
	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()
	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "missing start or end room") {
		t.Fatalf("Expected error 'missing start or end room', got %v", err)
	}
}

// Missing End Room: Ensures that a missing end room is flagged.
func TestParser_ParseFile_MissingEndRoom(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "missing start or end room") {
		t.Fatalf("Expected error 'missing start or end room', got %v", err)
	}
}

// No Connections: Verifies that rooms without connections are handled correctly.
func TestParser_ParseFile_NoConnections(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	colony, err := parser.ParseFile(fileName)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}

	if len(colony.Rooms) != 8 {
		t.Errorf("Expected 8 rooms, got %d", len(colony.Rooms))
	}

	// Ensure no rooms have connections
	for _, room := range colony.Rooms {
		if len(room.Neighbours) > 0 {
			t.Errorf("Expected no neighbours for room '%s', got %d", room.Name, len(room.Neighbours))
		}
	}
}

// No Rooms: Ensures that a file with no rooms results in an error.
func TestParser_ParseFile_NoRooms(t *testing.T) {
	content := `10
##start
##end
0-5
1-5
2-3
2-4
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "room 0 not found") {
		t.Fatalf("Expected error 'room 0 not found', got %v", err)
	}
}

// No Ants: Ensures a file with no ants is flagged as invalitiond.
func TestParser_ParseFile_NoAnts(t *testing.T) {
	content := `
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "ERROR: invalid data format\nnumber of ants invalid") {
		t.Fatalf("Expected 'ERROR: invalid data format\nnumber of ants invalid', got %v", err)
	}
}

//No Ants: Ensures a file with an invalid number of is flagged as invalitiond.
func TestParser_ParseFile_InvalidNumberofAnts(t *testing.T) {
	content := `
	-12
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "ERROR: invalid data format\nnumber of ants invalid") {
		t.Fatalf("Expected 'ERROR: invalid data format\nnumber of ants invalid', got %v", err)
	}
}


// Room name startting with L
func TestParser_ParseFile_RoomNameStartsWithL(t *testing.T) {
	content := `10
##start
L1 23 3
##end
0 9 5
0-L1
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "ERROR: room starts with # or L ") {
		t.Fatalf("Expected error 'ERROR: room starts with # or L ', got %v", err)
	}
}

// Room name startting with #
// Room name startting with L
func TestParser_ParseFile_RoomNameStartsWithH(t *testing.T) {
	content := `10
##start
#room 23 3
##end
0 9 5
0-#room
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "#room not found") {
		t.Fatalf("Expected error '#room not found', got %v", err)
	}
}

// Room name  with spaces
func TestParser_ParseFile_RoomNameWithSpaces(t *testing.T) {
	content := `10
##start
Room 1 23 3
##end
0 9 5
0-Room 1
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "invalid room format") {
		t.Fatalf("Expected error 'invalid room format', got %v", err)
	}
}

// duplicate tunnel between two rooms
func TestParser_ParseFile_DuplicateTunnels(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
##end
3 9 5
1-2
1-2
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "duplicate tunnel") {
		t.Fatalf("Expected error 'duplicate tunnel', got %v", err)
	}
}

// tunnel connecting more than two rooms
func TestParser_ParseFile_InvalidTunnelFormat(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
3 12 8
4 11 9
##end
3 9 5
1-2-3

`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "invalid connection format: 1-2-3") {
		t.Fatalf("Expected error 'invalid connection format: 1-2-3', got %v", err)
	}
}

// Room connected to itself
func TestParser_ParseFile_SelfLoop(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
##end
3 9 5
1-1
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "invalid connection: 1-1 (self loop)") {
		t.Fatalf("Expected error 'invalid connection: 1-1 (self loop)', got %v", err)
	}
}

// Rooms with the same name
func TestParser_ParseFile_DuplicateRoomNames(t *testing.T) {
	content := `10
##start
1 23 3
2 16 7
1 10 10
##end
3 9 5
1-2
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err != nil {
		t.Fatalf("Expected error 'no error', got %v", err)
	}
}

// Invalid room coordinats
func TestParser_ParseFile_InvalidRoomCoordinates(t *testing.T) {
	content := `10
##start
1 s 3
2 16 7
3 10 10
##end
3 9 5
1-2
`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	parser := NewParser()

	_, err = parser.ParseFile(fileName)
	if err == nil || !strings.Contains(err.Error(), "invalid x coordinate: s") {
		t.Fatalf("Expected error 'invalid x coordinate: s', got %v", err)
	}
}

