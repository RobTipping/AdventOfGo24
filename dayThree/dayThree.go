package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var topString = ""
var middleString = ""
var bottomString = ""

func main() {
    //sum := 0
    dat, err := os.Open("testInput")
    if err != nil {
        panic(err)
    }
    defer dat.Close()
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
	isValidPartNumber(scanner.Text())	
    }
    isValidPartNumber("")
    //fmt.Println(sum)
}

func isValidPartNumber(newString string) int {
    bottomString = newString
    tempstring := ""
    check := false
    if middleString != "" {
	for i, c := range middleString {
	    if unicode.IsDigit(c) == true {
		tempstring += string(c)
		if check == false {
		    check = nextToSymbol(i)
		}
	    } else {
		if check == true {
		    fmt.Println("has a number above", tempstring)
		    check = false
		    // add to sum
		}
		//fmt.Println(tempstring)
		tempstring = ""
		check = false
	    }
	}
    }
    topString = middleString
    middleString = bottomString
    return 0
}

func nextToSymbol(pos int) bool {
    topLength := len(topString)
    middleLength := len(middleString)
    bottomLength := len(bottomString)

    // check top row
    if pos < topLength && pos < 0{
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
    if pos < middleLength && pos < 0{
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

    // check bottom row
    if pos < bottomLength && pos < 0{
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
    if unicode.IsDigit(character) == false && string(character) != "." {
	//fmt.Println("is a symbol: ", string(character))
	return true
    }

    return false
}
