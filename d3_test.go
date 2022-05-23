package main

import "testing"

func TestD3P1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	res := d3P1(input, 5)

	if res != 198 {
		t.Fatalf("got %d; expected %d", res, 198)
	}
}

func TestD3P2(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	res := d3P2(input, 5)

	if res != 230 {
		t.Fatalf("got %d; expected %d", res, 230)
	}
}
