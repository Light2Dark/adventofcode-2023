package day4

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Light2Dark/adventofcode-2023/utils"
)

// points is 2^n, where n is num of matches

func Run() {
	data, err := os.ReadFile("day4/input.txt")
	utils.CheckError(err)

	var totalPoints int

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		stringsTicket := strings.Split(line, "|")

		re := regexp.MustCompile(`\d+`)
		ownNumbersMatches := re.FindAllStringSubmatch(stringsTicket[1], -1)
		winningNumbersMatches := re.FindAllStringSubmatch(stringsTicket[0], -1)
		winningNumbersMatches = winningNumbersMatches[1:] // remove card number

		var winningNumbers []int = getIntSlice(winningNumbersMatches)
		var ownNumbers []int = getIntSlice(ownNumbersMatches)

		var count int
		for _, winningNum := range winningNumbers {
			for _, ownNum := range ownNumbers {
				if ownNum == winningNum {
					count += 1
				}
			}
		}

		if count > 0 {
			points := math.Pow(2, float64(count-1))
			totalPoints += int(points)
		}
	}

	fmt.Println(totalPoints)
}

func getIntSlice(matches [][]string) []int {
	var intSlice []int
	for _, match := range matches {
		num, err := strconv.Atoi(match[0])
		utils.CheckError(err)
		intSlice = append(intSlice, num)
	}

	return intSlice
}
