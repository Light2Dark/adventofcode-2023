package day21

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Run2() {
	filepath := flag.String("file", "day21/input.txt", "path to filename")
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

	var cache = map[BoardCache]bool{} // key: x,y, num of steps
	var boardPosition = Position{x: 0, y: 0}
	var steps int = 26501365

	if len(board) != len(board[0]) {
		panic(errors.New("grid is not a square"))
	}
	var size = len(board)
	var half = size / 2

	var halfCount = traverseEndless(board, startPos, boardPosition, half, cache)
	var fullHalfCount = traverseEndless(board, startPos, boardPosition, half+size, cache)
	var overFullCount = traverseEndless(board, startPos, boardPosition, half+2*size, cache)
	
	a := (overFullCount + halfCount - 2*fullHalfCount) / 2
	b := fullHalfCount - halfCount - a
	c := halfCount
	n := steps / size
	result := a*n*n + b*n + c

	fmt.Println(result)
}

type BoardCache struct {
	current      Position
	numStepsLeft int
	board        Position
}

func traverseEndless(board [][]rune, position Position, boardPosition Position, numStepsLeft int, cache map[BoardCache]bool) int {
	if position.y < 0 {
		position.y = len(board) - 1
		boardPosition.y -= 1
	} else if position.y > len(board)-1 {
		position.y = 0
		boardPosition.y += 1
	}

	if position.x < 0 {
		position.x = len(board[position.y]) - 1
		boardPosition.x -= 1
	} else if position.x > len(board[position.y])-1 {
		position.x = 0
		boardPosition.x += 1
	}

	var key = BoardCache{
		current:      Position{x: position.x, y: position.y},
		numStepsLeft: numStepsLeft,
		board:        Position{x: boardPosition.x, y: boardPosition.y},
	}
	if _, ok := cache[key]; ok {
		return 0
	}
	cache[key] = true

	if board[position.y][position.x] == '#' {
		return 0
	}

	if numStepsLeft == 0 {
		return 1
	}

	var count int
	var possibleMoves = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // x,y
	for _, move := range possibleMoves {
		newPosition := Position{x: position.x + move[0], y: position.y + move[1]}
		count += traverseEndless(board, newPosition, boardPosition, numStepsLeft-1, cache)
	}

	return count
}
