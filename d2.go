package main

import (
	"strconv"
	"strings"
)

func d2P1(lines []string) int {
	x, y := 0, 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch strings.ToLower(parts[0]) {
		case "forward":
			x += n
		case "down":
			y += n
		case "up":
			y -= n
		}
	}

	return x * y
}

func d2P2(lines []string) int {
	x, y, aim := 0, 0, 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		n, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		switch strings.ToLower(parts[0]) {
		case "forward":
			x += n
			y += n * aim
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}

	return x * y
}
