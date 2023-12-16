package day15

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	data, _ := os.ReadFile("day15/input.txt")
	line := string(data)

	// line = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	sequences := strings.Split(line, ",")

	var sum int
	for _, seq := range sequences {
		var currentVal int
		for _, runeVal := range seq {
			currentVal = (currentVal + int(runeVal)) * 17 % 256
		}

		sum += currentVal
	}

	fmt.Println(sum)
}
