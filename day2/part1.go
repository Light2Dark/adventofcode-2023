package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	data, err := os.ReadFile("day2/input.txt")
	checkError(err)

	var lines []string = strings.Split(string(data), "\n")
	var possibleIDs []int

	for _, game := range lines {
		var gameStrings []string = strings.Split(game, ":")
		var stringID string = gameStrings[0]
		var id, err = strconv.Atoi(strings.Split(stringID, " ")[1])
		checkError(err)

		var allReveals = strings.Split(gameStrings[1], ";")
		var possible bool = true

		for _, reveal := range allReveals {
			var cubes = strings.Split(reveal, ",") // [4 blue, 2 green]
			for _, cube := range cubes {
				vals := strings.Split(strings.TrimSpace(cube), " ")
				cubeAmount, err := strconv.Atoi(vals[0])
				colour := vals[1]
				checkError(err)

				if (colour == "blue" && cubeAmount > 14) || (colour == "green" && cubeAmount > 13) || (colour == "red" && cubeAmount > 12) {
					possible = false
					break
				}
			}

			if !possible {
				break
			} 
		}

		if !possible {
			continue
		} else {
			possibleIDs = append(possibleIDs, id)
		}
	}

	var sum int
	for _, val := range possibleIDs {
		sum += val
	}

	fmt.Println(sum)
}
