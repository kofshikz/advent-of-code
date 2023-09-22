package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// getInput returns the contents of the input file into as a string
func getInput(input string) []string {
	// Read input into a slice of bytes
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.ReplaceAll(string(content[:]), "\r\n", "\n"), "\n")
}

func hasBadStrings(s string) bool {
	badStrings := []string{"ab", "cd", "pq", "xy"}
	for _, v := range badStrings {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func hasCoupleLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasThreeVowles(s string) bool {
	vowelCount := 0
	for _, v := range s {
		if strings.ContainsAny(string(v), "aeiou") {
			vowelCount++
		}
	}
	if vowelCount >= 3 {
		return true
	}
	return false
}

func hasLetterBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func hasTwoPairs(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if strings.Contains(string(s[i+2:]), string(s[i:i+2])) {
			return true
		}
	}
	return false
}

func partOne(input []string) int {
	niceStrings := 0
	for _, v := range input {
		if hasBadStrings(v) {
			continue
		}
		if hasThreeVowles(v) && hasCoupleLetters(v) {
			niceStrings++
		}
	}
	return niceStrings
}

func partTwo(input []string) int {
	niceStrings := 0
	for _, v := range input {
		if hasLetterBetween(v) && hasTwoPairs(v) {
			niceStrings++
		}
	}
	return niceStrings
}

func main() {
	content := getInput("input.txt")
	fmt.Printf("There are %v nice strings in a first part", partOne(content))
	fmt.Printf("\nThere are %v nice strings in a first part", partTwo(content))
}
