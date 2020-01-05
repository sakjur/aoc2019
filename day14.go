package aoc2019

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

type chemical string

type chemQuantity struct {
	consume int
	chem    chemical
	produce int
}

type reaction struct {
	m    map[chemical][]chemQuantity
	rest map[chemical]int
}

func ParseReaction(r io.Reader) (reaction, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return reaction{}, err
	}

	res := reaction{
		m:    make(map[chemical][]chemQuantity),
		rest: make(map[chemical]int),
	}
	lines := bytes.Split(raw, []byte("\n"))
	for _, line := range lines {
		transformation := bytes.Split(line, []byte(" => "))
		components := bytes.Split(transformation[0], []byte(", "))
		produces, chem := parseChem(transformation[1])

		comps := []chemQuantity{}
		for _, comp := range components {
			consumes, chem := parseChem(comp)
			comps = append(comps, chemQuantity{
				consume: consumes,
				chem:    chem,
				produce: produces,
			})
		}
		res.m[chem] = comps
	}
	res.rest = make(map[chemical]int)
	return res, nil
}

func parseChem(b []byte) (num int, chem chemical) {
	parts := bytes.Split(b, []byte(" "))
	num, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		panic(err)
	}

	return num, chemical(parts[1])
}

func (r reaction) Ore(chem chemical) int {
	ore := 0
	produce := 0
	if r.rest[chem] == 0 {
		for _, c := range r.m[chem] {
			produce = c.produce
			switch c.chem {
			case "ORE":
				ore += c.consume
			default:
				for i := 0; i < c.consume; i++ {
					ore += r.Ore(c.chem)
				}
			}
		}
		r.rest[chem] += produce
	}
	r.rest[chem]--
	return ore
}

func (r reaction) ProduceFuel(oreSupply int) int {
	fuel := 0
	for oreSupply > 0 {
		oreSupply -= r.Ore("FUEL")
		if oreSupply < 0 {
			break
		}
		fuel++
		zeroed := true
		for _, val := range r.rest {
			if val != 0 {
				zeroed = false
				break
			}
		}
		if zeroed {
			fmt.Println(fuel)
		}
	}
	return fuel
}
