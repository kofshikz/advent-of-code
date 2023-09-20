package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var overlays int

func main() {
	fileToRead, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fileToRead)
	for scanner.Scan() {
		if getBitmasks(scanner.Text()) == true {
			overlays += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d overlays in assignments", overlays)
}

func getBitmasks(input string) bool {
	pair := strings.Split(input, ",")
	var assignments [][]uint8
	var bitmask []uint8
	var doesOverlay bool
	for _, assignment := range pair {
		days := strings.Split(assignment, "-")
		begin, _ := strconv.Atoi(days[0])
		end, _ := strconv.Atoi(days[1])
		bitmask = make([]uint8, 100)
		for i := range bitmask {
			if ((i + 1) >= begin) && ((i + 1) <= end) {
				bitmask[i] = 1
			} else {
				bitmask[i] = 0
			}
		}
		assignments = append(assignments, bitmask)
	}
	var bitwiseOr []uint8
	for i := 0; i < len(bitmask); i++ {
		bitwiseOr = append(bitwiseOr, assignments[0][i]|assignments[1][i])
	}
	//fmt.Println(reflect.DeepEqual(bitwiseOr, assignments[1]))
	if reflect.DeepEqual(bitwiseOr, assignments[0]) || reflect.DeepEqual(bitwiseOr, assignments[1]) {
		doesOverlay = true
	} else {
		doesOverlay = false
	}
	return doesOverlay
}
