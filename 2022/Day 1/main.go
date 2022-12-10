package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var element int
var maxTotalCalories int = 0

func main() {
	allCalories := []int{}
	content, err := os.ReadFile("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	elves := strings.Split(string(content), "\n\n")
	// fmt.Println(v[1])
	for i := 0; i < len(elves); i++ {
		elfCalories := strings.Split(elves[i], "\n")
		var totalCalories int = 0
		for _, element := range elfCalories {
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
	var topThreeSum = 0
	for _, element := range allCalories[:3] {
		topThreeSum += element
	}
	var topCalories = 0
	for _, element := range allCalories[:1] {
		topCalories += element
	}
	fmt.Println("Top 3 sum: ", topThreeSum)
	fmt.Println(topCalories)
}
