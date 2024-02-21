package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"unicode"
)

func main() {
	sum := 0
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		sum += findNumbers(scanner.Text())
	}
	fmt.Println(sum)
}

func findNumbers(inputString string) int {

	var firstNumber int = 0
	var secondNumber int = 0

	for i, c := range inputString {
		if unicode.IsDigit(c) == true {
			firstNumber = (int(c) - 48)
			break
		}
		firstNumber = findTextNumber(inputString[i:])
		if firstNumber != 0 {
			break
		}
	}

	tempArray := []byte(inputString)
	slices.Reverse(tempArray)
	inputString = string(tempArray)

	for i, c := range inputString {
		if unicode.IsDigit(c) == true {
			secondNumber = (int(c) - 48)
			break
		}
		secondNumber = findBackwardsTextNumber(inputString[i:])
		if secondNumber != 0 {
			break
		}
	}

	return (firstNumber * 10) + secondNumber
}

func findTextNumber(stringSegment string) int {

	if len(stringSegment) < 6 {
		stringSegment = stringSegment + "    "
	}
	if stringSegment[0:3] == "one" {
		return 1
	} else if stringSegment[0:3] == "two" {
		return 2
	} else if stringSegment[0:5] == "three" {
		return 3
	} else if stringSegment[0:4] == "four" {
		return 4
	} else if stringSegment[0:4] == "five" {
		return 5
	} else if stringSegment[0:3] == "six" {
		return 6
	} else if stringSegment[0:5] == "seven" {
		return 7
	} else if stringSegment[0:5] == "eight" {
		return 8
	} else if stringSegment[0:4] == "nine" {
		return 9
	} else {
		return 0
	}
}

func findBackwardsTextNumber(stringSegment string) int {

	if len(stringSegment) < 6 {
		stringSegment = stringSegment + "    "
	}
	if stringSegment[0:3] == "eno" {
		return 1
	} else if stringSegment[0:3] == "owt" {
		return 2
	} else if stringSegment[0:5] == "eerht" {
		return 3
	} else if stringSegment[0:4] == "ruof" {
		return 4
	} else if stringSegment[0:4] == "evif" {
		return 5
	} else if stringSegment[0:3] == "xis" {
		return 6
	} else if stringSegment[0:5] == "neves" {
		return 7
	} else if stringSegment[0:5] == "thgie" {
		return 8
	} else if stringSegment[0:4] == "enin" {
		return 9
	} else {
		return 0
	}
}
