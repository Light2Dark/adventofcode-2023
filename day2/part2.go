package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run2() {
	data, err := os.ReadFile("day2/input.txt")
	checkError(err)

	var lines []string = strings.Split(string(data), "\n")
	var sum int

	for _, game := range lines {
		var gameStrings []string = strings.Split(game, ":")

		// since each bag must have at least 1 colour of each (assumption)
		var colourToAmounts = map[string]int{
			"green": 1,
			"blue": 1,
			"red": 1,
		}

		var allReveals = strings.Split(gameStrings[1], ";")

		for _, reveal := range allReveals {
			var cubes = strings.Split(reveal, ",") // [4 blue, 2 green]
			for _, cube := range cubes {
				vals := strings.Split(strings.TrimSpace(cube), " ")
				cubeAmount, err := strconv.Atoi(vals[0])
				colour := vals[1]
				checkError(err)

				// if unset, set it to an initial value
				if colourToAmounts[colour] == 1 {
					colourToAmounts[colour] = cubeAmount
				} else {
					// if set, check if value is more than what has been set
					if cubeAmount > colourToAmounts[colour] {
						colourToAmounts[colour] = cubeAmount
					}
				}
			}
		}

		var product int = 1
		for _,v := range colourToAmounts {
			product *= v
		}
		sum += product
	}

	fmt.Println(sum)
}
