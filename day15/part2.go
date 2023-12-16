package day15

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type Lens struct {
	label       string
	focalLength int
}

func Run2() {
	data, _ := os.ReadFile("day15/input.txt")
	line := string(data)

	// line = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	sequences := strings.Split(line, ",")

	var hashmap = make(map[int][]Lens)

	for _, seq := range sequences {
		var re = regexp.MustCompile("[a-zA-Z]+")
		var letters = re.FindString(seq)
		var boxNum = hash(letters)

		var lastChar rune = rune(seq[len(seq)-1])

		// equals symbol, lastChar is the focal length
		if unicode.IsNumber(rune(lastChar)) {
			var focalLength = int(lastChar - '0')
			var overwrite bool
			for i, lensInBox := range hashmap[boxNum] {
				if lensInBox.label == letters {
					hashmap[boxNum][i].focalLength = focalLength // if lensInBox.focalLength, that is just a copy of the struct as struct is passed by value
					overwrite = true
					break
				}
			}

			if !overwrite {
				var lens = Lens{label: letters, focalLength: focalLength}
				hashmap[boxNum] = append(hashmap[boxNum], lens)
			}

		} else { // -
			lenses, ok := hashmap[boxNum]
			if !ok {
				continue
			}

			for i := range lenses {
				if lenses[i].label == letters {
					newLenses := deleteAtIndex(lenses, i)
					hashmap[boxNum] = newLenses
					break
				}
			}
		}
	}

	fmt.Println(hashmap)

	var sum int
	for boxNum, lenses := range hashmap {
		var res int
		for i := range lenses {
			res += (boxNum+1) * (i+1) * lenses[i].focalLength
		}
		sum += res
	}

	fmt.Println(sum)
}

func deleteAtIndex(slice []Lens, index int) []Lens {
	return append(slice[:index], slice[index+1:]...)
}

func hash(chars string) int {
	var currentVal int
	for _, runeVal := range chars {
		currentVal = (currentVal + int(runeVal)) * 17 % 256
	}

	return currentVal
}
