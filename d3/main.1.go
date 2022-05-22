package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const bits = 12

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// count 1 for each column
	counters := make([]int, bits)

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		for i := 0; i < len(counters); i++ {
			if l[i:i+1] == "1" {
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

	fmt.Println(gamma * epsilon)

}
