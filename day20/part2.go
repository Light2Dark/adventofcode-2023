package day20

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 4 conjunction modules feed into xn module (conjunction module)
// xn conjunc module feed into rx mod
// the concept is that the 4 conjunction modules when aligned (all high) will produce a low pulse to the xn mod
// so in order to find when they all align, need to find their lcm

func Run2() {
	data, _ := os.ReadFile("day20/input.txt")
	lines := strings.Split(string(data), "\n")

	moduleMap := getModuleMap(lines)

	// get the input to rx mod
	var inputRx Module
	if mod, ok := moduleMap["rx"].(*Conjunction); ok {
		inputMods := mod.inputModulesPulses
		if len(inputMods) > 1 {
			panic(errors.New("input to rx mod should only be 1"))
		}
		for k := range inputMods {
			inputRx = k
		}
	}

	fourConjucMods := inputRx.(*Conjunction).inputModulesPulses

	var cycleLengths = map[Module][]int{}
	for k := range fourConjucMods {
		cycleLengths[k] = []int{}
	}

	for i := 1;i<10000; i++ {
		var firstEvent = Event{pulse: false, moduleTargeted: moduleMap[broadcaster]}
		var queue = &Queue{events: []Event{firstEvent}}
	
		for queue.length() > 0 {
			event, _ := queue.dequeue()
	
			var newEvents []Event
			if conjMod, ok := event.moduleTargeted.(*Conjunction); ok {
				newEvents = conjMod.ReceivePulseInput(event.pulse, event.moduleSource)
			} else {
				newEvents = event.moduleTargeted.ReceivePulse(event.pulse)
			}
	
			if event.moduleTargeted == inputRx && event.pulse {
				_, ok := cycleLengths[event.moduleSource]
				if ok {
					cycleLengths[event.moduleSource] = append(cycleLengths[event.moduleSource], i)
				}
			}
	
			queue.enqueue(newEvents...)
		}
	}

	// first val in cycle length represents the num they repeat at
	var nums []int
	for _, v := range cycleLengths {
		nums = append(nums, v[0])
	}

	fmt.Println(lcm(nums...))
}

func lcm(integers ...int) int {
	if len(integers) < 2 {
		panic(errors.New("at least two integers are required"))
	}

	result := integers[0]
	for i := 1; i < len(integers); i++ {
		result = result * integers[i] / gcd(result, integers[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}