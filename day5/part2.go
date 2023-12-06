package day5

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Light2Dark/adventofcode-2023/utils"
)

func Run2() {
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

	var seedPairs [][2]int
	for i := 0; i < len(seeds); i += 2 {
		seed1 := validateNum(seeds[i][0])
		seed2 := validateNum(seeds[i+1][0])

		pair := [2]int{seed1, seed2}
		seedPairs = append(seedPairs, pair)
	}

	var minLocation int
	var wg = sync.WaitGroup{}
	var mutex = sync.RWMutex{}

	for m, pair := range seedPairs {
		fmt.Printf("Time: %v, pairNum: %v, pair: %v\n", time.Now(), m, pair)
		wg.Add(1)
		go func(pair [2]int) {
			for i := pair[0]; i < pair[0]+pair[1]; i++ {
				seed := i
				var soil = findMappedNum(linesSeedToSoil, seed)
				var fertilizer = findMappedNum(linesSoilToFertilizer, soil)
				var water = findMappedNum(linesFertilizerToWater, fertilizer)
				var light = findMappedNum(linesWaterToLight, water)
				var temp = findMappedNum(linesLightToTemp, light)
				var humidity = findMappedNum(linesTempToHumidity, temp)
				var location = findMappedNum(linesHumidityToLoc, humidity)

				mutex.RLock()
				if location < minLocation || minLocation == 0 {
					mutex.RUnlock()
					mutex.Lock()
					minLocation = location
					fmt.Println(minLocation)
					mutex.Unlock()
				} else {
					mutex.RUnlock()
				}
			}
			wg.Done()
			fmt.Println(pair)
		}(pair)
	}

	wg.Wait()
	fmt.Println(minLocation)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %v time\n", name, elapsed)
}
