package day8

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve2() int {
	file, err := os.Open("2025/day8/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	coords := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		coord := []int{}
		for _, ele := range parts {
			i, _ := strconv.Atoi(ele)
			coord = append(coord, i)
		}
		coords = append(coords, coord)
	}

	h := &PairHeap{}
	heap.Init(h)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			heap.Push(h, Pair{getDist(coords[i], coords[j]), i, j})
		}
	}

	d := make(DSU)
	isOne := false
	for len(*h) > 0 && !isOne {
		closest := heap.Pop(h).(Pair)

		d.makeSet(closest.First)
		d.makeSet(closest.Second)
		d.union(closest.First, closest.Second)

		// Check all connected in one circuit
		if len(d) != len(coords) {
			continue
		}

		last := -1
		isOne = true
		for x := range d {
			if last == -1 {
				last = d.find(x)
			}
			if last != d.find(x) {
				isOne = false
				break
			}
		}

		if isOne {
			return coords[closest.First][0] * coords[closest.Second][0]
		}
	}

	return -1
}
