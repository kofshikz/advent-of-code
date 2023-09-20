package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type box struct {
	length int
	width  int
	height int
}

func (b *box) area() int {
	return 2 * (b.length*b.width + b.width*b.height + b.height*b.length)
}

func (b *box) slack() int {
	slack := b.length * b.width
	if b.width*b.height < slack {
		slack = b.width * b.height
	}
	if b.height*b.length < slack {
		slack = b.height * b.length
	}
	return slack
}

func (b *box) volume() int {
	return b.height * b.length * b.width
}

func (b *box) smallestPerimeter() int {
	smallestPerimeter := 2 * (b.length + b.width)
	if 2*(b.height+b.width) < smallestPerimeter {
		smallestPerimeter = 2 * (b.height + b.width)
	}
	if 2*(b.height+b.length) < smallestPerimeter {
		smallestPerimeter = 2 * (b.height + b.length)
	}
	return smallestPerimeter
}

// getInput returns the contents of the input file into as a string
func getInput(input string) []string {
	// Read input into a slice of bytes
	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.ReplaceAll(string(content[:]), "\r\n", "\n"), "\n")
}

func convertDimensions(d []string) []int {
	intDimensions := make([]int, len(d))
	for i, s := range d {
		intDimensions[i], _ = strconv.Atoi(s)
	}
	return intDimensions
}

func partOne(input []string) int {
	var surfaceArea = 0
	for _, d := range input {
		var box box
		dimensions := convertDimensions(strings.Split(d, "x"))
		box.length = dimensions[0]
		box.width = dimensions[1]
		box.height = dimensions[2]
		surfaceArea += (box.area() + box.slack())
	}
	return surfaceArea
}

func partTwo(input []string) int {
	var ribbonLength = 0
	for _, d := range input {
		var box box
		dimensions := convertDimensions(strings.Split(d, "x"))
		box.length = dimensions[0]
		box.width = dimensions[1]
		box.height = dimensions[2]
		ribbonLength += (box.volume() + box.smallestPerimeter())
	}
	return ribbonLength
}

func main() {
	content := getInput("input.txt")
	answerOne := partOne(content)
	answerTwo := partTwo(content)
	fmt.Printf("The total surface area needed is %v", answerOne)
	fmt.Printf("\nThe total ribbon length required is %v", answerTwo)
}
