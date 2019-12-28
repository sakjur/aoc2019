package aoc2019

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestOrbit_Simple(t *testing.T) {
	test := orbit{
		"COM": "",
		"B":   "COM",
		"C":   "B",
		"D":   "C",
		"E":   "D",
		"F":   "E",
		"G":   "B",
		"H":   "G",
		"I":   "D",
		"J":   "E",
		"K":   "J",
		"L":   "K",
	}

	if test.totalOrbits() != 42 {
		t.Errorf("expected total orbits = 42, got %v", test.totalOrbits())
	}
}

const testOrbit = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

func TestOrbit_Parse(t *testing.T) {
	o := parseOrbit(testOrbit)
	if o.totalOrbits() != 42 {
		t.Errorf("expected total orbits = 42, got %v", o.totalOrbits())
	}
}

func TestOrbit_Task(t *testing.T) {
	f, err := os.Open("testdata/day6.txt")
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	o := parseOrbit(string(b))

	fmt.Println(o.totalOrbits())
}
