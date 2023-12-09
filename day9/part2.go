package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run2() {
	data, _ := os.ReadFile("day9/input.txt")
	lines := strings.Split(string(data), "\n")
	var res = []int{}

	for _, line := range lines {
		var nums = []int{}
		var str = strings.Split(line, " ")
		for _, val := range str {
			num, _ := strconv.Atoi(val)
			nums = append(nums, num)
		}

		res = append(res, findPrevious(nums))
	}

	var result int
	for _, val := range res {
		result += val
	}

	fmt.Println(result)
}

func findPrevious(integers []int) int {
	var allZeros bool = true
	var result []int = []int{}
	for i := 0; i < len(integers)-1; i++ {
		if integers[i] != 0 {
			allZeros = false
		}
		result = append(result, integers[i+1]-integers[i])
	}

	if allZeros {
		return integers[0]
	}

	return integers[0] - findPrevious(result)
}
