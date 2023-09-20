package main

import (
	"fmt"
	"log"
	"os"
)

// getInput saves the input file into a byte slices
func getInput(input string) []byte {
	// Read input into a slice of bytes
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func partOne(content []byte) int {
	// Declare local variables
	var currentFloor int = 0
	// Iterate over a slice of b ytes
	for _, v := range content {
		if v == 40 {
			currentFloor++
		} else {
			currentFloor--
		}
	}
	// Return the result
	return currentFloor
}

func partTwo(content []byte) int {
	// Declare local variables
	var (
		currentFloor      int = 0
		characterPosition int = 0
	)
	// Iterate over a slice of bytes
	for i, v := range content {
		if v == 40 {
			currentFloor++
		} else {
			currentFloor--
		}
		// Check the current floor
		if currentFloor < 0 {
			characterPosition = i + 1
			break
		}
	}
	// Return the result
	return characterPosition
}

func main() {

	content := getInput("input.txt")
	answerOne := partOne(content)
	answerTwo := partTwo(content)

	fmt.Printf("Santa came up to the %v floor\n", answerOne)
	fmt.Printf("The character position is %v", answerTwo)
}
