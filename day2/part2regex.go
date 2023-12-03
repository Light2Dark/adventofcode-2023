package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Regex() {
	data, err := os.ReadFile("day2/input.txt")
	checkError(err)

	lines := strings.Split(string(data), "\n")
	
	var sum int

	for _, line := range lines {
		blueNumbers := extractNumbers(line, "blue")
		redNumbers := extractNumbers(line, "red")
		greenNumbers := extractNumbers(line, "green")

		maxRed := findMax(redNumbers)
		maxGreen := findMax(greenNumbers)
		maxBlue := findMax(blueNumbers)

		sum += maxBlue * maxGreen * maxRed
	}

	fmt.Println(sum)
}

func extractNumbers(line string, colour string) []int {
	re := regexp.MustCompile(fmt.Sprintf(`(\d+)\s%s`, colour))
	matches := re.FindAllStringSubmatch(line, -1) // 

	var numbers []int
	
	for _, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func findMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}