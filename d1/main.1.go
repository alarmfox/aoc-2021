package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	lastRead, counter := 0, 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		n, err := strconv.Atoi(l)

		if err != nil {
			panic(err)
		}
		if n > lastRead {
			counter++
		}

		lastRead = n
	}
	fmt.Println(counter - 1)
}
