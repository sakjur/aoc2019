package aoc2019

import (
	"sync"
	"testing"
)

// Added extra IntCode instructions to day2.go

var day5_input = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 2, 171, 209, 224, 1001, 224, -1040, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 102, 65, 102, 224, 101, -3575, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 9, 82, 224, 1001, 224, -738, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1101, 52, 13, 224, 1001, 224, -65, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 1102, 82, 55, 225, 1001, 213, 67, 224, 1001, 224, -126, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1, 217, 202, 224, 1001, 224, -68, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1002, 176, 17, 224, 101, -595, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1102, 20, 92, 225, 1102, 80, 35, 225, 101, 21, 205, 224, 1001, 224, -84, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1101, 91, 45, 225, 1102, 63, 5, 225, 1101, 52, 58, 225, 1102, 59, 63, 225, 1101, 23, 14, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 359, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 1001, 223, 1, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 419, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 1001, 223, 1, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 449, 1001, 223, 1, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 524, 101, 1, 223, 223, 107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 554, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 1001, 223, 1, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 599, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 614, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 644, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

func TestDay5_Simple(t *testing.T) {
	i := make(chan int, 1)
	i <- 1
	o := make(chan int)
	res := []int{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := range o {
			res = append(res, i)
		}
		wg.Done()
	}()
	_, _ = IntCode([]int{3, 0, 4, 0, 99}, i, o)
	close(o)
	wg.Wait()

	if len(res) != 1 && res[0] != 1 {
		t.Errorf("expected len(res) = 1 & res[0] = 1, got len(res) = %d & res[0] = %d", len(res), res[0])
	}

	res, _ = IntCode([]int{1002, 4, 3, 4, 33}, nil, nil)
	if res[4] != 99 {
		t.Errorf("expected res[4] = 99, got res = %v", res)
	}

	res, _ = IntCode([]int{1101, 100, -1, 4, 0}, nil, nil)
	if res[4] != 99 {
		t.Errorf("expected res[4] = 99, got res = %v", res)
	}
}

func TestDay5_Tasks(t *testing.T) {
	for _, num := range []int{1, 5} {
		i := make(chan int, 1)
		i <- num
		o := make(chan int)
		output := []int{}
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for i := range o {
				output = append(output, i)
			}
			wg.Done()
		}()
		_, err := IntCode(day5_input, i, o)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		close(o)
		wg.Wait()

		for i, val := range output {
			if i == len(output)-1 {
				continue
			}
			if val != 0 {
				t.Errorf("expected all outputs except last to be 0, got %d", val)
			}
		}

		switch num {
		case 1:
			if output[len(output)-1] != 9006673 {
				t.Errorf("expected output from input 1 = 9006673")
			}
		case 5:
			if output[len(output)-1] != 3629692 {
				t.Errorf("expected output from input 5 = 3629692")
			}
		default:
			t.Error("unknown case")
		}
	}
}

func TestDay5_Output(t *testing.T) {
	type expected struct {
		program []int
		in      int
		out     int
	}

	longExample := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	tests := []expected{
		{
			program: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			in:      8,
			out:     1,
		},
		{
			program: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			in:      7,
			out:     0,
		},
		{
			program: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			in:      9,
			out:     0,
		},
		{
			program: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			in:      7,
			out:     1,
		},
		{
			program: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			in:      8,
			out:     0,
		},
		{
			program: longExample,
			in:      8,
			out:     1000,
		},
		{
			program: longExample,
			in:      7,
			out:     999,
		},
		{
			program: longExample,
			in:      9,
			out:     1001,
		},
	}

	for _, test := range tests {
		i := make(chan int, 1)
		i <- test.in
		o := make(chan int, 1)
		_, _ = IntCode(test.program, i, o)
		val := <-o
		if val != test.out {
			t.Errorf("expected IntCode(%v) = %d with input %d, got %d", test.program, test.out, test.in, val)
		}
	}
}
