package main

// TODO:
//// parse input file
//// print number of ants
//// print rooms
//// print start room
//// print end room
//// print links
// Make struct for rooms and links
// create graph
// find shortest path
// print moves

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Room struct {
	name string
	x    int
	y    int
}

type Link struct {
	room1 string
	room2 string
}

type Graph struct {
	rooms map[Room][]string
}

func (g *Graph) addRoom(r Room) {
	g.rooms[r] = []string{}
}

func (g *Graph) addLink(l Link) {
	for room := range g.rooms {
		if room.name == l.room1 {
			g.rooms[room] = append(g.rooms[room], l.room2)
		}
		if room.name == l.room2 {
			g.rooms[room] = append(g.rooms[room], l.room1)
		}
	}
}

func (g *Graph) Display() {
	for room, links := range g.rooms {
		fmt.Printf("%+v -> %v\n", room, links)
	}
}

func main() {
	// Read os args
	if len(os.Args[1:]) != 1 {
		log.Fatalln("[USAGE]: ./lem-in <filename>.txt")
	}
	filename := os.Args[1]
	// Read input file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("[USAGE]: ./lem-in <filename>.txt")
	}
	defer file.Close()
	// Parse input file
	file_lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		file_lines = append(file_lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the number of ants
	num_ants, err := strconv.Atoi(file_lines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of ants:", reflect.TypeOf(num_ants), num_ants)

	// Print the rooms
	r1, err := regexp.Compile(`^([a-zA-Z0-9]+)\s(\d+)\s(\d+)$`)
	// '^([a-zA-Z0-9]+) (\d+) (\d+)$'
	if err != nil {
		log.Fatal(err)
	}
	rStartRoom, err := regexp.Compile(`^##start$`)
	if err != nil {
		log.Fatal(err)
	}
	rEndRoom, err := regexp.Compile(`^##end$`)
	if err != nil {
		log.Fatal(err)
	}

	lastLineRoom := 0
	startRoom := ""
	endRoom := ""
	rooms_list := []Room{}
	for i := 1; i < len(file_lines); i++ {
		if r1.MatchString(file_lines[i]) {
			// fmt.Println("Room:", file_lines[i])
			room_details := strings.Split(file_lines[i], " ")
			room_name := room_details[0]
			room_x, err := strconv.Atoi(room_details[1])
			if err != nil {
				log.Fatal(err)
			}
			// TESTING
			// fmt.Println("Room X:", room_x)
			room_y, err := strconv.Atoi(room_details[2])
			if err != nil {
				log.Fatal(err)
			}
			// TESTING
			// fmt.Println("Room Y:", room_y)
			rooms_list = append(rooms_list, Room{name: room_name, x: room_x, y: room_y})
		} else if rStartRoom.MatchString(file_lines[i]) {
			room_details := strings.Split(file_lines[i+1], " ")
			startRoom = room_details[0]
		} else if rEndRoom.MatchString(file_lines[i]) {
			room_details := strings.Split(file_lines[i+1], " ")
			endRoom = room_details[0]
		} else {
			if len(r1.FindStringIndex(file_lines[i-1])) > 0 && r1.FindStringIndex(file_lines[i-1])[1] > lastLineRoom {
				lastLineRoom = r1.FindStringIndex(file_lines[i-1])[1]
			}
			break
		}
	}

	// Print the start room and end room
	fmt.Println("Start Room:", startRoom)
	fmt.Println("End Room:", endRoom)
	fmt.Println("Line Of Last Room:", lastLineRoom)

	// Print the links
	links_list := []Link{}
	r2, err := regexp.Compile(`^([a-zA-Z0-9]+)-([a-zA-Z0-9]+)$`)
	if err != nil {
		log.Fatal(err)
	}
	for i := lastLineRoom; i < len(file_lines); i++ {
		if r2.MatchString(file_lines[i]) {
			// fmt.Println("Link:", file_lines[i])
			link_details := strings.Split(file_lines[i], "-")
			room1 := link_details[0]
			room2 := link_details[1]
			links_list = append(links_list, Link{room1: room1, room2: room2})
		}
	}
	// Create graph
	graph := Graph{rooms: make(map[Room][]string)}
	for _, room := range rooms_list {
		graph.addRoom(room)
	}
	for _, link := range links_list {
		graph.addLink(link)
	}
	// fmt.Println("Graph:", rooms_list)
	// fmt.Println("Graph:", links_list)
	fmt.Println("Graph:", graph)
	graph.Display()
	for room, link := range graph.rooms {
		if room.name == startRoom {
			fmt.Print("Start Room:")
			fmt.Printf("%+v -> %v\n", room, link)
		}
		if room.name == endRoom {
			fmt.Print("End Room:")
			fmt.Printf("%+v -> %v\n", room, link)
		}
	}
	// Find shortest path
	// Print moves

}
