package day6

import (
	"bufio"
	"log"
	"os"
)

func Solve2() int {
	file, err := os.Open("2025/day6/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	res := 0
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	operators := lines[len(lines)-1]
	nums := lines[:len(lines)-1]

	currOp := operators[0]
	currOperands := []int{}
	for i := range len(operators) {
		if i != 0 && operators[i] != ' ' {
			currRes := 0
			if currOp == '*' {
				currRes = 1
			}
			for _, ele := range currOperands {
				if currOp == '+' {
					currRes += ele
				} else {
					currRes *= ele
				}
			}
			res += currRes

			currOp = operators[i]
			currOperands = []int{}
		}

		currOperand := 0
		separator := true
		for _, line := range nums {
			if line[i] != ' ' {
				separator = false
				currOperand = currOperand*10 + int(line[i]-'0')
			}
		}
		if !separator {
			currOperands = append(currOperands, currOperand)
		}
	}

	currRes := 0
	if currOp == '*' {
		currRes = 1
	}
	for _, ele := range currOperands {
		if currOp == '+' {
			currRes += ele
		} else {
			currRes *= ele
		}
	}
	res += currRes

	return res
}
