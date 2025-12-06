package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	file, err := os.Open("2025/day3/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		jolt := strings.Builder{}
		start := 0
		for n := 11; n >= 0; n-- {
			largest := line[start]
			largestIdx := start
			for j := start; j < len(line)-n; j++ {
				if line[j] > largest {
					largest = line[j]
					largestIdx = j
				}
			}
			jolt.WriteByte(largest)
			start = largestIdx + 1
		}

		joltInt, _ := strconv.Atoi(jolt.String())
		res += joltInt
	}

	return res
}
