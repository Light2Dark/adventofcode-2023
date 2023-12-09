package day8

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Run() {
	data, _ := os.ReadFile("day8/input.txt")
	lines := strings.Split(string(data), "\n")

	var maps  = map[string][2]string{}

	for _, line := range lines[2:] {
		re := regexp.MustCompile(`\w+`)
		matches := re.FindAllStringSubmatch(line, -1)

		var node string = matches[0][0]
		var leftRoute, rightRoute string = matches[1][0], matches[2][0]
		
		maps[node] = [2]string{leftRoute, rightRoute}
	}

	instructions := lines[0]
	for i := 0; i < 10; i++ {
		instructions += instructions
	}
	instructionsRune := []rune(instructions)
	
	var startNode string = "AAA"
	for i := 0; i < len(instructionsRune); i++ {
		var destNodes [2]string = maps[startNode]
		if string(instructionsRune[i]) == "L" {
			startNode = destNodes[0]
		} else if string(instructionsRune[i]) == "R" {
			startNode = destNodes[1]
		} else {
			fmt.Println(string(instructionsRune[i]))
		}

		if startNode == "ZZZ" {
			fmt.Println(i+1)
			break
		}
	}
}