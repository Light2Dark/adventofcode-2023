package day19

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WorkflowFunc func(variableMap VariableMap) (workflowName string, accepted bool, err error)
type WorkflowMap map[string]WorkflowFunc
type VariableMap map[string]int

func Run() {
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
	workflowStr, partStr := configs[0], configs[1]
	var workflowMap WorkflowMap = createWorkflowMap(workflowStr)

	parts := strings.Split(partStr, "\n")
	var sum int
	for _, part := range parts {
		part = part[1 : len(part)-1] // remove brackets
		components := strings.Split(part, ",")
		var variableMap = VariableMap{}
		for _, c := range components {
			comps := strings.Split(c, "=")
			num, _ := strconv.Atoi(comps[1])
			variableMap[comps[0]] = num
		}

		sum += StartWorkflow(variableMap, workflowMap, "in")
	}

	fmt.Println(sum)
}

func createWorkflowMap(workflowStr string) WorkflowMap {
	var workflowMap = WorkflowMap{}
	lines := strings.Split(workflowStr, "\n")
	for _, line := range lines {
		line = line[:len(line)-1]
		parts := strings.Split(line, "{")
		funcName := parts[0]
		workflowMap[funcName] = makeFunction(parts[1])
	}
	return workflowMap
}

func StartWorkflow(variableMap VariableMap, workflowMap WorkflowMap, functionName string) int {
	function := workflowMap[functionName]
	workflowName, accepted, err := function(variableMap)

	if err != nil {
		panic(err)
	}

	if workflowName == "" {
		if accepted {
			return findSumVariables(variableMap)
		} 
		return 0 // rejected
	} else {
		return StartWorkflow(variableMap, workflowMap, workflowName)
	}
}

func findSumVariables(variableMap VariableMap) int {
	var sum int
	for _, v := range variableMap {
		sum += v
	}
	return sum
}

func makeFunction(args string) WorkflowFunc {
	steps := strings.Split(args, ",")

	var result = func(variableMap VariableMap) (workflowName string, accepted bool, err error) {
		for _, step := range steps {
			if step == "A" {
				return "", true, nil
			} else if step == "R" {
				return "", false, nil
			}

			if strings.Contains(step, ":") {
				parts := strings.Split(step, ":")
				condition, outputWorkflow := parts[0], parts[1]
				variableName, num, operator := getVar(condition)

				val, ok := variableMap[variableName]
				if ok {
					if (operator == "<" && val < num) || (operator == ">" && val > num) {
						if outputWorkflow == "A" {
							return "", true, nil
						} else if outputWorkflow == "R" {
							return "", false, nil
						}
						return outputWorkflow, false, nil
					}
				}
			} else {
				return step, false, nil // only workflow name is found, and returned
			}
		}

		return "", false, errors.New("none of the conditions met")
	}

	return result
}

func getVar(condition string) (string, int, string) {
	if strings.Contains(condition, "<") {
		x := strings.Split(condition, "<")
		variableName := x[0]
		number, _ := strconv.Atoi(x[1])
		return variableName, number, "<"
	} else if strings.Contains(condition, ">") {
		x := strings.Split(condition, ">")
		variableName := x[0]
		number, _ := strconv.Atoi(x[1])
		return variableName, number, ">"
	} else {
		return "", 0, ""
	}
}
