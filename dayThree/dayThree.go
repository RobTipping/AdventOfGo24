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
		//fmt.Println(sum)
		findAstrix()
		// need to move the line swaps here instead of in valid part number
		topString = middleString
		middleString = bottomString
	}
	sum += isValidPartNumber("")
	fmt.Println(sum)
	testString := "....12345.."
	testlength := len(testString)
	findNumber(5, testString, testlength)
}

func findAstrix() {
	if middleString == "" {
		return
	}
	for i, c := range middleString {
		if c == '*' {
			fmt.Println("i found an astrix")
			//howManyDigits(i)
		}
		howManyDigits(i)
	}
}

func howManyDigits(pos int) int {
	sum := 0
	topLength := len(topString)
	middleLength := len(middleString)
	bottomLength := len(bottomString)

	// check top row
	if pos < topLength && pos > 0 {
		if unicode.IsDigit(rune(topString[pos-1])) == true {
			sum += 1
		}
	}
	if pos+1 < topLength {
		// check middle
		if unicode.IsDigit(rune(topString[pos])) == true {
			sum += 1
		}
	}
	if pos+2 <= topLength {
		// check top right
		if unicode.IsDigit(rune(topString[pos+1])) == true {
			sum += 1
		}
	}

	// check middle row
	if pos < middleLength && pos > 0 {
		if unicode.IsDigit(rune(middleString[pos-1])) == true {
			sum += 1
		}
	}
	if pos+2 <= middleLength {
		if unicode.IsDigit(rune(middleString[pos+1])) == true {
			sum += 1
		// check top right
		}
	}

	// check bottom row
	if pos < bottomLength && pos > 0 {
		if unicode.IsDigit(rune(bottomString[pos-1])) == true {
			sum += 1
		}
	}
	if pos+1 < bottomLength {
		// check middle
		if unicode.IsDigit(rune(bottomString[pos])) == true {
			sum += 1
		}
	}
	if pos+2 <= bottomLength {
		// check top right
		if unicode.IsDigit(rune(bottomString[pos+1])) == true {
			sum += 1
		}
	}

	return sum
}

func findNumber(pos int, text string, length int) int {
	// find findNumber
	low := 0
	high := 0
	for x:= pos; x >= 0; x-- {
		fmt.Println(x)
		if unicode.IsDigit(rune(text[x])) != true {
			fmt.Println("did this execute")
			low = x + 1
			break
		}
	}
	fmt.Println("did this execute")
	for x:= pos; x <= length; x++ {
		if unicode.IsDigit(rune(text[x])) != true {
			high = x - 1
			break
		}
	}
	fmt.Println(low, high)
	fmt.Println(text[low:high])
	return 0
}

func isValidPartNumber(newString string) int {
	bottomString = newString
	tempstring := ""
	check := false
	sum := 0
	if middleString != "" {
		for i, c := range middleString {
			//fmt.Println("square to checks value:", string(c))
			if unicode.IsDigit(c) == true {
				tempstring += string(c)
				if check == false {
					check = nextToSymbol(i)
				}
			} else {
				if check == true {
					//fmt.Println("has a symbol above", tempstring)
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

	if check == true {
		//fmt.Println("has a symbol above", tempstring)
		check = false
		temp, err := strconv.Atoi(tempstring)
		if err == nil {
			sum += temp
		}
	}
	//topString = middleString
	//middleString = bottomString
	//println(sum)
	return sum
}

func nextToSymbol(pos int) bool {
	topLength := len(topString)
	middleLength := len(middleString)
	bottomLength := len(bottomString)

	//fmt.Println("position:", pos)
	//fmt.Println(topString)
	//fmt.Println("top corner value:", string(topString[pos+1]))

	// check top row
	if pos < topLength && pos > 0 {
		if checkSymbol(rune(topString[pos-1])) == true {
			return true
		}
	}
	if pos+1 < topLength {
		// check middle
		if checkSymbol(rune(topString[pos])) == true {
			return true
		}
	}
	if pos+2 <= topLength {
		// check top right
		if checkSymbol(rune(topString[pos+1])) == true {
			return true
		}
	}

	// check middle row
	if pos < middleLength && pos > 0 {
		if checkSymbol(rune(middleString[pos-1])) == true {
			return true
		}
	}
	if pos+2 <= middleLength {
		// check top right
		if checkSymbol(rune(middleString[pos+1])) == true {
			return true
		}
	}

	//fmt.Println("checking the bottom")
	//fmt.Println("the bottom length is", bottomLength)
	//fmt.Println("pos is", pos)
	// check bottom row
	if pos < bottomLength && pos > 0 {
		//fmt.Println("checking the bottom")
		if checkSymbol(rune(bottomString[pos-1])) == true {
			return true
		}
	}
	if pos+1 < bottomLength {
		// check middle
		if checkSymbol(rune(bottomString[pos])) == true {
			return true
		}
	}
	if pos+2 <= bottomLength {
		// check top right
		if checkSymbol(rune(bottomString[pos+1])) == true {
			return true
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
