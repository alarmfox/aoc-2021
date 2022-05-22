package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var p Position
	for s.Scan() {
		l := s.Text()
		parts := strings.Split(l, " ")
		n, _ := strconv.Atoi(parts[1])

		switch strings.ToLower(parts[0]) {
		case "forward":
			p.X += n
		case "down":
			p.Y += n
		case "up":
			p.Y -= n
		}
	}

	fmt.Println(p.X * p.Y)
}
