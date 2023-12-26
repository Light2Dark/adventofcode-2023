package day21

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// after 16 steps, count the number of places he could be at

type Position struct {
	x int
	y int
}

func Run() {
	filepath := flag.String("file", "day21/test1.txt", "path to filename")
	flag.Parse()

	data, _ := os.ReadFile(*filepath)
	lines := strings.Split(string(data), "\n")

	var board [][]rune
	var startPos Position

	for i, line := range lines {
		var row = []rune{}
		for j, val := range line {
			row = append(row, val)
			if byte(val) == 'S' {
				startPos = Position{x: j, y: i}
			}
		}
		board = append(board, row)
	}

	board = padBoard(board)
	startPos.x, startPos.y = startPos.x+1, startPos.y+1
	var cache = map[[3]int]bool{} // key: x,y, num of steps

	fmt.Println(traverse(board, startPos, 64, cache))

}

func traverse(board [][]rune, position Position, numStepsLeft int, cache map[[3]int]bool) int {

	var key = [3]int{position.x, position.y, numStepsLeft}
	if _, ok := cache[key]; ok {
		return 0
	}
	cache[key] = true // this place is now visited, put it above the return statements so that in the next iteration it doesn't count this

	if board[position.y][position.x] == '#' {
		return 0
	}

	if numStepsLeft == 0 {
		return 1
	}

	var count int
	var possibleMoves = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // x,y
	for _, move := range possibleMoves {
		newPosition := Position{x: position.x+move[0], y: position.y+move[1]}
		count += traverse(board, newPosition, numStepsLeft-1, cache)
	}
	
	return count
}

func padBoard(board [][]rune) [][]rune {
	var emptyRow []rune = []rune(strings.Repeat("#", len(board[0])))
	board = append(board, emptyRow)
	board = append([][]rune{emptyRow}, board...)

	for i := range board {
		board[i] = append([]rune("#"), board[i]...)
		board[i] = append(board[i], rune('#'))
	}

	return board
}

func printBoard(board [][]rune) {
	for _, row := range board {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}
