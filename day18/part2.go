package day18

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

func Run2() {
	data, _ := os.ReadFile("day18/input.txt")
	lines := strings.Split(string(data), "\n")

// 	var data string = `R 6 (#70c710)
// D 5 (#0dc571)
// L 2 (#5713f0)
// D 2 (#d2c081)
// R 2 (#59c680)
// D 2 (#411b91)
// L 5 (#8ceee2)
// U 2 (#caa173)
// L 1 (#1b58a2)
// U 2 (#caa171)
// R 2 (#7807d2)
// U 3 (#a77fa3)
// L 2 (#015232)
// U 2 (#7a21e3)`

// 	lines := strings.Split(data, "\n")

	var vertices = []Position{}
	var current = Position{0, 0}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		hexStr := parts[len(parts)-1]
		hexStr = string(hexStr[1:len(hexStr)-1])

		numStr := hexStr[1:len(hexStr)-1]
		direction := string(hexStr[len(hexStr)-1])
		
		num64, _ := strconv.ParseInt(numStr, 16, 0)
		num := int(num64)

		if direction == "0" {
			current.x += num
		} else if direction == "3" {
			current.y -= num
		} else if direction == "1" {
			current.y += num
		} else if direction == "2" {
			current.x -= num
		}

		vertices = append(vertices, current)
	}

	// vertices = []Position{ // area is 21
	// 	{4, 0}, {4, 2}, {2, 2}, {2, 4}, {0, 4}, {0, 0},
	// }

	// vertices = []Position{
	// 	{0, 0}, {5, 0}, {5, 2}, {1, 2}, {1, 1}, {0, 1},
	// }

	area := getArea(vertices)
	perimeter := getPerimeter(vertices)
	interiorPoints := getInteriorPoints(area, perimeter)

	fmt.Println(interiorPoints + perimeter)
}