package day7

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func Solve2() int {
	file, err := os.Open("2025/day7/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		temp := []int{}
		for _, ele := range line {
			switch ele {
			case '.':
				temp = append(temp, 0)
			case '^':
				temp = append(temp, -1)
			default:
				temp = append(temp, -2)
			}
		}
		lines = append(lines, temp)
	}

	lines[1][slices.Index(lines[0], -2)] = 1

	for i := 1; i < len(lines)-2; i += 2 {
		for j := range lines[i] {
			if lines[i][j] != 0 {
				if lines[i+1][j] == -1 {
					lines[i+2][j-1] += lines[i][j]
					lines[i+2][j+1] += lines[i][j]
				} else {
					lines[i+2][j] += lines[i][j]
				}
			}
		}
	}

	res := 0
	for _, ele := range lines[len(lines)-1] {
		res += ele
	}

	return res
}
