package aoc2019

import (
	"os"
	"strings"
	"testing"
)

func TestDay14_Simple(t *testing.T) {
	r := reaction{
		m: map[chemical][]chemQuantity{
			"A":    {{10, "ORE", 10}},
			"B":    {{1, "ORE", 1}},
			"C":    {{7, "A", 1}, {1, "B", 1}},
			"D":    {{7, "A", 1}, {1, "C", 1}},
			"E":    {{7, "A", 1}, {1, "D", 1}},
			"FUEL": {{7, "A", 1}, {1, "E", 1}},
		},
		rest: make(map[chemical]int),
	}
	if r.Ore("FUEL") != 31 {
		t.Fail()
	}
}

func TestDay14_Parse(t *testing.T) {
	test := `157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`

	r, _ := ParseReaction(strings.NewReader(test))
	if ore := r.Ore("FUEL"); ore != 13312 {
		t.Error("expected 13312, got: ", ore)
	}
}

func TestDay14_Task(t *testing.T) {
	f, _ := os.Open("testdata/day14.txt")
	r, _ := ParseReaction(f)
	if ore := r.Ore("FUEL"); ore != 443537 {
		t.Error("expected 443537, got: ", ore)
	}
}

func TestDay14_Task2(t *testing.T) {
	t.Skip("this takes literally hours to run. Needs to optimize to solve task")
	f, _ := os.Open("testdata/day14.txt")
	r, _ := ParseReaction(f)
	r.ProduceFuel(1e12)
}
