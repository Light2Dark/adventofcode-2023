package day7

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Light2Dark/adventofcode-2023/utils"
)

type Hand struct {
	cards string
	bid   int
}

func Run() {
	data, err := os.ReadFile("day7/input.txt")
	utils.CheckError(err)

	lines := strings.Split(string(data), "\n")
	var fiveOfAKind, fourOfAKind, fullHouse, threeOfAKind, twoPair, onePair, highCard []Hand

	// place all cards into their respective group
	for _, line := range lines {
		lineInfo := strings.Split(line, " ")
		bidAmount, _ := strconv.Atoi(lineInfo[1])
		var groupFound bool

		// to iterate through string, use runes as the underlying structure of string is bytes.
		hand := []rune(lineInfo[0])
		var charCount = map[string]int{}

		for _, runeVal := range hand {
			charCount[string(runeVal)] += 1
		}

		var cards = Hand{string(hand), bidAmount}
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

	ordering := "AKQJT98765432"
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

	writeHands("day7/output.txt", hands)
}

func sortHand(slice []Hand, ordering string) []Hand {
	slices.SortFunc(slice, func(a, b Hand) int {
		return customComp(a.cards, b.cards, ordering)
	})
	return slice
}

func customComp(a, b, ordering string) int {
	length := len(a) // b length is the same

	aRune, bRune := []rune(a), []rune(b)
	orderingRune := []rune(ordering)

	for i := 0; i < length; i++ {
		indexA := indexInOrdering(aRune[i], orderingRune)
		indexB := indexInOrdering(bRune[i], orderingRune)

		if indexA < indexB {
			return 1
		} else if indexB < indexA {
			return -1
		}
	}
	return 0
}

func indexInOrdering(a rune, orderingRune []rune) int {
	for i := range orderingRune {
		if string(a) == string(orderingRune[i]) {
			return i
		}
	}
	fmt.Println("Not found in ordering!")
	return len(orderingRune) // place at last
}

func writeHands(savepath string, hands []Hand) {
	// Uneccesary but useful for checking
	f, err := os.Create(savepath)
	utils.CheckError(err)
	defer f.Close()

	for _, hand := range hands {
		bidStr := strconv.Itoa(hand.bid)
		f.WriteString(hand.cards + " " + bidStr + "\n")
	}
}
