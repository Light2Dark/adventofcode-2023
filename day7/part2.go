package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Light2Dark/adventofcode-2023/utils"
)


func Run2() {
	data, err := os.ReadFile("day7/input.txt")
	utils.CheckError(err)

	lines := strings.Split(string(data), "\n")
	var fiveOfAKind, fourOfAKind, fullHouse, threeOfAKind, twoPair, onePair, highCard []Hand // highCard not possible with J wildcard

	// place all cards into their respective group
	for _, line := range lines {
		lineInfo := strings.Split(line, " ")
		bidAmount, _ := strconv.Atoi(lineInfo[1])
		var groupFound bool

		// to iterate through string, use runes as the underlying structure of string is bytes.
		hand := []rune(lineInfo[0])
		var charCount = map[string]int{}

		// first remove all J's
		var newHand []rune
		var missingJnum int
		for _, runeVal := range hand {
			if string(runeVal) != "J" {
				newHand = append(newHand, runeVal)
				continue
			}
			missingJnum += 1
		}

		// count characters as normal
		for _, runeVal := range newHand {
			charCount[string(runeVal)] += 1
		}

		var highestCharacter string
		var highestCount int

		for char, count := range charCount {
			if count > highestCount {
				highestCharacter = char
				highestCount = count
			}
		}

		// add to highest count for missing chars due to removal
		for i := 0; i < missingJnum; i++ {
			charCount[highestCharacter] += 1	
		}

		var cards = Hand{string(hand), bidAmount}

		if len(newHand) == 0 {
			fiveOfAKind = append(fiveOfAKind, cards)
			continue
		}

		for _, v := range charCount {
			if v == 5 {
				fiveOfAKind = append(fiveOfAKind, cards)
				groupFound = true
				break
			} else if v == 4 {
				fourOfAKind = append(fourOfAKind, cards)
				groupFound = true
				break
			} else if (v == 3 || v == 2) && len(charCount) == 2 {
				fullHouse = append(fullHouse, cards)
				groupFound = true
				break
			} else if v == 3 {
				threeOfAKind = append(threeOfAKind, cards)
				groupFound = true
				break
			} else if v == 2 && len(charCount) == 3 {
				twoPair = append(twoPair, cards)
				groupFound = true
				break
			} else if v == 2 {
				onePair = append(onePair, cards)
				groupFound = true
				break
			}
		}

		if !groupFound {
			highCard = append(highCard, cards)
		}
	}

	ordering := "AKQT98765432J"
	fiveOfAKind = sortHand(fiveOfAKind, ordering)
	fourOfAKind = sortHand(fourOfAKind, ordering)
	threeOfAKind = sortHand(threeOfAKind, ordering)
	fullHouse = sortHand(fullHouse, ordering)
	twoPair = sortHand(twoPair, ordering)
	onePair = sortHand(onePair, ordering)
	highCard = sortHand(highCard, ordering)

	var hands []Hand
	hands = append(hands, highCard...)
	hands = append(hands, onePair...)
	hands = append(hands, twoPair...)
	hands = append(hands, threeOfAKind...)
	hands = append(hands, fullHouse...)
	hands = append(hands, fourOfAKind...)
	hands = append(hands, fiveOfAKind...)

	var sum int
	for i := range hands {
		sum += hands[i].bid * (i + 1)
	}

	fmt.Println(sum)

	writeHands("day7/output2.txt", hands)
}