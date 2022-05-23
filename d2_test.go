package main

import "testing"

func TestD2P1(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	res := d2P1(input)

	if res != 150 {
		t.Fatalf("got %d; expected %d", res, 150)
	}
}

func TestD2P2(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	res := d2P2(input)

	if res != 900 {
		t.Fatalf("got %d; expected %d", res, 900)
	}
}
