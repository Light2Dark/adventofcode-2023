package day10

import (
	"fmt"
	"os"
	"strings"
)

/*
	Traverse the pipe while keeping count
	Answer is count / 2 or count // 2 (depend on even or odd)
*/

type location struct {
	row int
	col int
}

const (
	Down  = "down"
	Up    = "up"
	Left  = "left"
	Right = "right"
)

func Run() {
	data, _ := os.ReadFile("day10/input.txt")
	var board [][]rune
	var startLoc location = location{
		row: 0,
		col: 0,
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		var row []rune = []rune(line)
		board = append(board, row)

		for j, runeVal := range row {
			if string(runeVal) == "S" {
				startLoc.row = i
				startLoc.col = j
			}
		}
	}

	// board = [][]rune{
	// 	{'7', '-', 'F', '7', '-'},
	// 	{'.', 'F', 'J', '|', '7'},
	// 	{'S', 'J', 'L', 'L', '7'},
	// 	{'|', 'F', '-', '-', 'J'},
	// 	{'L', 'J', '.', 'L', 'J'},
	// }
	// startLoc.row, startLoc.col = 2, 0
	// board = [][]rune{
	// 	{'-', 'L', '|', 'F', '7'},
	// 	{'7', 'S', '-', '7', '|'},
	// 	{'L', '|', '7', '|', '|'},
	// 	{'-', 'L', '-', 'J', 'I'},
	// 	{'L', '|', '-', 'J', 'F'},
	// }
	// startLoc.row, startLoc.col = 1, 1

	// padding the board
	var emptyRow []rune = []rune(strings.Repeat(".", len(board[0])))
	board = append(board, emptyRow)
	board = append([][]rune{emptyRow}, board...)

	for i := range board {
		board[i] = append([]rune("."), board[i]...)
		board[i] = append(board[i], rune('.'))
	}
	startLoc.row, startLoc.col = startLoc.row+1, startLoc.col+1

	var count int
	// var prevNode string
	var cameFromDirection string
	for {
		row, col := startLoc.row, startLoc.col
		node := string(board[row][col])

		var right, left, down, up string = string(board[row][col+1]), string(board[row][col-1]), string(board[row+1][col]), string(board[row-1][col])

		if node == "S" {
			if count > 0 {
				fmt.Println("end found!", count/2+1)
				break
			}
			count++

			if up == "|" || up == "7" || up == "F" {
				if up == "|" {
					startLoc.row = row - 2
					cameFromDirection = Down
				} else if up == "7" {
					startLoc.row, startLoc.col = row-1, col-1
					cameFromDirection = Right
				} else if up == "F" {
					startLoc.row, startLoc.col = row-1, col+1
					cameFromDirection = Left
				}
				continue
			}

			if down == "|" || down == "L" || down == "J" {
				if down == "|" {
					startLoc.row += 2
					cameFromDirection = Up
				} else if down == "L" {
					startLoc.row, startLoc.col = row+1, col+1
					cameFromDirection = Left
				} else if down == "J" {
					startLoc.row, startLoc.col = row+1, col-1
					cameFromDirection = Right
				}
				continue
			}

			if left == "F" || left == "L" || left == "-" {
				if left == "-" {
					startLoc.col = col - 2
					cameFromDirection = Right
				} else if left == "F" {
					startLoc.col, startLoc.row = col-1, row+1
					cameFromDirection = Up
				} else if left == "L" {
					startLoc.col, startLoc.row = col-1, row-1
					cameFromDirection = Down
				}
				continue
			}

			if right == "J" || right == "7" || right == "-" {
				if right == "-" {
					startLoc.col = col + 2
					cameFromDirection = Left
				} else if right == "J" {
					startLoc.col, startLoc.row = col+1, row-1
					cameFromDirection = Down
				} else if right == "7" {
					startLoc.col, startLoc.row = col+1, row+1
					cameFromDirection = Up
				}
				continue
			}
		}

		fmt.Println(node, row, col, cameFromDirection)

		// all other combinations wont be possible because there is only 2 routes at each connected pipe (no divergents)
		if node == "F" {
			if cameFromDirection == Down {
				startLoc.col += 1
				cameFromDirection = Left
			} else if cameFromDirection == Right {
				startLoc.row += 1
				cameFromDirection = Up
			}
		} else if node == "7" {
			if cameFromDirection == Down {
				startLoc.col -= 1
				cameFromDirection = Right
			} else if cameFromDirection == Left {
				startLoc.row += 1
				cameFromDirection = Up
			}
		} else if node == "J" {
			if cameFromDirection == Up {
				startLoc.col -= 1
				cameFromDirection = Right
			} else if cameFromDirection == Left {
				startLoc.row -= 1
				cameFromDirection = Down
			}
		} else if node == "L" {
			if cameFromDirection == Up {
				startLoc.col += 1
				cameFromDirection = Left
			} else if cameFromDirection == Right {
				startLoc.row -= 1
				cameFromDirection = Down
			}
		} else if node == "-" { // direction doesn't change
			if cameFromDirection == Left {
				startLoc.col += 1
			} else if cameFromDirection == Right {
				startLoc.col -= 1
			}
		} else if node == "|" {
			if cameFromDirection == Down {
				startLoc.row -= 1
			} else if cameFromDirection == Up {
				startLoc.row += 1
			}
		}

		count++

		// if count%5 == 0 {
		// 	break
		// }
	}
}
