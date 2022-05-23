package main

import (
	"math"
	"strconv"
)

func d3P1(lines []string, bits int) int {
	// count 1 for each column
	counters := make([]int, bits)

	for _, line := range lines {
		for i := 0; i < len(counters); i++ {
			if line[i:i+1] == "1" {
				counters[i]++
			} else {
				counters[i]--
			}
		}
	}

	gamma, epsilon := 0, 0
	for i, counter := range counters {
		pos := bits - i - 1
		// MCB is 1 of i-column
		if counter >= 0 {
			gamma |= int(math.Exp2(float64(pos)))
		} else {
			epsilon |= int(math.Exp2(float64(pos)))
		}
	}
	return gamma * epsilon
}

func d3P2(lines []string, bits int) int {
	oxygen := getLifeRatingParameter(lines, bits, func(mcb bool) string {
		if mcb {
			return "1"
		}
		return "0"
	})
	co2 := getLifeRatingParameter(lines, bits, func(mcb bool) string {
		if mcb {
			return "0"
		}
		return "1"
	})

	return oxygen * co2
}

func getLifeRatingParameter(lines []string, bits int, fn func(mcb bool) string) int {
	idx, mcbCounter := 0, 0
	var result string
	for i := 0; i < bits; i++ {
		for _, line := range lines {
			// get most common bit
			if line[idx:idx+1] == "1" {
				mcbCounter++
			} else {
				mcbCounter--
			}
		}

		result += fn(mcbCounter >= 0)

		var temp []string
		for _, line := range lines {
			if line[idx:idx+1] == result[idx:idx+1] {
				temp = append(temp, line)
			}
		}
		if len(temp) == 0 {
			result = lines[0]
			break
		}
		lines = temp
		idx++
		mcbCounter = 0
	}

	n, err := strconv.ParseInt(result, 2, 32)

	if err != nil {
		panic(err)
	}

	return int(n)
}
