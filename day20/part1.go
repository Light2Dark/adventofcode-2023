package day20

import (
	"fmt"
	"regexp"
	"os"
	"strings"
)

const (
	conjunction = "conjunction"
	broadcaster = "broadcaster"
	flipflop    = "flipflop"
)

func Run() {
	data, _ := os.ReadFile("day20/input.txt")
	lines := strings.Split(string(data), "\n")

	moduleMap := getModuleMap(lines)
	
	var high, low int
	for i := 0; i < 1000; i++ {
		highCount, lowCount := pushButton(moduleMap)
		high += highCount
		low += lowCount
	}

	fmt.Println(high * low)
}

func pushButton(moduleMap map[string]Module) (int, int) {
	var firstEvent = Event{pulse: false, moduleTargeted: moduleMap[broadcaster]}
	var queue = &Queue{events: []Event{firstEvent}}
	var highPulseCount, lowPulseCount int

	for queue.length() > 0 {
		event, ok := queue.dequeue()
		if !ok {
			break
		}

		if event.pulse {
			highPulseCount++
		} else {
			lowPulseCount++
		}

		var newEvents []Event
		if conjMod, ok := event.moduleTargeted.(*Conjunction); ok {
			newEvents = conjMod.ReceivePulseInput(event.pulse, event.moduleSource)
		} else {
			newEvents = event.moduleTargeted.ReceivePulse(event.pulse)
		}

		queue.enqueue(newEvents...)
	}

	return highPulseCount, lowPulseCount
}

func getModuleMap(lines []string) map[string]Module {
	var moduleMap = map[string]Module{}

	// get initial modules
	for _, line := range lines {
		parts := strings.Split(line, "->")
		sourceModule := strings.Trim(parts[0], " ")
		var newModule Module

		if sourceModule == broadcaster {
			newModule = &Broadcaster{name: sourceModule}
		} else {
			if string(sourceModule[0]) == `%` {
				newModule = &Flipflop{onState: false, name: sourceModule}
			} else if byte(sourceModule[0]) == '&' {
				newModule = &Conjunction{name: sourceModule, inputModulesPulses: map[Module]bool{}}
			}
			sourceModule = sourceModule[1:]
		}
		moduleMap[sourceModule] = newModule
	}

	for _, line := range lines {
		parts := strings.Split(line, "->")
		sourceModule := strings.Trim(parts[0], " ")

		re := regexp.MustCompile(`\w+`)
		destModulesList := re.FindAllString(parts[1], -1)

		if sourceModule != broadcaster {
			sourceModule = sourceModule[1:] // remove symbols
		}

		var destModules []Module
		sourceMod := moduleMap[sourceModule]

		for _, destMod := range destModulesList {
			var mod, ok = moduleMap[destMod]
			if !ok {
				mod = &Conjunction{name:destMod, inputModulesPulses: map[Module]bool{}} // no dest modules
				moduleMap[destMod] = mod
			}

			destModules = append(destModules, mod)

			if conjMod, ok := mod.(*Conjunction); ok {
				conjMod.AddInputModule(sourceMod)
			}
		}
		sourceMod.AddDestModules(destModules)
	}

	return moduleMap
}