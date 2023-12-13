package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	var lines []string

	data, _ := os.ReadFile("day12/input.txt")
	lines = strings.Split(string(data), "\n")

	// lines = []string{
	// 	"???.### 1,1,3",
	// 	".??..??...?##. 1,1,3",
	// 	"?#?#?#?#?#?#?#? 1,3,16",
	// }

	var sum int

	for _, line := range lines {
		input := strings.Split(line, " ")
		var springs []rune = []rune(input[0])

		numString := strings.Split(input[1], ",")
		var nums []int
		for _, numStr := range numString {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		fmt.Println(string(springs), nums)
		sum += count(springs, nums)
	}

	fmt.Println(sum)
}

func count(cfg []rune, nums []int) int {
	if string(cfg) == "" {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(nums) == 0 {
		for _, runeVal := range cfg {
			if string(runeVal) == "#" {
				return 0
			}
		}
		return 1
	}

	var result int
	startChar := string(cfg[0])

	if startChar == "." || startChar == "?" {
		result += count(cfg[1:], nums)
	}

	if startChar == "#" || startChar == "?" {
		if nums[0] <= len(cfg) && charNotIn(".", cfg[:nums[0]]) && (len(cfg) == nums[0] || string(cfg[nums[0]]) != "#") {
			if len(cfg) == nums[0] {
				result += count([]rune{}, nums[1:])
			} else {
				result += count(cfg[nums[0]+1:], nums[1:])
			}
		}
	}

	return result
}

// if character is not in runeStr, returns true
func charNotIn(char string, runeStr []rune) bool {
	for _, runeVal := range runeStr {
		if string(runeVal) == char {
			return false
		}
	}
	return true
}
