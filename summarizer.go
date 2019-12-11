package aoc2019

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func summarizer(r io.Reader, fn func(int) int) (int, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(d), "\n")

	res := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		v, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		res += fn(v)
	}
	return res, nil
}
