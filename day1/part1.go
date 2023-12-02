package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
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