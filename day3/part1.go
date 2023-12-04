package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	data, err := os.ReadFile("day3/input.txt")
	checkError(err)

	var lines []string = strings.Split(string(data), "\n")

	var board [][]rune
	var partNums []int

	// populating the board with input
	for _, line := range lines {
		var row []rune = []rune(line)
		board = append(board, row)
	}

	// adding empty row before & after
	var emptyRow = []rune(strings.Repeat(".", len(board[0])))
	board = append([][]rune{emptyRow}, board...) // as the first arg must be a slice of slices, as we are adding slices to the slice of slices
	board = append(board, emptyRow)              // adding another slice, to slice of slices

	// adding empty column before & after
	for i := range board {
		board[i] = append([]rune("."), board[i]...)
		board[i] = append(board[i], '.')
	}

	// compute based on inner board (ignore padding)
	for i := 1; i < len(board)-1; i++ {
		for m := 1; m < len(board[i])-1; m++ {
			var runeVal rune = board[i][m]
			if unicode.IsNumber(runeVal) {
				symbolPresent, indexToSkip, numFound := isSurroundingSymbolPresent(i, m, board)
				fmt.Println(symbolPresent, indexToSkip)

				if indexToSkip > 0 {
					m += indexToSkip - 1
				}

				if numFound != 0 {
					partNums = append(partNums, numFound)
				}
			}
		}
	}

	var sum int
	for _, val := range partNums {
		sum += val
	}
	fmt.Println(sum)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isSurroundingSymbolPresent(row int, colStart int, board [][]rune) (isPresent bool, indexToSkip int, numberFound int) {
	// get the entire number
	var reFindDigits = regexp.MustCompile(`\d+`)

	var stringTillEnd string = string(board[row][colStart:])
	var numbers []string = reFindDigits.FindStringSubmatch(stringTillEnd)
	var number []rune = []rune(numbers[0])

	var surroundedBySymbol bool

	for i := range number {
		var leftRune, rightRune = board[row][colStart-1+i], board[row][colStart+1+i]
		var upRune, downRune = board[row-1][colStart+i], board[row+1][colStart+i]
		var topLeftRune, topRightRune = board[row-1][colStart-1+i], board[row-1][colStart+1+i]
		var bottomLeftRune, bottomRightRune = board[row+1][colStart-1+i], board[row+1][colStart+1+i]

		var surroundingRunes []rune = []rune{leftRune, rightRune, upRune, downRune, topLeftRune, topRightRune, bottomLeftRune, bottomRightRune}

		var output = map[string]string{
			"number":      string(number),
			"left":        string(leftRune),
			"right":       string(rightRune),
			"up":          string(upRune),
			"down":        string(downRune),
			"topLeft":     string(topLeftRune),
			"topRight":    string(topRightRune),
			"bottomLeft":  string(bottomLeftRune),
			"bottomRight": string(bottomRightRune),
		}
		fmt.Println(output)

		var re *regexp.Regexp = regexp.MustCompile(`[^\.\s\w]`) // finds symbols
		for _, runeSurround := range surroundingRunes {
			if re.MatchString(string(runeSurround)) {
				surroundedBySymbol = true
				break
			}
		}

		if surroundedBySymbol {
			break
		}
	}

	if surroundedBySymbol {
		num, err := strconv.Atoi(string(number))
		checkError(err)
		return surroundedBySymbol, len(number), num
	} else {
		return false, 0, 0
	}
}
