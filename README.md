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
	- use **breadth-first search** BFS to find the shortest path from `##start` to `##end`
	- consider addional paths for parallel traversal to minimize moves

- CHALLENGES:
	- avoid loops and self-links
	- handle multiple shortest paths efficiently

4. SIMULATION
- Ant Movement Rules:
	- each ant can move once per turn
	- a room (except `##start` and `##end`) can only hold one ant at a time.
	- ants follow the shortest paths but avoid congestion




