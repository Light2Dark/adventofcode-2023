package day14

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	data, _ := os.ReadFile("day14/input.txt")

	lines := strings.Split(string(data), "\n")

// 	var sample string = `O....#....
// O.OO#....#
// .....##...
// OO.#O....O
// .O.....O#.
// O.#..O.#.#
// ..O..#O..O
// .......O..
// #....###..
// #OO..#....`

// 	lines = strings.Split(sample, "\n")

	var slices []string

	// each col is a slice
	for i := 0; i < len(lines[0]); i++ {
		var slice string
		for row := 0; row < len(lines); row++ {
			line := []rune(lines[row])
			slice += string(line[i])
		}
		slices = append(slices, slice)
	}

	// fmt.Println(slices)

	var sum int
	for _, slice := range slices {
		sum += getStringCount(slice)
		// fmt.Println("slice count",getStringCount(slice))
	}

	fmt.Println(sum)
}

func getStringCount(slice string) int {
	runeString := []rune(slice)
	var sum int

	var count int
	var lastHashPos int = -1
	var defaultLength int = len(runeString)

	for i, runeVal := range runeString {
		runeChar := string(runeVal)
		if runeChar == "O" {
			count++
		}

		if runeChar == "#" {
			sum += addUpCount(lastHashPos, defaultLength, count)
			lastHashPos = i
			count = 0
		}
	}

	if count > 0 {
		sum += addUpCount(lastHashPos, defaultLength, count)
	}

	return sum
}

func addUpCount(lastHashPosition, defaultLength, count int) int {
	var startFrom int = defaultLength - lastHashPosition - 1
	if lastHashPosition == -1 { // no hash before O
		startFrom = defaultLength
	}


	var sum int
	for i := 0; i < count; i++ {
		sum += startFrom
		startFrom--
	}
	return sum
}
