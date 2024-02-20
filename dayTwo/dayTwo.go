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
    sum := 0
    dat, err := os.Open("testInput")
    if err != nil {
        panic(err)
    }
    defer dat.Close()
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
	sum += validGame(scanner.Text())
    }
    fmt.Println(sum)
}

func validGame(inputString string) int {
    firstSplit := strings.Split(inputString, ": ")
    gameNumber, err := strconv.Atoi(strings.TrimPrefix(firstSplit[0], "Game "))
    if err != nil {
        fmt.Println("there is no game")
        fmt.Println(gameNumber)
    }
    for _, x := range strings.Split(firstSplit[1], "; ") {
        for _, y := range strings.Split(x, ", ") {
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
