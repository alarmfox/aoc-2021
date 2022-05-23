package main

import (
	"strconv"
)

func d1P1(input []string) int {
	lastRead, counter := 0, 0
	for _, line := range input[1:] {
		n, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}
		if n > lastRead {
			counter++
		}

		lastRead = n
	}

	return counter
}

func d1P2(lines []string) int {
	var nums []int = make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}

	var currWindow []int = []int{0, 0, 0}
	counter := 0
	for i := 0; i < len(nums)-3; i++ {
		w := nums[i : i+3]
		if sum(w) > sum(currWindow) {
			counter++
		}
		currWindow[0], currWindow[1], currWindow[2] = w[0], w[1], w[2]
	}

	return counter

}
func sum(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
