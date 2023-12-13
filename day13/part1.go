package day13

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	data, _ := os.ReadFile("day13/input.txt")
	blocks := strings.Split(string(data), "\n\n")

// 	pattern1 := `#.##..##.
// ..#.##.#.
// ##......#
// ##......#
// ..#.##.#.
// ..##..##.
// #.#.##.#.`

// 	pattern2 := `#...##..#
// #....#..#
// ..##..###
// #####.##.
// #####.##.
// ..##..###
// #....#..#`

// 	blocks = []string{pattern1, pattern2}

	// fmt.Println(blocks[0], "\n")
	var sum int

	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		var rowReflection bool

		for i := 1; i < len(lines); i++ {
			if hasReflection(lines[0:i], lines[i:]) {
				rowReflection = true
				sum += i * 100 // num of rows above point of reflection is i
				break
			}
		}

		if rowReflection {
			continue
		}

		// column reflection
		for i := 1; i < len(lines[0]); i++ {
			lines1, lines2 := getLines(lines, i)
			if hasReflection(lines1, lines2) {
				sum += i
				break
			}
		}
	}

	fmt.Println(sum)
}

func getLines(lines []string, colStop int) ([]string, []string) {
	var lines1, lines2 []string
	colLength := len([]rune(lines[0]))

	for col := 0; col < colStop; col++ {
		var newLine string
		for row := 0; row < len(lines); row++ {
			newLine += string(lines[row][col])
		}
		lines1 = append(lines1, newLine)
	}

	for col := colStop; col < colLength; col++ {
		var newLine string
		for row := 0; row < len(lines); row++ {
			newLine += string(lines[row][col])
		}
		lines2 = append(lines2, newLine)
	}

	return lines1, lines2
}

// return true if lines1 and lines2 are a reflection
// row reflection
func hasReflection(lines1 []string, lines2 []string) bool {
	minLength := len(lines1)
	if len(lines2) < len(lines1) {
		minLength = len(lines2)
	}

	// get last minLength elements from lines1
	lines1 = lines1[len(lines1)-minLength:]
	lines1 = reverse(lines1)

	// get first minLength elements from lines2
	lines2 = lines2[:minLength]

	var allEqual bool = true
	for i := 0; i < minLength; i++ {
		equal := lines1[i] == lines2[i]
		if !equal {
			allEqual = false
		}
	}

	// fmt.Println(lines1, lines2, minLength)
	return allEqual
}

func reverse(slice []string) []string {
	var newSlice []string
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}

	return newSlice
}
