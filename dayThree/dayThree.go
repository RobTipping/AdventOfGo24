package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var topString = ""
var middleString = ""
var bottomString = ""

func main() {
    sum := 0
    dat, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    defer dat.Close()
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
	sum += isValidPartNumber(scanner.Text())	
    }
    sum += isValidPartNumber("")
    fmt.Println(sum)
}

func isValidPartNumber(newString string) int {
    bottomString = newString
    tempstring := ""
    check := false
    sum := 0
    if middleString != "" {
	for i, c := range middleString {
	    if unicode.IsDigit(c) == true {
		tempstring += string(c)
		if check == false {
		    check = nextToSymbol(i)
		}
	    } else {
		if check == true {
		    fmt.Println("has a symbol above", tempstring)
		    check = false
		    temp, err := strconv.Atoi(tempstring)
		    if err == nil {
			sum += temp
		    }
		} else {
		    //fmt.Println("does not have a symbol", tempstring)
		}
		//fmt.Println(tempstring)
		tempstring = ""
		check = false
	    }
	}
    }
    topString = middleString
    middleString = bottomString
    println(sum)
    return sum
}

func nextToSymbol(pos int) bool {
    topLength := len(topString)
    middleLength := len(middleString)
    bottomLength := len(bottomString)
    

    // check top row
    if pos < topLength && pos > 0{
	if checkSymbol(rune(topString[pos-1])) == true {
	    return true;
	}
    }
    if pos + 1 < topLength {
	// check middle
	if checkSymbol(rune(topString[pos])) == true {
	    return true;
	}
    }
    if pos + 2 < topLength {
	// check top right
	if checkSymbol(rune(topString[pos+1])) == true {
	    return true;
	}
    }

    // check middle row
    if pos < middleLength && pos > 0{
	if checkSymbol(rune(middleString[pos-1])) == true {
	    return true;
	}
    }
    if pos + 2 < middleLength {
	// check top right
	if checkSymbol(rune(middleString[pos+1])) == true {
	    return true;
	}
    }

	//fmt.Println("checking the bottom")
    //fmt.Println("the bottom length is", bottomLength)
    //fmt.Println("pos is", pos)
    // check bottom row
    if pos < bottomLength && pos > 0{
	//fmt.Println("checking the bottom")
	if checkSymbol(rune(bottomString[pos-1])) == true {
	    return true;
	}
    }
    if pos + 1 < bottomLength {
	// check middle
	if checkSymbol(rune(bottomString[pos])) == true {
	    return true;
	}
    }
    if pos + 2 < bottomLength {
	// check top right
	if checkSymbol(rune(bottomString[pos+1])) == true {
	    return true;
	}
    }

    return false
}

func checkSymbol(character rune) bool {
    //fmt.Println("is testing the rune: ", string(character))
    //fmt.Println(string(character))
    if unicode.IsDigit(character) == false && string(character) != "." {
	//fmt.Println("is a symbol: ", string(character))
	return true
    }

    return false
}
