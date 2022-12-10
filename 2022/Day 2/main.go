package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var moveScore int
var totalScore int
var correctMoveScore int
var correctTotalScore int

const drawScore = 3
const winScore = 6
const lossScore = 0
const rockScore = 1
const paperScore = 2
const scissorsScore = 3

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	allMoves := strings.ReplaceAll(string(content), "\r\n", "\n")
	allRounds := strings.Split(allMoves, "\n")
	for i := 0; i < len(allRounds); i++ {
		theMoves := strings.Split(allRounds[i], " ")
		moveScore = yourScore(theMoves[1]) + currentPlay(theMoves[1], theMoves[0])
		totalScore += moveScore
		correctMoveScore = yourScore(neededResult(theMoves[1], theMoves[0])) + currentPlay(neededResult(theMoves[1], theMoves[0]), theMoves[0])
		correctTotalScore += correctMoveScore
	}
	// Part 1
	fmt.Println(totalScore)
	// Part 2
	fmt.Println(correctTotalScore)
}
func yourScore(yourMove string) int {
	switch yourMove {
	case "X":
		return rockScore
	case "Y":
		return paperScore
	case "Z":
		return scissorsScore
	}
	return 0
}

func currentPlay(yourMove, opponentMove string) int {
	switch yourMove {
	case "X": //rock
		if opponentMove == "A" { //rock
			return drawScore
		} else if opponentMove == "B" { //paper
			return lossScore
		}
		return winScore //scissors
	case "Y": //paper
		if opponentMove == "A" { //rock
			return winScore
		} else if opponentMove == "B" { //paper
			return drawScore
		}
		return lossScore //scissors
	case "Z": //scissors
		if opponentMove == "A" { //rock
			return lossScore
		} else if opponentMove == "B" { //paper
			return winScore
		}
		return drawScore //scissors
	}
	return 0
}

func neededResult(yourMove, opponentMove string) string {
	switch yourMove {
	case "X": //Lose
		if opponentMove == "A" { //rock
			return "Z"
		} else if opponentMove == "B" { //paper
			return "X"
		}
		return "Y" //scissors
	case "Y": //Draw
		if opponentMove == "A" { //rock
			return "X"
		} else if opponentMove == "B" { //paper
			return "Y"
		}
		return "Z" //scissors
	case "Z": //Win
		if opponentMove == "A" { //rock
			return "Y"
		} else if opponentMove == "B" { //paper
			return "Z"
		}
		return "X" //scissors
	}
	return "Error"
}
