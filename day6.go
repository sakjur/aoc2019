package aoc2019

import "strings"

func parseOrbit(s string) orbit {
	o := orbit{}
	lines := strings.Split(s, "\n")
	for _, s := range lines {
		orb := strings.Split(s, ")")
		o[orb[1]] = orb[0]
	}
	return o
}

type orbit map[string]string

func (o orbit) totalOrbits() int {
	orbits := map[string]int{}
	for body, parent := range o {
		c := 0
		for parent != "" {
			parent = o[parent]
			c++
		}
		orbits[body] = c
	}

	sum := 0
	for _, val := range orbits {
		sum += val
	}
	return sum
}
