package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var allCalories []int
var topThree = 0

func main() {
	content, err := os.ReadFile("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	// This will make the code Windows compatible
	allElves := strings.ReplaceAll(string(content), "\r\n", "\n")
	groupedCalories := strings.Split(allElves, "\n\n")
	for i := 0; i < len(groupedCalories); i++ {
		calorieGroup := strings.Split(groupedCalories[i], "\n")
		var totalCalories = 0
		for _, element := range calorieGroup {
			element, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			totalCalories += element
		}
		allCalories = append(allCalories, totalCalories)
	}
	sort.Slice(allCalories, func(i, j int) bool {
		return allCalories[i] > allCalories[j]
	})
	// Part 1
	fmt.Printf("The Elf carrying the most Calories carries %d of them\n", allCalories[0])
	// Part 2
	for _, element := range allCalories[:3] {
		topThree += element
	}
	fmt.Printf("The top three Elves are carrying %d calories combined\n", topThree)
}
