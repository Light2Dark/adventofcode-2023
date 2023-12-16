package day16

import (
	"fmt"
	"strings"
	"os"
)

type Ray struct {
	x int
	y int
}

func Run2() {
// 	data := `.|...\....
// |.-.\.....
// .....|-...
// ........|.
// ..........
// .........\
// ..../.\\..
// .-.-/..|..
// .|....-|.\
// ..//.|....`
// 	lines := strings.Split(data, "\n")

	data, _ := os.ReadFile("day16/input.txt")
	lines := strings.Split(string(data), "\n")

	var board = [][]tile{}

	for _, line := range lines {
		var row []tile
		for _, runeVal := range line {
			var newTile = tile{string(runeVal), false}
			row = append(row, newTile)
		}
		board = append(board, row)
	}

	max := getMax(board)
	fmt.Println(max)
}

func getMax(board [][]tile) int {
	var max int

	checkAndUpdateMax := func(board [][]tile, ray Ray, direction string) {
		newBoard := makeBoardCopy(board)
		visited := map[Key]int{}
		traverseRay(newBoard, ray, direction, visited)
		result := scoreBoard(newBoard)

		if result > max {
			max = result
		}
	}

	// all positions at col 0, row 1 to second last row
	for i := 1; i < len(board)-1; i++ {
		var ray = Ray{0, i}
		checkAndUpdateMax(board, ray, leftToRight)
	}

	// last col
	for i := 1; i < len(board)-1; i++ {
		var ray = Ray{len(board) - 1, i}
		checkAndUpdateMax(board, ray, rightToLeft)
	}

	// top row
	for i := 1; i < len(board[0])-1; i++ {
		var ray = Ray{i, 0}
		checkAndUpdateMax(board, ray, upToDown)
	}

	// last row
	for i := 1; i < len(board[0])-1; i++ {
		var ray = Ray{i, len(board) - 1}
		checkAndUpdateMax(board, ray, downToUp)
	}

	// corners
	var topLeftRay = Ray{0, 0}
	checkAndUpdateMax(board, topLeftRay, upToDown)
	checkAndUpdateMax(board, topLeftRay, leftToRight)

	var topRightRay = Ray{len(board[0]) - 1, 0}
	checkAndUpdateMax(board, topRightRay, upToDown)
	checkAndUpdateMax(board, topRightRay, rightToLeft)

	var bottomLeftRay = Ray{0, len(board) - 1}
	checkAndUpdateMax(board, bottomLeftRay, downToUp)
	checkAndUpdateMax(board, bottomLeftRay, leftToRight)

	var bottomRightRay = Ray{len(board[0]) - 1, len(board) - 1}
	checkAndUpdateMax(board, bottomRightRay, downToUp)
	checkAndUpdateMax(board, bottomRightRay, rightToLeft)

	return max
}

func makeBoardCopy(board [][]tile) [][]tile {
	var boardCopy [][]tile
	for _, row := range board {
		var rowCopy []tile
		rowCopy = append(rowCopy, row...)
		boardCopy = append(boardCopy, rowCopy)
	}
	return boardCopy
}
