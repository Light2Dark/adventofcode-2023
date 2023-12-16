package day16

import (
	"fmt"
	"os"
	"strings"
)

const (
	leftToRight = "leftToRight"
	upToDown    = "upToDown"
	rightToLeft = "rightToLeft"
	downToUp    = "downToUp"
)

type tile struct {
	char      string
	energized bool
}

type Key struct {
	x         int
	y         int
	direction string
}

func Run() {
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

	var ray = struct{ x, y int }{0, 0}
	visited := map[Key]int{}
	traverseRay(board, ray, leftToRight, visited)
	printEnergizedBoard(board)
	fmt.Println(scoreBoard(board))
	writeBoard(board, "day16/output.txt", false)
}

func scoreBoard(board [][]tile) int {
	var sum int
	for _, row := range board {
		for _, boardTile := range row {
			if boardTile.energized {
				sum += 1
			}
		}
	}

	return sum
}

// rayDirection is direction ray is travelling towards (right = left to right)
func traverseRay(board [][]tile, ray struct{ x, y int }, rayDirection string, visited map[Key]int) {
	rowLength, colLength := len(board), len(board[0])

	for {
		if ray.x >= colLength || ray.x < 0 || ray.y >= rowLength || ray.y < 0 {
			return
		}

		key := Key{x: ray.x, y: ray.y, direction: rayDirection}
		_, ok := visited[key]
		if ok {
			return
		} else {
			visited[key] = 1
		}

		board[ray.y][ray.x].energized = true
		var charTile = board[ray.y][ray.x]

		// fmt.Println("ray", ray.x, ray.y, rayDirection)
		// writeBoard(board, "day16/output.txt", true)

		switch charTile.char {
		case "/":
			if rayDirection == rightToLeft {
				rayDirection = upToDown
				ray.y += 1
			} else if rayDirection == upToDown {
				rayDirection = rightToLeft
				ray.x -= 1
			} else if rayDirection == leftToRight {
				rayDirection = downToUp
				ray.y -= 1
			} else if rayDirection == downToUp {
				rayDirection = leftToRight
				ray.x += 1
			}
		case "\\":
			if rayDirection == rightToLeft {
				rayDirection = downToUp
				ray.y -= 1
			} else if rayDirection == upToDown {
				rayDirection = leftToRight
				ray.x += 1
			} else if rayDirection == leftToRight {
				rayDirection = upToDown
				ray.y += 1
			} else if rayDirection == downToUp {
				rayDirection = rightToLeft
				ray.x -= 1
			}
		case "|":
			if rayDirection == leftToRight || rayDirection == rightToLeft {
				traverseRay(board, ray, downToUp, visited)
				traverseRay(board, ray, upToDown, visited)
				return
			} else if rayDirection == upToDown {
				ray.y += 1
			} else if rayDirection == downToUp {
				ray.y -= 1
			}
		case "-":
			if rayDirection == leftToRight {
				ray.x += 1
			} else if rayDirection == upToDown || rayDirection == downToUp {
				traverseRay(board, ray, leftToRight, visited)
				traverseRay(board, ray, rightToLeft, visited)
				return
			} else if rayDirection == rightToLeft {
				ray.x -= 1
			}
		case ".":
			if rayDirection == leftToRight {
				ray.x += 1
			} else if rayDirection == rightToLeft {
				ray.x -= 1
			} else if rayDirection == upToDown {
				ray.y += 1
			} else if rayDirection == downToUp {
				ray.y -= 1
			}
		}
	}
}

func printBoardWithRay(board [][]tile, rayPosition struct{ x, y int }, rayDirection string) {
	var boardCopy [][]tile
	for _, row := range board {
		var rowCopy []tile
		rowCopy = append(rowCopy, row...)
		boardCopy = append(boardCopy, rowCopy)
	}

	var rayChar string
	switch rayDirection {
	case leftToRight:
		rayChar = ">"
	case rightToLeft:
		rayChar = "<"
	case upToDown:
		rayChar = "v"
	case downToUp:
		rayChar = "^"
	}

	boardCopy[rayPosition.y][rayPosition.x].char = rayChar
	// printBoard(boardCopy)
	// writeBoard(board, "day16/output.txt", true)
}

func printBoard(board [][]tile) {
	for _, row := range board {
		for _, boardTile := range row {
			fmt.Print(boardTile.char)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printEnergizedBoard(board [][]tile) {
	for _, row := range board {
		for _, boardTile := range row {
			if boardTile.energized {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func writeBoard(board [][]tile, filename string, append bool) {
	var f *os.File
	if append {
		f, _ = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		f, _ = os.Create(filename)
	}
	defer f.Close()

	for _, row := range board {
		for _, boardTile := range row {
			if boardTile.energized {
				f.WriteString("#")
			} else {
				f.WriteString(".")
			}
			// f.WriteString(boardTile.char)
		}
		f.WriteString("\n")
	}
	if append {
		f.WriteString("\n")
	}
}
