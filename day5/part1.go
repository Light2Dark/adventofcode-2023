package day5

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Light2Dark/adventofcode-2023/utils"
)

func Run() {
	data, err := os.ReadFile("day5/input.txt")
	utils.CheckError(err)

	lines := strings.Split(string(data), "\n")

	linesSeedToSoil := lines[3:23]
	linesSoilToFertilizer := lines[25:72]
	linesFertilizerToWater := lines[74:107]
	linesWaterToLight := lines[109:149]
	linesLightToTemp := lines[151:179]
	linesTempToHumidity := lines[181:228]
	linesHumidityToLoc := lines[230:258]
	seedLine := strings.Split(lines[0], ":")[1]

	re := regexp.MustCompile(`\d+`)
	seeds := re.FindAllStringSubmatch(seedLine, -1)

	var destinationSeeds []int

	for i := range seeds {
		seed := validateNum(seeds[i][0])
		var soil = findMappedNum(linesSeedToSoil, seed)
		var fertilizer = findMappedNum(linesSoilToFertilizer, soil)
		var water = findMappedNum(linesFertilizerToWater, fertilizer)
		var light = findMappedNum(linesWaterToLight, water)
		var temp = findMappedNum(linesLightToTemp, light)
		var humidity = findMappedNum(linesTempToHumidity, temp)
		var location = findMappedNum(linesHumidityToLoc, humidity)

		destinationSeeds = append(destinationSeeds, location)
	}

	var minLocation int
	fmt.Println(destinationSeeds)
	for _, seed := range destinationSeeds {
		if minLocation == 0 {
			minLocation = seed
		} else if seed < minLocation {
			minLocation = seed
		}
	}
	fmt.Println(minLocation)
}

func findMappedNum(lines []string, num int) int {
	for _, line := range lines {
		nums := strings.Split(line, " ")
		var destNum, sourceNum, numRange = validateNum(nums[0]), validateNum(nums[1]), validateNum(nums[2])
		correctDestNum, ok := getDestNum(sourceNum, destNum, numRange, num)

		if !ok {
			continue
		} else {
			return correctDestNum
		}
	}

	return num // default
}

func validateNum(numString string) int {
	num, err := strconv.Atoi(numString)
	utils.CheckError(err)
	return num
}

func getDestNum(sourceNum, destNum, numRange, num int) (int, bool) {
	if (num >= sourceNum) && (num < sourceNum + numRange) { // means mapping exists
		var diff int = num - sourceNum
		return (destNum + diff), true
	} else {
		return 0, false
	}
}