package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")

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

	fmt.Println(counter)

}

func readInt(s *bufio.Scanner) (int, error) {
	if !s.Scan() {
		return 0, fmt.Errorf("scan error")
	}
	l := s.Text()

	return strconv.Atoi(l)
}

func sum(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
