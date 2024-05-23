package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	name string // Room Identifier (Could be digit/string/whatever)
	x, y int // Coordinates for visualization
	links []*Room // Rooms linking to this room struct 
}

// type Path struct {
// 	roomsInPath []*Room // List of rooms in this path
// 	antsInPath int // Number of ants in the
// }

type AntFarm struct {
	rooms map[string]*Room // If visited already, then ignore on second pass of bfs exploration
	numberOfAnts int // Number of ants in the ant farm
	startRoom, endRoom *Room
}

func parseFile(filepath string) (*AntFarm, error) {
	// Opening File
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err // Return error if file opening doesnt work
	}
	defer file.Close() // Close file when function is closed

	// Initializing Scanner and AntFarm structure
	
	scanner := bufio.NewScanner(file) // Scanner object created to read file line by line

	antFarm := &AntFarm{
		rooms: make(map[string]*Room),
	}

	// Storing either start or end room
	var pendingType string
	
	// Reading file line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trim the line for any trailing whitespace

		if line == "" { // If no contents in the line
			continue
		} 

		if antFarm.numberOfAnts == 0 {
			antFarm.numberOfAnts, err = strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("Invalid number of ants")
			}	
			continue
		}
		
		if strings.HasPrefix(line, "#") { // If the line has a prefix with # we check if its the pending Room type either start or end
			if line == "##start" {
				pendingType = "start"
			} else if line == "##end" {
				pendingType = "end"
			}
			continue // We skip if we have comments or anything that is not a starting or ending room
		}

		// Parsing Rooms
		if strings.Contains(line, " ") {
			parts := strings.Split(line, " ") // []string
			if len(parts) != 3 {
				return nil, fmt.Errorf("Invalid room format")
			}

			// Parse to store coordinates for bonus repository 
			x, _ := strconv.Atoi(parts[1]) // Coordinate X
			y, _ := strconv.Atoi(parts[2]) // Coordinate Y

			room := &Room{
				name: parts[0],
				x: x,
				y: y,
				links: []*Room{},
			}

			antFarm.rooms[room.name] = room // The room name identifies the room struct for O(1) lookup

			if pendingType == "start" {
				antFarm.startRoom = room // If the pendingType is populated with "start", then the room after that line directly is the start room
				pendingType = ""
			} else if pendingType == "end" { // If the pendingType is populated with "end", then the room after that line directly is the end room
				antFarm.endRoom = room
				pendingType = ""
			}
		}
		
		// Parsing Links
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid link format")
			}
			room1 := antFarm.rooms[parts[0]]
			room2 := antFarm.rooms[parts[1]]
			room1.links = append(room1.links, room2)
			room2.links = append(room2.links, room1)
		}
	}	
	
	// Error: Hits an end of file condition without reading any data
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("Scanner Error")
	}	

	// Error: Ant Farm does not have a start or end room

	if antFarm.startRoom == nil || antFarm.endRoom == nil {
		return nil, fmt.Errorf("Missing Start Room or End Room")
	}

	return antFarm, nil
}

/*
Breadth First Search answers two questions:
1) Is there a path from node A to node B?
2) What is the shortest path from node A to node B?

Using queues: search nodes in the order they were added 



Summary:
-> Initialize a queue to store paths and a map to track visited paths. 
-> 
->
->





Returns a list of paths which is a list of rooms, so a 2d room array
*/

func bfsTraversal(antFarm *AntFarm) [][]*Room{
	// 1) Initializing a queue to keep track of explored paths
	// Each element in the queue is a path (a list of rooms)
	// -> Starting Point: Enqueue start room
	var paths [][]*Room
	queue := [][]*Room{{antFarm.startRoom}} 
	visited := map[string]bool{antFarm.startRoom.name:true}

	// We would like to prevent cycles in the farm and we use a visited map by doing so
	for len(queues) != 0 {
		path := queue[0] // extract from the front of the queue 
		queue = queue[1:] // mechanism to dequeue from the front of the queue
		if visited[]
	}
	return paths
}

func main() {

}
