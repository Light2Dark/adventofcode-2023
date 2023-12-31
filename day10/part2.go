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

func Run2() {
	data, _ := os.ReadFile("day10/input.txt")
	lines := strings.Split(string(data), "\n")

	startPosition := findPoint(&lines)
	fmt.Println(startPosition)
	// visited := map[location]int{startPosition:0}
	// notChecked := []location{startPosition}
}

/*
	Could not find a good way to solve Part2, will look into it in the future (mybe)
	But the below is a solution from another Redditor
	https://github.com/rumkugel13/AdventOfCode2023/day10.go
*/

func findPoint(board *[]string) location {
	for i, line := range *board {
		var row = []rune(line)
		for j, col := range row {
			if byte(col) == 'S' {
				return location{i,j}
			}
		}
	}
	return location{0,0}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Point struct {
	x, y int
}
func Run3() {
	grid := getLines("day10/input.txt")

	startingPoint := findStart(grid)
	visited := map[Point]int{startingPoint:0}
	notChecked := []Point{startingPoint}

	maxDist := 0
	for len(notChecked) > 0 {
		current := notChecked[0]
		notChecked = notChecked[1:]
		next := nextPoints(grid, current)
		for _,point := range next {
			if _,found := visited[point]; !found {
				visited[point] = visited[current] + 1
				maxDist = max(maxDist, visited[current] + 1)
				notChecked = append(notChecked, point)
			}
		}
	}
	
	var result = maxDist
	fmt.Println("Day 10 Part 1 Result: ", result)

	countInside := 0
	for y, row := range grid {
		for x := range row {
			if isInside(grid, Point{x, y}, visited) {
				countInside++
			}
		}
	}

	var result2 = countInside
	fmt.Println("Day 10 Part 2 Result: ", result2)
}

func isInside(grid []string, p Point, theLoop map[Point]int) bool {
	if _,part := theLoop[p]; part {
		return false
	}
	count := 0
	cornerCounts := map[byte]int{}
	for y := p.y + 1; y < len(grid); y++ {
		check := Point{p.x, y}
		tile := grid[y][p.x]
		if tile == 'S' {
			tile = findStartTile(Point{p.x, y}, grid)
		}
		if _,part := theLoop[check]; part {
			if (tile == '-') {
				count++
			} else if tile != '|' && tile != '.' {
				cornerCounts[tile]++
			}
		}
	}

	count += max(cornerCounts['L'], cornerCounts['7']) - abs(cornerCounts['L'] - cornerCounts['7'])
	count += max(cornerCounts['F'], cornerCounts['J']) - abs(cornerCounts['F'] - cornerCounts['J'])
	return count % 2 == 1
}

func findStart(grid []string) Point {
	for y, row := range grid {
		for x, col := range row {
			if byte(col) == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{}
}

func findStartTile(start Point, grid []string) byte {
	points := nextPoints(grid, start)
	minx, maxx, miny, maxy := min(points[0].x, points[1].x), max(points[0].x, points[1].x), min(points[0].y, points[1].y), max(points[0].y, points[1].y)
	if points[0].x == points[1].x {
		return '|'
	} else if points[0].y == points[1].y {
		return '-'
	} else if minx < start.x && miny < start.y {
		return 'J'
	} else if maxx > start.x && maxy > start.y {
		return 'F'
	} else if maxx > start.x && miny < start.y {
		return 'L'
	} else if minx < start.x && maxy > start.y {
		return '7'
	}
	return '.'
}

func nextPoints(grid []string, p Point) []Point {
	points := []Point{}
	switch grid[p.y][p.x] {
	case '|':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x, p.y - 1})
	case '-':
		points = append(points, Point{p.x + 1, p.y})
		points = append(points, Point{p.x - 1, p.y})
	case 'L':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x + 1, p.y})
	case 'J':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x - 1, p.y})
	case '7':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x - 1, p.y})
	case 'F':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x + 1, p.y})
	case '.':
	case 'S':
		down, right, up, left := grid[p.y+1][p.x], grid[p.y][p.x+1], grid[p.y-1][p.x], grid[p.y][p.x-1]
		if down == '|' || down == 'L' || down == 'J' {
			points = append(points, Point{p.x, p.y + 1})
		}
		if right == '-' || right == '7' || right == 'J' {
			points = append(points, Point{p.x + 1, p.y})
		}
		if up == '|' || up == '7' || up == 'F' {
			points = append(points, Point{p.x, p.y - 1})
		}
		if left == '-' || left == 'L' || left == 'F' {
			points = append(points, Point{p.x - 1, p.y})
		}
	}
	return points
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}