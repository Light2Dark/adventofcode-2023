package day12

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	cfg  string
	nums string
}

func Run2() {
	var lines []string

	data, _ := os.ReadFile("day12/input.txt")
	lines = strings.Split(string(data), "\n")

	// lines = []string{
	// 	".# 1",
	// 	"???.### 1,1,3",
	// 	".??..??...?##. 1,1,3",
	// 	"?#?#?#?#?#?#?#? 1,3,1,6",
	// 	"????.#...#... 4,1,1",
	// 	"????.######..#####. 1,6,5",
	// 	"?###???????? 3,2,1",
	// }

	var duplicateNum = 4
	var sum int
	var cache = make(map[Key]int)

	for _, line := range lines {
		input := strings.Split(line, " ")
		var springs []rune = []rune(input[0])
		var originalSpring = springs
		for i := 0; i < duplicateNum; i++ {
			springs = []rune(string(springs) + "?" + string(originalSpring))
		}

		numString := strings.Split(input[1], ",")
		var nums []int
		for _, numStr := range numString {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		originalNums := nums
		for i := 0; i < duplicateNum; i++ {
			nums = append(nums, originalNums...)
		}
		 
		sum += countMemo(springs, nums, cache)
	}

	fmt.Println(sum)
}

func countMemo(cfg []rune, nums []int, cache map[Key]int) int {
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

	var key = Key{cfg: string(cfg), nums: intSliceToString(nums)}
	val, ok := cache[key]
	if ok {
		return val
	}
	
	if startChar == "." || startChar == "?" {
		result += countMemo(cfg[1:], nums, cache)
	}

	if startChar == "#" || startChar == "?" {
		if nums[0] <= len(cfg) && charNotIn(".", cfg[:nums[0]]) && (len(cfg) == nums[0] || string(cfg[nums[0]]) != "#") {
			if len(cfg) == nums[0] {
				result += countMemo([]rune{}, nums[1:], cache)
			} else {
				result += countMemo(cfg[nums[0]+1:], nums[1:], cache)
			}
		}
	}

	cache[key] = result
	return result
}


func intSliceToString(nums []int) string {
	var numString []string
	for _, val := range nums {
		numString = append(numString, strconv.Itoa(val))
	}
	return strings.Join(numString, ",")
}

func Run2NotWorking() {
	var lines []string

	data, _ := os.ReadFile("day12/input.txt")
	lines = strings.Split(string(data), "\n")

	var sum int64
	var duplicateNum = 4

	res_1 := getEachCombination(lines, 0)
	res_2 := getEachCombination(lines, 1)

	for i := 0; i < len(res_1); i++ {
		factor := res_2[i] / res_1[i]
		result := res_1[i] * int(math.Pow(float64(factor), float64(duplicateNum)))
		sum += int64(result)
	}

	fmt.Println(sum)
}

func getEachCombination(lines []string, duplicateNum int) []int {
	var combinations []int
	for _, line := range lines {
		input := strings.Split(line, " ")
		var springs []rune = []rune(input[0])
		var originalSpring = springs
		for i := 0; i < duplicateNum; i++ {
			springs = []rune(string(springs) + "?" + string(originalSpring))
		}

		numString := strings.Split(input[1], ",")
		var nums []int
		for _, numStr := range numString {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		originalNums := nums
		for i := 0; i < duplicateNum; i++ {
			nums = append(nums, originalNums...)
		}

		// fmt.Println(string(springs), nums, count(springs, nums))
		combinations = append(combinations, count(springs, nums))
	}

	return combinations
}
