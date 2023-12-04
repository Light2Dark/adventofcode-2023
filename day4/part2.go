package day4

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Light2Dark/adventofcode-2023/utils"
)

func Run2() {
	data, err := os.ReadFile("day4/input.txt")
	utils.CheckError(err)

	lines := strings.Split(string(data), "\n")

	var numberOfCards = map[int]int{}
	for i := 1; i < len(lines)+1; i++ {
		numberOfCards[i] = 1
	}

	// 0th line is Card 1
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		stringsTicket := strings.Split(line, "|")

		re := regexp.MustCompile(`\d+`)
		ownNumbersMatches := re.FindAllStringSubmatch(stringsTicket[1], -1)
		winningNumbersMatches := re.FindAllStringSubmatch(stringsTicket[0], -1)
		winningNumbersMatches = winningNumbersMatches[1:] // remove card number

		var winningNumbers []int = getIntSlice(winningNumbersMatches)
		var ownNumbers []int = getIntSlice(ownNumbersMatches)

		var matches int
		for _, winningNum := range winningNumbers {
			for _, ownNum := range ownNumbers {
				if ownNum == winningNum {
					matches += 1
				}
			}
		}

		if i == len(lines)-1 { // last card, so doesn't matter
			break
		}

		// add winnings to next cards
		for j := 1; j < matches+1; j++ {
			currNumCards := numberOfCards[i+1]
			oriNumCards := numberOfCards[i+j+1] // next card
			numberOfCards[i+j+1] = currNumCards + oriNumCards
		}
	}

	var totalCards int
	for _, val := range numberOfCards {
		totalCards += val
	}

	fmt.Println(totalCards)
}
