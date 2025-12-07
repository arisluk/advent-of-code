package day6

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	file, err := os.Open("2025/day6/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	res := 0
	nums := [][]int{}
	operators := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if parts[0] == "*" || parts[0] == "+" {
			operators = parts
		} else {
			curr := []int{}
			for _, ele := range parts {
				temp, _ := strconv.Atoi(ele)
				curr = append(curr, temp)
			}
			nums = append(nums, curr)
		}
	}

	for i, op := range operators {
		var currRes int
		if op == "*" {
			currRes = 1
		} else {
			currRes = 0
		}

		for j := range len(nums) {
			if op == "*" {
				currRes *= nums[j][i]
			} else {
				currRes += nums[j][i]
			}
		}

		res += currRes
	}

	return res
}
