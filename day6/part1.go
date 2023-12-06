package day6

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	data, _ := os.ReadFile("day6/input.txt")
	lines := strings.Split(string(data), "\n")

	re := regexp.MustCompile(`\d+`)
	times := re.FindAllStringSubmatch(lines[0], -1)
	distanceRecords := re.FindAllStringSubmatch(lines[1], -1)

	var numWaysWin []int

	for i := range times {
		time, _ := strconv.Atoi(times[i][0])
		distance, _ := strconv.Atoi(distanceRecords[i][0])
		var wins int

		for m := 1; m < time; m++ {
			travelTime := time - m
			distTravelled := travelTime * m

			if distTravelled > distance {
				wins += 1
			}
		}

		numWaysWin = append(numWaysWin, wins)
	}

	fmt.Println(numWaysWin)

	var prod int = 1
	for _, val := range numWaysWin {
		prod = prod * val
	}
	fmt.Println(prod)
}
