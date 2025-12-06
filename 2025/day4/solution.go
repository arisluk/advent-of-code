package day4

import (
	"bufio"
	"log"
	"os"
)

func Solve() int {
	file, err := os.Open("2025/day4/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	nRows, nCols := len(grid), len(grid[0])

	res := 0

	for {
		removed := 0
		for row := range nRows {
			for col := range nCols {
				if grid[row][col] == '@' && isAccessible(&grid, row, col, nRows, nCols) {
					removed++
					grid[row][col] = 'x'
				}
			}
		}

		for row := range nRows {
			for col := range nCols {
				if grid[row][col] == 'x' {
					grid[row][col] = '.'
				}
			}
		}

		if removed == 0 {
			break
		}

		res += removed
	}

	return res
}

func isAccessible(grid *[][]byte, row, col int, nRows, nCols int) bool {
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}

	count := 0
	for _, dir := range dirs {
		tempRow := row + dir[0]
		tempCol := col + dir[1]
		if tempRow >= nRows || tempRow < 0 || tempCol >= nCols || tempCol < 0 {
			continue
		} else if (*grid)[tempRow][tempCol] == '@' || (*grid)[tempRow][tempCol] == 'x' {
			count++
		}
	}

	return count < 4
}
