package main

import (
	"fmt"
	"strings"
)

var redMax = 12
var greenMax = 13
var blueMax = 14

var testString = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"


func main() {
    fmt.Println("test world")
    //tempChar:= testString[5]
    fmt.Println(strings.Split(testString, ":"))
    fmt.Println(strings.SplitAfter(testString, ":"))
    firstSplit := strings.Split(testString, ": ")
    fmt.Println(firstSplit[0])
    secondSplit := strings.Split(firstSplit[1], "; ")
    for _, x := range secondSplit {
        fmt.Println(x)
        for _, y := range strings.Split(x, ", ") {
            fmt.Println(y)
        }
    }
}
