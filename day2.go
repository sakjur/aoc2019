package aoc2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func stringToIntSlice(input string) ([]int, error) {
	ops := strings.Split(input, ",")

	output := []int{}
	for _, val := range ops {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		output = append(output, i)
	}
	return output, nil
}

func IntCodeString(codes []int) string {
	s := []string{}
	for _, code := range codes {
		s = append(s, strconv.Itoa(code))
	}

	return strings.Join(s, ",")
}

func IntCode(input []int, keyboard chan int, output chan int) ([]int, error) {
	codes := make([]int, len(input))
	copy(codes, input)
	for i := 0; i < len(codes); {
		switch codes[i] % 100 {
		case 99:
			return codes, nil
		case 1:
			codes[codes[i+3]] = val(i, 1, codes) + val(i, 2, codes)
			i += 4
		case 2:
			codes[codes[i+3]] = val(i, 1, codes) * val(i, 2, codes)
			i += 4
		case 3:
			codes[codes[i+1]] = <-keyboard
			i += 2
		case 4:
			output <- val(i, 1, codes)
			i += 2
		default:
			return nil, fmt.Errorf("unknown op code '%d'", codes[i])
		}
	}
	return nil, fmt.Errorf("program incorrectly terminated")
}

func val(i int, param int, codes []int) int {
	op := codes[i]
	mode := (op / int(math.Pow10(param+1))) % 10

	switch mode {
	case 0:
		return codes[codes[i+param]]
	case 1:
		return codes[i+param]
	default:
		panic(fmt.Errorf("unexpected mode %d", mode))
	}
}
