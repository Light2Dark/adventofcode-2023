package day19

import (
	"fmt"
	"os"
	"strings"
)

type Xmas struct {
	x,m,a,s [2]int
}

func Run2() {
	data, _ := os.ReadFile("day19/input.txt")
// 	data := `px{a<2006:qkq,m>2090:A,rfg}
// pv{a>1716:R,A}
// lnx{m>1548:A,A}
// rfg{s<537:gd,x>2440:R,A}
// qs{s>3448:A,lnx}
// qkq{x<1416:A,crn}
// crn{x>2662:A,R}
// in{s<1351:px,qqz}
// qqz{s>2770:qs,m<1801:hdj,R}
// gd{a>3333:R,R}
// hdj{m>838:A,pv}

// {x=787,m=2655,a=1222,s=2876}
// {x=1679,m=44,a=2067,s=496}
// {x=2036,m=264,a=79,s=2244}
// {x=2461,m=1339,a=466,s=291}
// {x=2127,m=1623,a=2188,s=1013}`

	configs := strings.Split(string(data), "\n\n")
	workflowStr := configs[0]

	workflows := GetWorkflows(workflowStr)

	var ranges = struct{ x, m, a, s [2]int }{
		x: [2]int{1, 4000},
		m: [2]int{1, 4000},
		a: [2]int{1, 4000},
		s: [2]int{1, 4000},
	}

	fmt.Println(Count(ranges, "in", workflows))
}

type Workflow struct {
	varName        string
	operator       string
	value          int
	targetWorkflow string
}

func GetWorkflows(workflowStr string) map[string][]Workflow {
	lines := strings.Split(workflowStr, "\n")
	var workflowMap = map[string][]Workflow{}

	for _, line := range lines {
		line = line[:len(line)-1]
		parts := strings.Split(line, "{")
		workflowName, step := parts[0], parts[1][:len(parts[1])]

		steps := strings.Split(step, ",")
		var instructions = []Workflow{}

		for _, step := range steps[:len(steps)-1] { // last one is fallback (diff structure)
			parts := strings.Split(step, ":")
			condition, outputWorkflow := parts[0], parts[1]
			variableName, num, operator := getVar(condition)

			var workflow = Workflow{varName: variableName, operator: operator, value: num, targetWorkflow: outputWorkflow}
			instructions = append(instructions, workflow)
		}

		var lastWorkflow = Workflow{varName: "", operator: "", value: 0, targetWorkflow: steps[len(steps)-1]}
		instructions = append(instructions, lastWorkflow)

		workflowMap[workflowName] = instructions
	}

	return workflowMap
}

func Count(ranges Xmas, workflowName string, workflows map[string][]Workflow) int {
	if workflowName == "R" {
		return 0
	} else if workflowName == "A" {
		return (ranges.x[1] - ranges.x[0] + 1) * (ranges.m[1] - ranges.m[0] + 1) * (ranges.a[1] - ranges.a[0] + 1) * (ranges.s[1] - ranges.s[0] + 1)
	}

	var workflowSteps = workflows[workflowName]
	var rules, fallback = workflowSteps[:len(workflowSteps)-1], workflowSteps[len(workflowSteps)-1]

	var key [2]int
	var total int
	var remaining bool = true

	for _, rule := range rules {
		switch rule.varName {
		case "x":
			key = ranges.x
		case "m":
			key = ranges.m
		case "a":
			key = ranges.a
		case "s":
			key = ranges.s
		}

		var trueHalf, falseHalf [2]int
		if rule.operator == "<" {
			trueHalf = [2]int{key[0], rule.value-1}
			falseHalf = [2]int{rule.value, key[1]}
		} else if rule.operator == ">" {
			trueHalf = [2]int{rule.value+1, key[1]}
			falseHalf = [2]int{key[0], rule.value}
		}

		if trueHalf[0] <= trueHalf[1] {
			// create new struct where the ranges are within the trueHalf, it will call the target worflow for this whole range
			// e.g. if x is 1-1000, and the rule is x<500, then the new struct will be 1-499
			var newStruct = setStruct(rule.varName, ranges, trueHalf) 
			total += Count(newStruct, rule.targetWorkflow, workflows)
		}
		if falseHalf[0] <= falseHalf[1] {
			// create new struct where the ranges are within the falseHalf
			// e.g. if x is 1-1000, and the rule is x<500, then the new struct will be 500-1000
			// if it's empty, then it will be rejected
			ranges = setStruct(rule.varName, ranges, falseHalf)
		} else {
			remaining = false
			break
		}
	}

	if remaining {
		total += Count(ranges, fallback.targetWorkflow, workflows)
	}

	return total
}

func StartWorkflowBool(variableMap map[string]int, workflowMap WorkflowMap, functionName string) bool {
	function := workflowMap[functionName]
	workflowName, accepted, err := function(variableMap)

	if err != nil {
		panic(err)
	}

	if workflowName == "" {
		return accepted
	} else {
		return StartWorkflowBool(variableMap, workflowMap, workflowName)
	}
}

func setStruct(key string, x Xmas, val [2]int) Xmas {
	var newStruct = x
	switch key {
	case "x":
		newStruct.x = val
	case "m":
		newStruct.m = val
	case "a":
		newStruct.a = val
	case "s":
		newStruct.s = val
	}

	return newStruct
}