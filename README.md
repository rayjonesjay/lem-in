# lem-in

## About
Lem-in, a digital ant farm simulation, challenges you to implement algorithms for graph traversal
and optimization in Go.
Lem-in combines file parsing, graph theory, and simulation.


## Analogy: Navigating a Maze
An ant colony is a group of ants that live together in a structured community.
Imagine an ant colony is a maze. The ants are numbered 1 to N, they start at the entrance `##start`
and need to find the shortest way to exit `##end`.

The maze/colony has rooms(nodes) connected by tunnels (edges).
Ants can only move to adjacent rooms, but they cannot occupy the same room simultaneously - this avoids traffic jams.

The goal is to guide the ants to move from start to end and display their progress.

## How the program works
1. Takes in a file as an argument.
2. Parses the file into a colony
3. Finds optimum paths and distributes the ants to the paths accordingly.
4. Moves ant through the colony from the start room to the end room.


## Explanation

1. FILE PARSING
Here are the contents and structure of the file we are reading from:
- **Input structure**
	- the number of ants - the fist line of each file contains the number of ants
	- room definition - a room has this format `room_name x-cordinate y-cordinate` eg `A 1 2`.
	- tunnel definition - `room1-room2` eg `A-B`
	- commands `##start` and `#end` mark the start and end rooms.
- **Validation**:
  - Ensure:
	- number of ants is valid (1 to N).
	- there is one `##start` and one `##end`.
	- no duplicate rooms or invalid tunnels.
	- rooms and links conform to specified formats.
- **Errors**:
 - For invalid input return specific messages:
	- example `ERROR: invalid data format, no start room found`.
	

2. GRAPH CONSTRUCTION
- COLONY==GRAPH:
	- nodes: rooms (`A`,`B`,`C`)
	- edges: tunnels (`A-C`, `C-B`)
- USE ADJACENCY LIST FOR EFFICIENT TRAVERSAL:
```go
graph := map[string][]string{
    "A": {"C"},
    "C": {"A", "B"},
    "B": {"C"},
}
```

3. FINDING PATHS
- ALGORITHM:
	- use **depth-first search** DFS to find the paths from `##start` to `##end`
	- consider addional paths for parallel traversal to minimize moves

- CHALLENGES:
	- avoid loops and self-links
	- handle multiple shortest paths efficiently

4. SIMULATION
- Ant Movement Rules:
	- each ant can move once per turn
	- a room (except `##start` and `##end`) can only hold one ant at a time.
	- ants follow the shortest paths but avoid congestion


## Installation

### Step 1: Install Go
1. Download and install Go from the [official Go website](https://go.dev/dl/). Ensure that the Go version is `1.21` and above
2. Follow the instructions for your operating system.

### Step 2: Clone the Repository
Run the following command in your terminal to clone the repository:
```bash
git clone https://github.com/rayjonesjay/lem-in/
```

### Step 3: Navigate to the Project Directory
```bash
cd lem-in
```

### Step 4: Run the Project
Run the project using the `go run` command:
```bash
go run . filename.txt
```

### Example usage:
Consider a sample file `test.txt` with this structure:
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```
- To run the program, run the following command
```bash
go run . test.txt
```
- The following output will be printed on the terminal:
```bash
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```
- ants are labeled as L1, L2 and L3 and their movement per turn is displayed from start room to end room.

---

## Contributors
- [**Ramuiruri**:]('https://github.com/rayjonesjay')
- [**wonyango**: ]('https://github.com/WycliffeAlphus')
- [**josopondo**:]('https://github.com/josie-opondo')
- [**sheila**:]('https://github.com/Wambita')

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```





