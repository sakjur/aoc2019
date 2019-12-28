package aoc2019

import (
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

const testOrbitTransfer = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

func TestOrbit_Transfer(t *testing.T) {
	o := parseOrbit(testOrbitTransfer)
	if val := o.orbitTransfers("YOU", "SAN"); val != 4 {
		t.Errorf("expected 4 orbit transfers from YOU to SAN, got %v", val)
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

	//fmt.Println(o.totalOrbits())
	if val := o.totalOrbits(); val != 154386 {
		t.Errorf("expected 154386 orbits, got %d", val)
	}
	//fmt.Println(o.orbitTransfers("YOU", "SAN"))
	if val := o.orbitTransfers("YOU", "SAN"); val != 346 {
		t.Errorf("expected 346 orbits between YOU and SAN, got %d", val)
	}
}
