package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Run2() {
	data, err := os.ReadFile("day3/input.txt")
	checkError(err)

	var lines []string = strings.Split(string(data), "\n")

	var board [][]rune
	var sum int

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
			if string(runeVal) == "*" {
				var gears []int = getSurroundingGears(i, m, board)
				// fmt.Println(gears)

				if len(gears) == 2 {
					product := gears[0] * gears[1]
					sum += product
				}
			}
		}
	}

	fmt.Println(sum)
}

type runePos struct {
	row int
	col int
}

func getSurroundingGears(row int, col int, board [][]rune) []int {
	// start top left and move clockwise
	// for each rune, add digits to left and digits to right recursively(?)
	// if number in visited set, continue. else, add to visited set

	var visitedSlice []int

	var runesToVisit []runePos = []runePos{
		{row - 1, col - 1}, // top left
		{row - 1, col},     // up
		{row - 1, col + 1}, // top right
		{row, col - 1},     // left
		{row, col + 1},     // right
		{row + 1, col - 1}, // bottom left
		{row + 1, col},     // down
		{row + 1, col + 1}, // bottom right
	}

	for _, runeVisitPos := range runesToVisit {
		r, c := runeVisitPos.row, runeVisitPos.col
		var digitsToLeft = getLeftDigits(r, c, board)
		var digitsToRight = getRightDigits(r, c, board)

		var gear int
		if digitsToLeft != "" && digitsToRight != "" {
			numString := digitsToLeft[0:len(digitsToLeft)-1] + digitsToRight
			num, err := strconv.Atoi(numString)
			checkError(err)
			gear = num
		} else if digitsToLeft == "" && digitsToRight == "" {
			continue
		} else if digitsToLeft == "" {
			num, err := strconv.Atoi(digitsToRight)
			checkError(err)
			gear = num
		} else if digitsToRight == "" {
			num, err := strconv.Atoi(digitsToLeft)
			checkError(err)
			gear = num
		}

		if !contains(visitedSlice, gear) {
			visitedSlice = append(visitedSlice, gear)
		}
	}

	return visitedSlice
}

func getLeftDigits(row int, col int, board [][]rune) string {
	var digits string
	for i := col; i > 0; i-- {
		runeVal := board[row][i]
		if !unicode.IsNumber(runeVal) {
			break
		} else {
			digits = string(runeVal) + digits // reversed
		}
	}

	return digits
}

func getRightDigits(row int, col int, board [][]rune) string {
	var digits string
	for i := col; i < len(board[row]); i++ {
		runeVal := board[row][i]
		if !unicode.IsNumber(runeVal) {
			break
		} else {
			digits = digits + string(runeVal)
		}
	}

	return digits
}

func contains(slice []int, elem int) bool {
	for _, val := range slice {
		if val == elem {
			return true
		}
	}
	return false
}
