package aoc2019

import (
	"os"
	"strings"
	"testing"
)

func TestDay10_Simple(t *testing.T) {
	m := &meteors{
		height: 5,
		width:  5,
		field: []point{
			{1, 0}, {4, 0},
			{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
			{4, 3},
			{3, 4}, {4, 4},
		},
	}
	if loc := m.BestLoc(); loc.x != 3 || loc.y != 4 {
		t.Logf("expected BestLoc = {3 4}, got %v", loc)
	}

	m = ParseMeteorField(strings.NewReader(`.#..#
.....
#####
....#
...##`))

	t.Log(m.Polars(point{3, 4}))
	if loc := m.BestLoc(); loc.x != 3 || loc.y != 4 {
		t.Logf("expected BestLoc = {3 4}, got %v", loc)
	}

	m = ParseMeteorField(strings.NewReader(`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`))

	if loc := m.BestLoc(); loc.x != 5 || loc.y != 8 {
		t.Logf("expected BestLoc = {5 8}, got %v", loc)
	}
}

func TestDay10_Task(t *testing.T) {
	f, err := os.Open("testdata/day10.txt")
	if err != nil {
		panic(err)
	}

	m := ParseMeteorField(f)
	loc := m.BestLoc()

	if n := len(m.Polars(loc)); n != 278 {
		t.Errorf("expected 278, got %d", n)
	}

	if twohundredth := m.BlastOrder(loc)[199]; twohundredth.x != 14 || twohundredth.y != 17 {
		t.Errorf("expected 200th blasted comet to be {14 17}, was %v", twohundredth)
	}
}
