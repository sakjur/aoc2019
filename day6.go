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

func (o orbit) orbitTransfers(alpha, beta string) int {
	// in same orbit
	if o[alpha] == o[beta] {
		return 0
	}

	parent := o[alpha]
	alphaAncestors := map[string]int{}
	for c := 0; parent != ""; c++ {
		alphaAncestors[parent] = c
		parent = o[parent]
	}

	parent = o[beta]
	c := 0
	for parent != "" {
		if v, exists := alphaAncestors[parent]; exists {
			c += v
			break
		}
		parent = o[parent]
		c++
	}

	return c
}
