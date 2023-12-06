package day6

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getTimeDist() (int, int) {
	data, _ := os.ReadFile("day6/input.txt")
	lines := strings.Split(string(data), "\n")

	re := regexp.MustCompile(`\d+`)
	times := re.FindAllStringSubmatch(lines[0], -1)
	distanceRecords := re.FindAllStringSubmatch(lines[1], -1)

	var time int = concatArrStringToInt(times)
	var distance int = concatArrStringToInt(distanceRecords)

	return time, distance
}

func Run2Opt() {
	// there's 2 graphs that intersect.
	// y = -x^2 + nx , where n is a constant (n = race time)
	// y = r, where r is a constant (r = distance record)
	// -x^2 + nx - r = 0, we have to find both intersection points and subtract the difference which will be the answer
	// formula to find roots are (-b +- sqrt(b^2 - 4ac)) / 2a

	r, d := getTimeDist()
	raceTime, distanceRecord := float64(r), float64(d)

	// -x^2 + 58996469x - 478223210191071 = 0

	rootOne := (-(raceTime) + math.Sqrt((raceTime*raceTime)-(4*-1*(-distanceRecord)))) / (2 * -1)
	rootTwo := (-(raceTime) - math.Sqrt((raceTime*raceTime)-(4*-1*(-distanceRecord)))) / (2 * -1)

	fmt.Println(rootOne, rootTwo, int(math.Round(rootTwo - rootOne)))
}

func Run2() {
	time, distance := getTimeDist()
	var wins int

	for m := 1; m < time; m++ {
		travelTime := time - m
		distTravelled := travelTime * m

		if distTravelled > distance {
			wins += 1
		}
	}

	fmt.Println(wins)
}

func concatArrStringToInt(s [][]string) int {
	var str string
	for _, t := range s {
		str += t[0]
	}
	integer, _ := strconv.Atoi(str)
	return integer
}
