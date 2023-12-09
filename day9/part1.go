package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
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

		res = append(res, findNext(nums))
	}

	var result int
	for _, val := range res {
		result += val
	}

	fmt.Println(result)
}

func findNext(integers []int) int {
	var allZeros bool = true
	var result []int = []int{}
	for i := 0; i < len(integers)-1; i++ { // we don't check last elem (not sure if it's okayy)
		if integers[i] != 0 {
			allZeros = false
		}
		result = append(result, integers[i+1]-integers[i])
	}

	// base case
	if allZeros {
		return integers[len(integers)-1]
	}

	return integers[len(integers)-1] + findNext(result)
}
