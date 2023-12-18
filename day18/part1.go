package day18

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func Run() {
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
		direction, numStr := parts[0], parts[1]
		num, _ := strconv.Atoi(numStr)

		if direction == "R" {
			current.x += num
		} else if direction == "U" {
			current.y -= num
		} else if direction == "D" {
			current.y += num
		} else if direction == "L" {
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

// shoelace formula
func getArea(vertices []Position) int {

	var sum1, sum2 int
	for i := 0; i < len(vertices); i++ {
		nextI := (i + 1) % len(vertices)
		sum1 += vertices[i].x * vertices[nextI].y
		sum2 += vertices[i].y * vertices[nextI].x
	}

	return int(math.Abs(float64(sum1-sum2))) / 2
}

func getPerimeter(vertices []Position) int {
	var sum int
	for i := 0; i < len(vertices); i++ {
		nextI := (i + 1) % len(vertices)

		horizontal := math.Pow(math.Abs(float64(vertices[i].x-vertices[nextI].x)), 2)
		vertical := math.Pow(math.Abs(float64(vertices[i].y-vertices[nextI].y)), 2)
		sum += int(math.Sqrt(vertical + horizontal))
	}
	return sum
}

// pick's theorem: A = i + b/2 - 1, where i is interior points and b is boundary points (perimeter). A is derived from shoelace formula
// i = A - b/2 + 1
func getInteriorPoints(area int, boundaryPoints int) int {
	return area - (boundaryPoints)/2 + 1
}
