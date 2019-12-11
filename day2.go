package aoc2019

import (
	"fmt"
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

func IntCode(input []int) ([]int, error) {
	codes := make([]int, len(input))
	copy(codes, input)
	for i := 0; i < len(codes); i+=4 {
		switch codes[i] {
		case 99:
			return codes, nil
		case 1:
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
		case 2:
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
		default:
			return nil, fmt.Errorf("unknown op code '%d'", codes[i])
		}
	}
	return nil, fmt.Errorf("program incorrectly terminated")
}
