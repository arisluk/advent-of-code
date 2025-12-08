package day8

import (
	"bufio"
	"container/heap"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
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
	for i := 0; i < 1000; i++ {
		closest := heap.Pop(h).(Pair)

		d.makeSet(closest.First)
		d.makeSet(closest.Second)
		d.union(closest.First, closest.Second)
	}

	groups := make(map[int][]int)
	for x := range d {
		r := d.find(x)
		groups[r] = append(groups[r], x)
	}

	top3 := make([]int, 3)
	for _, g := range groups {
		length := len(g)
		if length >= top3[0] {
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = length
		} else if length >= top3[1] {
			top3[2] = top3[1]
			top3[1] = length
		} else if length >= top3[2] {
			top3[2] = length
		}
	}

	return top3[0] * top3[1] * top3[2]
}

func getDist(c1, c2 []int) float64 {
	return math.Sqrt(math.Pow(float64(c2[0]-c1[0]), 2) + math.Pow(float64(c2[1]-c1[1]), 2) + math.Pow(float64(c2[2]-c1[2]), 2))
}

type DSU map[int]int

func (d DSU) makeSet(x int) {
	if _, ok := d[x]; !ok {
		d[x] = x
	}
}

func (d DSU) find(x int) int {
	if d[x] != x {
		d[x] = d.find(d[x])
	}
	return d[x]
}

func (d DSU) union(a, b int) {
	ra, rb := d.find(a), d.find(b)
	if ra != rb {
		d[rb] = ra
	}
}

type Pair struct {
	Dist   float64
	First  int
	Second int
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() any {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}
