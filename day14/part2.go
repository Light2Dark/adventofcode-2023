package day14

import (
	"fmt"
	"os"
	"strings"
)

// type Key struct {
// 	board     string
// 	direction string
// }

func Run2() {
	data, _ := os.ReadFile("day14/input.txt")
	lines := strings.Split(string(data), "\n")

// 	var sample string = `O....#....
// O.OO#....#
// .....##...
// OO.#O....O
// .O.....O#.
// O.#..O.#.#
// ..O..#O..O
// .......O..
// #....###..
// #OO..#....`

// 	lines = strings.Split(sample, "\n")

	var board = [][]rune{}
	for _, line := range lines {
		var runeLine = []rune(line)
		board = append(board, runeLine)
	}

	printBoard(board)
	board = cycle(board, 1000000000)
	printBoard(board)

	fmt.Println(calculateSum(board))
}

func cycle(board [][]rune, count int) [][]rune {
	// var mapping = make(map[string][][]rune)
	var cache = make(map[string]int) // maps boardString to cycle number (count)

	var cycleStart int
	var i = 1
	for ; i < count+1; i++ {
		oneCycle(board)
		var key = getBoardInString(board)

		var cycleNum, ok = cache[key]
		if ok {
			cycleStart = cycleNum
			break
		}

		cache[key] = i
	}

	// remainder calculation because cycle length can vary
	// count - 1 = remaining cycles
	// i - cycleStart = super-cycle period
	rem := (count - i) % (i - cycleStart)
	fmt.Println("rem", rem)
	for j := 0; j < rem; j++ {
		oneCycle(board)
	}

	return board
}

func oneCycle(board [][]rune) {
	tiltCol(board, true)
	tiltRow(board, false)
	tiltCol(board, false)
	tiltRow(board, true)
}

func calculateSum(board [][]rune) int {
	var slices []string

	// each col is a slice
	for i := 0; i < len(board[0]); i++ {
		var slice string
		for row := 0; row < len(board); row++ {
			line := []rune(board[row])
			slice += string(line[i])
		}
		slices = append(slices, slice)
	}

	var sum int
	for _, slice := range slices {
		sum += getScoreInPlace(slice)
	}

	return sum
}

func getScoreInPlace(slice string) int {
	runeLine := []rune(slice)
	var sum int
	for i := range runeLine {
		if string(runeLine[i]) == "O" {
			sum += len(runeLine) - i
		}
	}
	return sum
}

func getBoardInString(board [][]rune) string {
	var lines []string
	for _, row := range board {
		lines = append(lines, string(row))
	}
	return strings.Join(lines, "\n")
}

// possible dp is key struct having fromPosition string, tilt string, and toPosition string

// maps, slices, pointers, channels, functions and interfaces are passed by ref as default
// but slices are pointers to array with len & capacity (known as slice header)
// operations that change the slice header (eg. reslicing, appending) will create a new slice, and as such will not affect original slice

func tiltCol(board [][]rune, isNorth bool) {
	var priorityChar1, priorityChar2 = 'O', '.'
	if !isNorth {
		priorityChar1, priorityChar2 = '.', 'O'
	}

	for col := 0; col < len(board[0]); col++ {
		var newLine []rune
		var count int
		for row := 0; row < len(board); row++ {
			val := byte(board[row][col])
			if val == byte(priorityChar1) {
				newLine = append(newLine, priorityChar1)
			} else if val == byte(priorityChar2) {
				count++
			} else if val == '#' {
				newLine = addRemainingLetter(newLine, priorityChar2, count)
				newLine = append(newLine, '#')
				count = 0
			}
		}

		if count > 0 {
			newLine = addRemainingLetter(newLine, priorityChar2, count)
		}

		// replacing values
		for row := 0; row < len(board); row++ {
			board[row][col] = newLine[row]
		}
	}
}

func tiltRow(board [][]rune, isEast bool) {
	var priorityChar1, priorityChar2 rune = '.', 'O'
	if !isEast {
		priorityChar1, priorityChar2 = 'O', '.'
	}

	for row := 0; row < len(board); row++ {
		var newLine []rune
		var count int
		for col := 0; col < len(board[row]); col++ {
			val := byte(board[row][col])
			if val == byte(priorityChar1) {
				newLine = append(newLine, priorityChar1)
			} else if val == byte(priorityChar2) {
				count++
			} else if val == '#' {
				newLine = addRemainingLetter(newLine, priorityChar2, count)
				newLine = append(newLine, '#')
				count = 0
			}
		}

		if count > 0 {
			newLine = addRemainingLetter(newLine, priorityChar2, count)
		}

		board[row] = newLine
	}
}

// Deeper mapping, but it's too slow as you need to perform 4 billion hashes / checks
/*
func tiltColMapping(board [][]rune, isNorth bool, mapping map[Key][][]rune) [][]rune {
	var direction string = "north"
	if !isNorth {
		direction = "south"
	}

	var key = Key{board: getBoardInString(board), direction: direction}
	var val, ok = mapping[key]
	if ok {
		// fmt.Println("cache hit")
		return val
	}

	var priorityChar1, priorityChar2 = 'O', '.'
	if !isNorth {
		priorityChar1, priorityChar2 = '.', 'O'
	}

	for col := 0; col < len(board[0]); col++ {
		var newLine []rune
		var count int
		for row := 0; row < len(board); row++ {
			val := byte(board[row][col])
			if val == byte(priorityChar1) {
				newLine = append(newLine, priorityChar1)
			} else if val == byte(priorityChar2) {
				count++
			} else if val == '#' {
				newLine = addRemainingLetter(newLine, priorityChar2, count)
				newLine = append(newLine, '#')
				count = 0
			}
		}

		if count > 0 {
			newLine = addRemainingLetter(newLine, priorityChar2, count)
		}

		// replacing values
		for row := 0; row < len(board); row++ {
			board[row][col] = newLine[row]
		}
	}

	mapping[key] = board
	return board
}

func tiltRowMapping(board [][]rune, isEast bool, mapping map[Key][][]rune) [][]rune {
	var direction string = "east"
	if !isEast {
		direction = "west"
	}

	var key = Key{board: getBoardInString(board), direction: direction}
	var val, ok = mapping[key]
	if ok {
		// fmt.Println("cache hit")
		return val
	}

	var priorityChar1, priorityChar2 rune = '.', 'O'
	if !isEast {
		priorityChar1, priorityChar2 = 'O', '.'
	}

	for row := 0; row < len(board); row++ {
		var newLine []rune
		var count int
		for col := 0; col < len(board[row]); col++ {
			val := byte(board[row][col])
			if val == byte(priorityChar1) {
				newLine = append(newLine, priorityChar1)
			} else if val == byte(priorityChar2) {
				count++
			} else if val == '#' {
				newLine = addRemainingLetter(newLine, priorityChar2, count)
				newLine = append(newLine, '#')
				count = 0
			}
		}

		if count > 0 {
			newLine = addRemainingLetter(newLine, priorityChar2, count)
		}

		board[row] = newLine
	}

	mapping[key] = board
	return board
}
*/

func addRemainingLetter(slice []rune, letter rune, count int) []rune {
	for i := 0; i < count; i++ {
		slice = append(slice, letter)
	}
	return slice
}

func printBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(string(row))
	}
	fmt.Println()
}
