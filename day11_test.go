package aoc2019

import (
	"image/color"
	"testing"
)

var test11_input = []int{3, 8, 1005, 8, 306, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 28, 2, 107, 3, 10, 1, 101, 19, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 59, 2, 5, 13, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1001, 8, 0, 85, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 107, 1006, 0, 43, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 132, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1001, 8, 0, 154, 2, 4, 1, 10, 2, 4, 9, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 183, 1, 1102, 5, 10, 1, 1102, 1, 10, 1006, 0, 90, 2, 9, 12, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1001, 8, 0, 221, 1006, 0, 76, 1006, 0, 27, 1, 102, 9, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 252, 2, 4, 9, 10, 1006, 0, 66, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 282, 1, 102, 19, 10, 101, 1, 9, 9, 1007, 9, 952, 10, 1005, 10, 15, 99, 109, 628, 104, 0, 104, 1, 21102, 1, 387240010644, 1, 21101, 0, 323, 0, 1105, 1, 427, 21102, 846541370112, 1, 1, 21101, 334, 0, 0, 1106, 0, 427, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102, 3425718295, 1, 1, 21102, 381, 1, 0, 1105, 1, 427, 21102, 179410541715, 1, 1, 21101, 0, 392, 0, 1106, 0, 427, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 718078255872, 1, 21101, 0, 415, 0, 1105, 1, 427, 21102, 1, 868494234468, 1, 21102, 1, 426, 0, 1105, 1, 427, 99, 109, 2, 21202, -1, 1, 1, 21101, 0, 40, 2, 21101, 458, 0, 3, 21101, 0, 448, 0, 1106, 0, 491, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 453, 454, 469, 4, 0, 1001, 453, 1, 453, 108, 4, 453, 10, 1006, 10, 485, 1102, 0, 1, 453, 109, -2, 2105, 1, 0, 0, 109, 4, 2102, 1, -1, 490, 1207, -3, 0, 10, 1006, 10, 508, 21102, 1, 0, -3, 22102, 1, -3, 1, 22101, 0, -2, 2, 21102, 1, 1, 3, 21102, 1, 527, 0, 1106, 0, 532, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 555, 2207, -4, -2, 10, 1006, 10, 555, 22101, 0, -4, -4, 1105, 1, 623, 22101, 0, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 574, 0, 0, 1105, 1, 532, 21202, 1, 1, -4, 21102, 1, 1, -1, 2207, -4, -2, 10, 1006, 10, 593, 21102, 0, 1, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 615, 21201, -1, 0, 1, 21101, 615, 0, 0, 106, 0, 490, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}

func TestDay11_Simple(t *testing.T) {
	tests := map[color.Color]int{
		color.Black: 1686,
		color.White: 249,
	}

	for c, res := range tests {
		r := NewRobot()
		go r.run(c)

		IntCode(test11_input, r.out, r.in)
		close(r.in)
		// wait for close
		<-r.out

		if painted := len(r.painted); painted != res {
			t.Errorf("expected %d, got %d", res, painted)
		}

		// uncomment to solve task 2:
		// t.Log(r.PrintPlate())
	}
}