package main

import (
	"bufio"
	"fmt"
	"os"
)

var testString1 = ""
var testString2 = "467..114.."

func main() {
    sum := 0
    dat, err := os.Open("testInput")
    if err != nil {
        panic(err)
    }
    defer dat.Close()
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {

    }
    fmt.Println(sum)



    fmt.Println("test to the world")
    if testString1 == "" {
        fmt.Println("blank string")
    }
}
