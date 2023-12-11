package day11

import (
	"fmt"
	"os"
	"strings"
)

func Run2() {
	var lines []string = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	data, _ := os.ReadFile("day11/input.txt")
	lines = strings.Split(string(data), "\n")

	// between each point
	// find num of empty rows & cols between them, then do math to get the dist for these 2 points

	// so first step is to get record all empty rows and cols

	var galaxies []Point
	var emptyRows, emptyCols []int

	for y, line := range lines {
		var row = []rune(line)
		var emptyRow bool = true
		for x, runeVal := range row {
			if byte(runeVal) == '#' {
				galaxies = append(galaxies, Point{x, y})
				emptyRow = false
			}
		}

		if emptyRow {
			emptyRows = append(emptyRows, y)
		}
	}

	for i := 0; i < len(lines); i++ {
		var emptyCol bool = true
		for _, line := range lines {
			var row = []rune(line)
			if byte(row[i]) == '#' {
				emptyCol = false
			}
		}

		if emptyCol {
			emptyCols = append(emptyCols, i)
		}
	}

	var res int
	var emptyLineSize int = 1000000
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {

			var leftmostGalaxy, rightmostGalaxy Point = galaxies[i], galaxies[j]
			if galaxies[j].x < galaxies[i].x {
				leftmostGalaxy, rightmostGalaxy = galaxies[j], galaxies[i]
			}

			var numEmptyCols int
			for _, col := range emptyCols {
				if col < rightmostGalaxy.x && col > leftmostGalaxy.x {
					numEmptyCols += 1
				}
			}

			var higherGalaxy, lowerGalaxy Point = galaxies[i], galaxies[j]
			if galaxies[j].y > galaxies[i].y {
				higherGalaxy, lowerGalaxy = galaxies[j], galaxies[i]
			}

			var numEmptyRows int
			for _, row := range emptyRows {
				if row < higherGalaxy.y && row > lowerGalaxy.y {
					numEmptyRows += 1
				}
			}

			verticalDist := abs(galaxies[j].y-galaxies[i].y) - numEmptyRows + numEmptyRows*emptyLineSize
			horizontalDist := abs(galaxies[j].x-galaxies[i].x) - numEmptyCols + numEmptyCols*emptyLineSize
			dist := verticalDist + horizontalDist
			// fmt.Println(dist, leftmostGalaxy, rightmostGalaxy, numEmptyCols, numEmptyRows)
			res += dist
		}
	}

	fmt.Println(emptyCols, emptyRows, res)
}
