package aoc2019

import (
	"testing"
)

// this is just using the IntCode computer from Day 2/day 5 with _no further modifications_.

var day7Program = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 51, 76, 101, 126, 207, 288, 369, 450, 99999, 3, 9, 102, 4, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 1002, 9, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 1001, 9, 2, 9, 102, 2, 9, 9, 101, 3, 9, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 5, 9, 9, 102, 5, 9, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 1002, 9, 5, 9, 1001, 9, 5, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}

func TestDay7_Task(t *testing.T) {
	// Task 1
	perms := permutations([]int{0, 1, 2, 3, 4})
	if _, v := day7BestConfiguration(perms); v != 422858 {
		t.Errorf("expected day 7 task 1 result to be 422858, got %d", v)
	}
	// Task 2
	perms = permutations([]int{5, 6, 7, 8, 9})
	if _, v := day7BestConfiguration(perms); v != 14897241 {
		t.Errorf("expected day 7 task 2 result to be 14897241, got %d", v)
	}
}

func day7BestConfiguration(perms [][]int) (highestPerm []int, highest int) {
	for _, perm := range perms {
		iA := make(chan int)
		go func() {
			iA <- 0
		}()

		i := iA
		for _, n := range perm {
			i = amplifier(n, i)
		}

		last := 0

		// feedback
		for v := range i {
			last = v
			iA <- v
		}
		close(iA)

		if highestPerm == nil || highest < last {
			highestPerm = perm
			highest = last
		}
	}

	return
}

func amplifier(phase int, i chan int) (out chan int) {
	out = make(chan int)

	go func() {
		in := make(chan int)
		go func() {
			in <- phase
			for v := range i {
				in <- v
			}
			close(in)
		}()
		IntCode(day7Program, in, out)
		close(out)
	}()
	return
}

func permutations(in []int) [][]int {
	if len(in) == 1 {
		return [][]int{in}
	}

	res := [][]int{}
	for i, n := range in {
		vals := make([]int, len(in))
		copy(vals, in)

		childPermutations := permutations(append(vals[:i], vals[i+1:]...))
		for _, perm := range childPermutations {
			res = append(res, append([]int{n}, perm...))
		}
	}
	return res
}
