package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Solve() int {
	file, err := os.Open("2025/day1/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	curr := 50
	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		dir := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("error parsing clicks: %v", err)
		}

		if dir == 'L' {
			temp := curr - clicks
			res += abs(temp / 100)
			if temp <= 0 && curr != 0 {
				res++
			}
			curr = (100 + temp%100) % 100
			// fmt.Println(string(dir), clicks, curr, temp, res)
		} else { // R
			temp := curr + clicks
			res += abs(temp / 100)
			if temp <= 0 && curr != 0 {
				res++
			}
			curr = (100 + temp%100) % 100
			// fmt.Println(string(dir), clicks, curr, temp, res)
		}
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
