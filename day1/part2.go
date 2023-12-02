package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var lettersToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// word var, iterate through line, append to word if first char matches first letter in map. If char matches last letter in map, then use the int.

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func Run2() {
	data, err := os.ReadFile("day1/input.txt")
	checkError(err)

	var lines []string = strings.Split(string(data), "\n")

	var calibrationValues []int

	for _, line := range lines {
		var characters []rune = []rune(line)
		var digits []int
		
		for i := range characters {
			runeVal := characters[i]
			if unicode.IsNumber(runeVal) {
				digits = append(digits, int(runeVal - '0'))
			} else {
				stringSoFar := string(characters[i:])
				for key, val := range lettersToInt {
					if strings.HasPrefix(stringSoFar, key) {
						digits = append(digits, val)
					}
				}
			}
		}

		var strVal string = strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits)-1])
		res, err := strconv.Atoi(strVal)
		checkError(err)
		calibrationValues = append(calibrationValues, res)
	}

	var sum int
	for _, val := range calibrationValues {
		sum += val
	}
	fmt.Println(sum)
}
