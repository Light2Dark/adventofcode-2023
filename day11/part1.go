package day11

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func Run() {

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

	// add more rows
	var addedRows []string
	for _, line := range lines {
		row := []rune(line)
		var emptyRow bool = true
		for _, runeVal := range row {
			if string(runeVal) != "." {
				emptyRow = false
				break
			}
		}
		if emptyRow {
			addedRows = append(addedRows, strings.Repeat(".", len(row)))
		}
		addedRows = append(addedRows, line)
	}

	// add cols
	var newLines []string = addedRows
	for i := 0; i < len(addedRows[0]); i++ {
		var emptyCol bool = true
		for _, line := range newLines {
			row := []rune(line)
			if byte(row[i]) != '.' {
				emptyCol = false
				break
			}
		}

		if emptyCol {
			for j := 0; j < len(newLines); j++ {
				oldStr := newLines[j]
				newStr := oldStr[:i] + "." + oldStr[i:]
				newLines[j] = newStr
			}
			i++
		}
	}

	lines = newLines
	var galaxies []Point

	for y, line := range lines {
		var row = []rune(line)
		for x, runeVal := range row {
			if byte(runeVal) == '#' {
				galaxies = append(galaxies, Point{x, y})
			}
		}
	}

	var res int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := abs(galaxies[j].y-galaxies[i].y) + abs(galaxies[j].x-galaxies[i].x)
			res += dist
		}
	}

	fmt.Println(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
