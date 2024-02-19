package main

import (
	"fmt"
	"strconv"
	"strings"
)

var redMax = 12
var greenMax = 13
var blueMax = 14

var testString = "Game 1: 3 blue, 4 red; 14 red, 2 green, 6 blue; 2 green"


func main() {
    fmt.Println("test world")
    fmt.Println(validGame(testString))
    //tempChar:= inputString[5]
}

func validGame(inputString string) int {
    firstSplit := strings.Split(inputString, ": ")
    fmt.Println(firstSplit[0])
    gameNumber, err := strconv.Atoi(strings.TrimPrefix(firstSplit[0], "Game "))
    if err == nil {
        fmt.Println("there is no game")
        fmt.Println(gameNumber)
    }
    secondSplit := strings.Split(firstSplit[1], "; ")
    for _, x := range secondSplit {
        fmt.Println(x)
        for _, y := range strings.Split(x, ", ") {
            fmt.Println(y)
            if strings.Contains(y, "red") == true {
                temp := strings.TrimSuffix(y," red")
                r, err := strconv.Atoi(temp)
                if err == nil {
                    if r > redMax {
                        fmt.Println("not possible game")
                        return 0
                    }
                }
            } else if strings.Contains(y, "green") == true {
                temp := strings.TrimSuffix(y," green")
                r, err := strconv.Atoi(temp)
                if err == nil {
                    if r > greenMax {
                        fmt.Println("not possible game")
                        return 0
                    }
                }
            } else if strings.Contains(y, "blue") == true {
                temp := strings.TrimSuffix(y," blue")
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
