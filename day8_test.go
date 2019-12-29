package aoc2019

import (
	"os"
	"strings"
	"testing"
)

func TestDay8_Simple(t *testing.T) {
	input := "123456789012"
	img := &sif{width: 3, height: 2}
	img.read(strings.NewReader(input))

	layer := img.cleanestLayer()
	if n := img.occurrences(layer, 1) * img.occurrences(layer, 2); n != 1 {
		t.Errorf("expected n = 1, got n = %d", n)
	}
}

func TestDay8_ImageRender(t *testing.T) {
	input := "0222112222120000"
	img := &sif{width: 2, height: 2}
	img.read(strings.NewReader(input))

	expected := "\n█ \n █"
	if img.decodeString() != expected {
		t.Errorf("expected: %s\n\ngot: %s", expected, img.decodeString())
	}
}

func TestDay8_Task(t *testing.T) {
	f, err := os.Open("testdata/day8.txt")
	if err != nil {
		t.Error(err)
	}

	img := &sif{width: 25, height: 6}
	img.read(f)

	layer := img.cleanestLayer()
	if n := img.occurrences(layer, 1) * img.occurrences(layer, 2); n != 1206 {
		t.Errorf("expected n = 1206, got n = %d", n)
	}

	t.Log(img.decodeString())
	t.Log("...not going to write an actual test for that.")
}
