package day7

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func Solve() int {
	file, err := os.Open("2025/day7/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	res := 0
	lines := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []byte(line))
	}

	lines[1][slices.Index(lines[0], 'S')] = '|'

	for i := 1; i < len(lines)-2; i += 2 {
		for j := range lines[i] {
			if lines[i][j] == '|' {
				if lines[i+1][j] == '^' {
					res++
					lines[i+2][j-1] = '|'
					lines[i+2][j+1] = '|'
				} else {
					lines[i+2][j] = '|'
				}
			}
		}
	}

	return res
}
