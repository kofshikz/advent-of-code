package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const INPUT string = "iwrupvqb"

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func partOne(input string) int {
	notFiveZeroes := true
	number := 0
	for notFiveZeroes {
		hash := getMD5Hash(input + strconv.Itoa(number))
		if strings.HasPrefix(hash, "00000") {
			notFiveZeroes = false
			return number
		}
		number++
	}
	return number
}

func partTwo(input string) int {
	notSixZeroes := true
	number := 0
	for notSixZeroes {
		hash := getMD5Hash(input + strconv.Itoa(number))
		if strings.HasPrefix(hash, "000000") {
			notSixZeroes = false
			return number
		}
		number++
	}
	return number
}

func main() {
	fmt.Println("The lowest possible number to produce a hash for part one is", partOne(INPUT))
	fmt.Println("The lowest possible number to produce a hash for part two is", partTwo(INPUT))
}
