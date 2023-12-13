package day13

import (
	"fmt"
	"os"
	"strings"
)

func Run2() {
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
			if hasReflectionRough(lines[0:i], lines[i:]) {
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
			if hasReflectionRough(lines1, lines2) {
				sum += i
				break
			}
		}
	}

	fmt.Println(sum)
}

// return true if lines1 and lines2 are a reflection with a smudge
func hasReflectionRough(lines1 []string, lines2 []string) bool {
	minLength := len(lines1)
	if len(lines2) < len(lines1) {
		minLength = len(lines2)
	}

	// get last minLength elements from lines1
	lines1 = lines1[len(lines1)-minLength:]
	lines1 = reverse(lines1)

	// get first minLength elements from lines2
	lines2 = lines2[:minLength]

	var mistakeCount int
	for i := 0; i < minLength; i++ {
		count := getMistakeCount(lines1[i], lines2[i])
		mistakeCount += count
	}

	return mistakeCount == 1
}

func getMistakeCount(line1, line2 string) int {
	var mistakeCount int
	line1Rune, line2Rune := []rune(line1), []rune(line2)

	for i := 0; i < len(line1Rune); i++ {
		if line1Rune[i] != line2Rune[i] {
			mistakeCount++
		}
	}

	return mistakeCount
}