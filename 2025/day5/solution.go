package day5

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	file, err := os.Open("2025/day5/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	intervals := [][]int{}
	ingredients := []int{}
	parseStatus := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			parseStatus = 1
			continue
		}

		if parseStatus == 0 { // parsing intervals
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			intervals = append(intervals, []int{start, end})
		} else { // reading ingredients
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	intervals = merge(intervals)

	res := 0
	for _, ing := range ingredients {
		lo, hi := 0, len(intervals)

		for lo < hi {
			mid := (lo + hi) / 2
			if intervals[mid][0] <= ing {
				lo = mid + 1
			} else {
				hi = mid
			}
		}

		targetInt := lo - 1
		if targetInt < 0 || intervals[targetInt][1] < ing {
			continue
		} else {
			res++
		}
	}

	freshCount := 0
	for _, interval := range intervals {
		freshCount += interval[1] - interval[0] + 1
	}

	return res, freshCount
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	curr := intervals[0]
	for _, interval := range intervals[1:] {
		if interval[0] > curr[1] {
			res = append(res, curr)
			curr = interval
		} else {
			curr[1] = max(curr[1], interval[1])
		}
	}
	res = append(res, curr)

	return res
}
