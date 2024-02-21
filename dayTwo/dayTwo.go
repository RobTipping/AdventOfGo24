package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var redMax = 12
var greenMax = 13
var blueMax = 14

func main() {
	sumA := 0
	sumB := 0
	dat, err := os.Open("mainInput")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		sumA += validGame(scanner.Text())
		sumB += findMinimumCubes(scanner.Text())
	}
	fmt.Println("part A = ", sumA)
	fmt.Println("part B = ", sumB)
}

func validGame(inputString string) int {
	firstSplit := strings.Split(inputString, ": ")
	gameNumber, err := strconv.Atoi(strings.TrimPrefix(firstSplit[0], "Game "))
	if err != nil {
		fmt.Println("there is no game")
		fmt.Println(gameNumber)
		return 0
	}
	for _, x := range strings.Split(firstSplit[1], "; ") {
		for _, y := range strings.Split(x, ", ") {
			if strings.Contains(y, "red") == true {
				temp := strings.TrimSuffix(y, " red")
				r, err := strconv.Atoi(temp)
				if err == nil {
					if r > redMax {
						fmt.Println("not possible game")
						return 0
					}
				}
			} else if strings.Contains(y, "green") == true {
				temp := strings.TrimSuffix(y, " green")
				r, err := strconv.Atoi(temp)
				if err == nil {
					if r > greenMax {
						fmt.Println("not possible game")
						return 0
					}
				}
			} else if strings.Contains(y, "blue") == true {
				temp := strings.TrimSuffix(y, " blue")
				r, err := strconv.Atoi(temp)
				if err == nil {
					if r > blueMax {
						fmt.Println("not possible game")
						return 0
					}
				}
			}
		}
	}
	return gameNumber
}

func findMinimumCubes(inputString string) int {
	red := 1
	green := 1
	blue := 1
	firstSplit := strings.Split(inputString, ": ")
	gameNumber, err := strconv.Atoi(strings.TrimPrefix(firstSplit[0], "Game "))
	if err != nil {
		fmt.Println("there is no game")
		fmt.Println(gameNumber)
		return 0
	}
	for _, x := range strings.Split(firstSplit[1], "; ") {
		for _, y := range strings.Split(x, ", ") {
			if strings.Contains(y, "red") == true {
				temp := strings.TrimSuffix(y, " red")
				r, err := strconv.Atoi(temp)
				if err == nil {
					if red == 1 {
						red = r
					} else if r > red {
						red = r
					}
				}
			} else if strings.Contains(y, "green") == true {
				temp := strings.TrimSuffix(y, " green")
				r, err := strconv.Atoi(temp)
				if err == nil {
					if green == 1 {
						green = r
					} else if r > green {
						green = r
					}
				}
			} else if strings.Contains(y, "blue") == true {
				temp := strings.TrimSuffix(y, " blue")
				r, err := strconv.Atoi(temp)
				if err == nil {
					if blue == 1 {
						blue = r
					} else if r > blue {
						blue = r
					}
				}
			}
		}
	}
	return red * green * blue
}
