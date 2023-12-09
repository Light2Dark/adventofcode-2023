package day8

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Run2() {
	data, _ := os.ReadFile("day8/input.txt")
	lines := strings.Split(string(data), "\n")

	var maps = map[string][2]string{}
	var startNodes []string

	for _, line := range lines[2:] {
		re := regexp.MustCompile(`\w+`)
		matches := re.FindAllStringSubmatch(line, -1)

		var node string = matches[0][0]
		var leftRoute, rightRoute string = matches[1][0], matches[2][0]

		maps[node] = [2]string{leftRoute, rightRoute}

		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}

	instructions := lines[0]

	fmt.Println(startNodes)
	num1 := findCycleNum([]string{"TVA"}, instructions, maps)
	num2 := findCycleNum([]string{"AAA"}, instructions, maps)
	num3 := findCycleNum([]string{"VBA"}, instructions, maps)
	num4 := findCycleNum([]string{"DVA"}, instructions, maps)
	num5 := findCycleNum([]string{"VPA"}, instructions, maps)
	num6 := findCycleNum([]string{"DTA"}, instructions, maps)

	fmt.Println(lcm(num1, num2, num3, num4, num5, num6))
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func findCycleNum(startNodes []string, instructions string, maps map[string][2]string) int {
	var instructionsRune []rune = []rune(instructions)
	var i, count int
	for i < len(instructionsRune) {
		var instruction string = string(instructionsRune[i])
		var zAll bool = true

		for m := range startNodes {
			var destNodes [2]string = maps[startNodes[m]]
			if instruction == "L" {
				startNodes[m] = destNodes[0]
			} else {
				startNodes[m] = destNodes[1]
			}

			if !strings.HasSuffix(startNodes[m], "Z") {
				zAll = false
			}
		}

		if zAll {
			fmt.Println("Steps taken", count+1, startNodes)
			return count + 1
		}

		if i == len(instructionsRune)-1 {
			i = 0
		} else {
			i++
		}
		count++
	}

	return 0
}
