package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const chunkSize = 3

type Rucksack struct {
	compartment1 string
	compartment2 string
	commonItem   rune
	itemPriority int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	allContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	allRucksacks := strings.Split(allContent, "\n")
	chunks := splitIntoChunks(allRucksacks)
	// fmt.Println(allRucksacks[0])
	var allCompartments []Rucksack
	for _, rucksack := range allRucksacks {
		commonItem, err := findCommonCharacter(rucksack[:(len(rucksack)/2)], rucksack[(len(rucksack)/2):])
		if err != nil {
			log.Fatal(err)
		}
		itemPriority, err := getItemPriority(commonItem)
		if err != nil {
			log.Fatal(err)
		}
		allCompartments = append(
			allCompartments,
			Rucksack{
				compartment1: rucksack[:(len(rucksack) / 2)],
				compartment2: rucksack[(len(rucksack) / 2):],
				commonItem:   commonItem,
				itemPriority: itemPriority,
			})
	}
	var itemSum = 0
	for _, rucksack := range allCompartments {
		itemSum += rucksack.itemPriority
	}
	fmt.Println(itemSum) //part 1
	var badgeSum = 0
	for _, chunk := range chunks {
		badge, err := findCommonBadge(chunk)
		if err != nil {
			log.Fatal(err)
		}
		priority, err := getItemPriority(badge)
		if err != nil {
			log.Fatal(err)
		}
		badgeSum += priority
	}
	fmt.Println(badgeSum)
}

func findCommonCharacter(compOne, compTwo string) (rune, error) {
	for _, item := range compOne {
		if strings.Contains(compTwo, string(item)) {
			return item, nil
		}
	}
	return rune(0), errors.New("nothing found")
}

func getItemPriority(item rune) (int, error) {
	if item >= 97 && item <= 122 {
		return int(item) - 96, nil
	} else if item >= 65 && item <= 90 {
		return int(item) - 38, nil
	}
	return 0, errors.New("wrong index")
}

func splitIntoChunks(input []string) [][]string {
	var currentChunkSize = 0
	var chunk []string
	var output [][]string
	for _, line := range input {
		currentChunkSize++
		chunk = append(chunk, line)
		if currentChunkSize == 3 {
			output = append(output, chunk)
			chunk = []string{}
			currentChunkSize = 0
		}
	}
	return output
}

func findCommonBadge(chunk []string) (rune, error) {
	for _, item := range chunk[0] {
		if strings.Contains(chunk[1], string(item)) {
			if strings.Contains(chunk[2], string(item)) {
				return item, nil
			}
		}
	}
	return rune(0), errors.New("nothing found")
}
