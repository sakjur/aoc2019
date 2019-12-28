package aoc2019

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestDay1_Simple(t *testing.T) {
	tests := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for input, output := range tests {
		if result := Fuel(input); result != output {
			t.Errorf("expected Fuel(%d) to return %d, got %d", input, output, result)
		}
	}
}

func TestDay1_FuelForFuel_Simple(t *testing.T) {
	tests := map[int]int{
		12:     2,
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for input, output := range tests {
		if result := FuelForFuel(input); result != output {
			t.Errorf("expected Fuel(%d) to return %d, got %d", input, output, result)
		}
	}
}

func TestDay1_Task1(t *testing.T) {
	f, err := os.Open("testdata/day1.txt")
	if err != nil {
		t.Error(err)
	}

	if n, err := summarizer(f, Fuel); n != 3464458 {
		t.Errorf("expected %d, got %d (with err: %v)", 3464458, n, err)
	}
}

func TestDay1_Task2(t *testing.T) {
	f, err := os.Open("testdata/day1.txt")
	if err != nil {
		t.Error(err)
	}

	if n, err := summarizer(f, FuelForFuel); n != 5193796 {
		t.Errorf("expected %d, got %d (with err: %v)", 5193796, n, err)
	}
}

func TestDay1_Sanity(t *testing.T) {
	f, err := os.Open("testdata/day1.txt")
	if err != nil {
		t.Error(err)
	}
	b, _ := ioutil.ReadAll(f)

	r := bytes.NewReader(b)
	if i, _ := summarizer(r, Fuel); i != 3464458 {
		t.Error("day 1 task 1 sanity failed")
	}

	r = bytes.NewReader(b)
	if i, _ := summarizer(r, Fuel); i != 3464458 {
		t.Error("day 1 task 2 sanity failed")
	}
}
