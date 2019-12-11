package aoc2019

import "fmt"

func PasswordCandidates(from, to int) []string {
	// task 2

	candidates := passwordCandidates(from, to)
	kept := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		if shouldKeep(candidate) {
			kept = append(kept, candidate)
		}
	}

	return kept
}

func shouldKeep(candidate string) bool {
	for i := 0; i < len(candidate); i++ {
		adj := 1
		for j := i+1; j < len(candidate); j++ {
			if candidate[i] == candidate[j] {
				adj++
				i++
			}
		}
		if adj == 2 {
			return true
		}
	}
	return false
}

func passwordCandidates(from, to int) []string {
	// task 1
	candidates := []string{}

	outer:
	for i := from; i < to; i++ {
		pwd := fmt.Sprintf("%d", i)
		if len(pwd) != 6 {
			continue
		}

		var twoAdj bool
		prev := '\000'

		for _, c := range pwd {
			if c == prev {
				twoAdj = true
			}
			if c < prev {
				continue outer
			}
			prev = c
		}

		if twoAdj {
			candidates = append(candidates, pwd)
		}
	}
	return candidates
}
