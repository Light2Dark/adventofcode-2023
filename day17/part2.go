package day17

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func Run2() {
// 	data := `2413432311323
// 3215453535623
// 3255245654254
// 3446585845452
// 4546657867536
// 1438598798454
// 4457876987766
// 3637877979653
// 4654967986887
// 4564679986453
// 1224686865563
// 2546548887735
// 4322674655533`

	data, _ := os.ReadFile("day17/input.txt")
	lines := strings.Split(string(data), "\n")

	var board = [][]int{}
	for _, line := range lines {
		var row = []int{}
		for _, runeVal := range line {
			var integer = int(runeVal - '0')
			row = append(row, integer)
		}
		board = append(board, row)
	}

	const paddingInt = 10_000

	// padding with rows
	var newRow []int
	for i := 0; i < len(board[0]); i++ {
		newRow = append(newRow, paddingInt)
	}
	board = append(board, newRow)
	board = append([][]int{newRow}, board...)

	// padding with columns
	for i := range board {
		board[i] = append([]int{paddingInt}, board[i]...)
		board[i] = append(board[i], paddingInt)
	}

	// printBoard(board)

	var pos1 Position = Position{
		x: 1, y: 1, direction: right, stepNum: 0,
		heatVal: 0, // first block not counted
		index:   0,
	}

	var pos2 = pos1
	pos2.index = 1
	pos2.direction = down

	type Key struct {
		x         int
		y         int
		direction string
		stepNum   int
	}
	var visited = map[Key]bool{}

	var pq PriorityQueue = make(PriorityQueue, 2)
	pq[0] = &pos1
	pq[1] = &pos2
	heap.Init(&pq)

	// heap.Push(&pq, &pos1)
	var i int = 2
	for pq.Len() > 0 {
		position := heap.Pop(&pq).(*Position) // pop topmost node
		x, y, direction, stepNum, heatVal := position.x, position.y, position.direction, position.stepNum, position.heatVal
		// fmt.Printf("Position: {%v, %v}, direction: %v, stepNumber: %v, heatValue: %v\n", x, y, direction, stepNum, heatVal)

		if x < 1 || x >= len(board[0])-1 || y < 1 || y >= len(board)-1 {
			continue
		}

		var key = Key{x, y, direction, stepNum}
		_, found := visited[key]
		if found {
			// fmt.Println("cache hit")
			continue
		}

		if x == len(board[0])-2 && y == len(board)-2 && stepNum >= 4 {
			fmt.Println("found", heatVal)
			break
		}

		var newPos *Position

		if direction == right || direction == left {
			if stepNum >= 4 { // go up or down only after 4 steps
				newPos = &Position{x: x, y: y - 1, stepNum: 1, direction: up, heatVal: heatVal + board[y-1][x], index: i}
				heap.Push(&pq, newPos)

				newPos = &Position{x: x, y: y + 1, stepNum: 1, direction: down, heatVal: heatVal + board[y+1][x], index: i}
				heap.Push(&pq, newPos)
			}

			if stepNum != 10 {
				if direction == right { // continue right
					newPos = &Position{x: x + 1, y: y, stepNum: stepNum + 1, direction: right, heatVal: heatVal + board[y][x+1], index: i}
					heap.Push(&pq, newPos)
				} else if direction == left { // continue left
					newPos = &Position{x: x - 1, y: y, stepNum: stepNum + 1, direction: left, heatVal: heatVal + board[y][x-1], index: i}
					heap.Push(&pq, newPos)
				}
			}
		}

		if direction == up || direction == down { 
			if stepNum >= 4 { // go left or right only after 4 steps
				newPos := &Position{x: x - 1, y: y, stepNum: 1, direction: left, heatVal: heatVal + board[y][x-1], index: i}
				heap.Push(&pq, newPos)
	
				newPos = &Position{x: x + 1, y: y, stepNum: 1, direction: right, heatVal: heatVal + board[y][x+1], index: i}
				heap.Push(&pq, newPos)
			}

			if stepNum != 10 {
				if direction == up { // continue up
					newPos := &Position{x: x, y: y - 1, stepNum: stepNum + 1, direction: up, heatVal: heatVal + board[y-1][x], index: i}
					heap.Push(&pq, newPos)
				} else if direction == down { // continue down
					newPos := &Position{x: x, y: y + 1, stepNum: stepNum + 1, direction: down, heatVal: heatVal + board[y+1][x], index: i}
					heap.Push(&pq, newPos)
				}
			}
		}

		visited[key] = true
		i++
	}
}
