package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	file, err := os.Open("2025/day2/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	reader := bufio.NewReader(file)
	res := 0

	for {
		line, err := reader.ReadBytes(',')
		interval := strings.Split(strings.TrimRight(string(line), ","), "-")
		start, _ := strconv.Atoi(interval[0])
		end, _ := strconv.Atoi(interval[1])

		for i := start; i <= end; i++ {
			width := numDigits(i)

			mod := 1
			for range width / 2 {
				mod *= 10

				temp := i
				repeated := temp % mod
				temp = temp / mod
				for temp > 0 {
					if temp%mod != repeated || mod > temp*10 {
						break
					}
					temp = temp / mod
				}
				if temp == 0 {
					fmt.Println("Found: ", i)
					res += i
					break
				}
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("file read err: %v", err)
		}
	}

	return res
}

func numDigits(i int) int {
	if i == 0 {
		return 1
	}
	if i < 0 {
		i = -i
	}

	width := 0
	for i > 0 {
		i /= 10
		width++
	}

	return width
}
