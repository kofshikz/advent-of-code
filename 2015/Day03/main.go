package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinates struct {
	xcoord int
	ycoord int
}

// Init is the main function that initiate the structure, and return it
func (c coordinates) init() coordinates {
	c.xcoord = 0
	c.ycoord = 0
	return c
}

func (c *coordinates) move(dir string) {
	switch dir {
	case "^":
		c.ycoord++
	case "v":
		c.ycoord--
	case ">":
		c.xcoord++
	case "<":
		c.xcoord--
	default:
		log.Fatalln("Error: unknown move")
	}
}

func removeDuplicates(slice []coordinates) []coordinates {
	seen := make(map[coordinates]bool)
	unique := []coordinates{}

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func getInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content[:]), "")
}

func partOne(input []string) int {
	coords := new(coordinates).init()
	path := []coordinates{coords}
	for _, v := range input {
		coords.move(v)
		path = append(path, coords)
	}
	return len(removeDuplicates(path))
}

func partTwo(input []string) int {
	sCoords := new(coordinates).init()
	rsCoords := new(coordinates).init()
	sPath := []coordinates{sCoords}
	rsPath := []coordinates{rsCoords}
	for i, v := range input {
		if i%2 == 0 {
			sCoords.move(v)
			sPath = append(sPath, sCoords)
		} else {
			rsCoords.move(v)
			rsPath = append(rsPath, rsCoords)
		}
	}
	combinedPath := append(sPath, rsPath...)
	return len(removeDuplicates(combinedPath))
}

func main() {
	content := getInput("input.txt")
	fmt.Printf("Santa has visited %v houses", partOne(content))
	fmt.Printf("\nSanta and Robosanta visited %v houses", partTwo(content))
}
