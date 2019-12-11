package aoc2019

import "testing"

const day2_input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,6,19,23,1,23,13,27,2,6,27,31,1,5,31,35,2,10,35,39,1,6,39,43,1,13,43,47,2,47,6,51,1,51,5,55,1,55,6,59,2,59,10,63,1,63,6,67,2,67,10,71,1,71,9,75,2,75,10,79,1,79,5,83,2,10,83,87,1,87,6,91,2,9,91,95,1,95,5,99,1,5,99,103,1,103,10,107,1,9,107,111,1,6,111,115,1,115,5,119,1,10,119,123,2,6,123,127,2,127,6,131,1,131,2,135,1,10,135,0,99,2,0,14,0"

func TestDay2_Simple(t *testing.T) {
	tests := map[string]string{
		"1,9,10,3,2,3,11,0,99,30,40,50": "3500,9,10,70,2,3,11,0,99,30,40,50",
		"1,0,0,0,99": "2,0,0,0,99",
		"2,3,0,3,99": "2,3,0,6,99",
		"2,4,4,5,99,0": "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99": "30,1,1,4,2,5,6,0,99",
	}

	for input, output := range tests {
		expected, err := stringToIntSlice(output)
		if err != nil {
			t.Errorf("when reading expected list, got err: %v", err)
		}

		vals, err := stringToIntSlice(input)
		if err != nil {
			t.Errorf("when reading input list, got err: %v", err)
		}

		res, err := IntCode(vals)
		if err != nil {
			t.Errorf("got error when running IntCode(%s): %v", input, err)
		}

		if IntCodeString(res) != IntCodeString(expected) {
			t.Errorf("expected IntCode(%s) == %s, got %s", input, IntCodeString(expected), IntCodeString(res))
		}
	}
}

func TestDay2_Task1(t *testing.T) {
	t.SkipNow()
	vals, err := stringToIntSlice(day2_input)
	if err != nil {
		t.Error(err)
		return
	}

	vals[1] = 12
	vals[2] = 2

	res, err := IntCode(vals)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res[0])
}

func TestDay2_Task2(t *testing.T) {
	t.SkipNow()
	vals, err := stringToIntSlice(day2_input)
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 1e4; i++ {
		vals[1] = i % 100
		vals[2] = i / 100

		res, err := IntCode(vals)
		if err != nil {
			continue
		}

		if res[0] == 19690720 {
			t.Logf("%02d%02d == %d", vals[1], vals[2], res[0])
			return
		}
	}
}

func TestDay2_Sanity(t *testing.T) {
	vals, err := stringToIntSlice(day2_input)
	if err != nil {
		t.Error(err)
		return
	}

	vals[1] = 12
	vals[2] = 02
	if res, _ := IntCode(vals); res[0] != 2782414 {
		t.Error("day 2 task 1 sanity fail")
	}

	vals[1] = 98
	vals[2] = 20
	if res, _ := IntCode(vals); res[0] != 19690720 {
		t.Error("day 2 task 2 sanity fail")
	}
}
