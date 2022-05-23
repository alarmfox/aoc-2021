package main

import "testing"

func TestD1P1(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}

	res := d1P1(input)

	if res != 7 {
		t.Fatalf("got %d; expected %d", res, 7)
	}
}

func TestD1P2(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}

	res := d1P2(input)

	if res != 5 {
		t.Fatalf("got %d; expected %d", res, 5)
	}
}
