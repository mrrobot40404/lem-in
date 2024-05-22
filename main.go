
package main

import (
	"bufio"
	"strings"
	// "fmt"
	// "log"
	"os"
	// "reflect"
	// "regexp"
	// "strconv"
	// "strings"
)

type Room struct {
	name string
	x, y int //coordinates for visualization
	visited bool // if visited already, then ignore on second pass of bfs exploration
	path []*Room // rooms linking to this room struct 
}

type Path struct {
	roomsInPath []*Room // list of rooms in this path
	numberOfAnts int
}

type AntFarm struct {
	rooms map[string]*Room
	paths []*Path
	startRoom, endRoom Room
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
		
		if strings.HasPrefix(line, "#") { // If the line has a prefix with # we check if its the pending Room type either start or end
			if line == "##start" {
				pendingType = "start"
			} else if line == "##end" {
				pendingType = "end"
			}
			continue // We skip if we have comments or anything that is not a starting or ending room
		}

		
	}	



	return antFarm, nil
}

