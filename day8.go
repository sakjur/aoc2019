package aoc2019

import (
	"io"
	"math"
	"strings"
)

// space image format
type sif struct {
	width  int
	height int
	layers [][]int
}

func (s *sif) read(reader io.Reader) {
	buf := make([]byte, s.width*s.height)
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}

		layer := make([]int, len(buf))
		for pos, b := range buf {
			layer[pos] = int(b - '0')
		}

		s.layers = append(s.layers, layer)
	}
}

func (s *sif) cleanestLayer() int {
	z := s.countByLayer(0)
	min := math.MaxInt64
	layer := -1

	for i, c := range z {
		if c < min {
			layer = i
			min = c
		}
	}

	return layer
}

func (s *sif) countByLayer(n int) []int {
	res := make([]int, len(s.layers))
	for i := range s.layers {
		res[i] = s.occurrences(i, n)
	}
	return res
}

func (s *sif) occurrences(layer int, n int) int {
	c := 0
	for _, i := range s.layers[layer] {
		if i == n {
			c++
		}
	}
	return c
}

func (s *sif) decodeString() string {
	builder := strings.Builder{}

	for pixel := 0; pixel < s.width*s.height; pixel++ {
		if pixel%s.width == 0 {
			builder.WriteRune('\n')
		}
		for _, layer := range s.layers {
			if layer[pixel] == 0 {
				builder.WriteRune('█') // full block
				break
			} else if layer[pixel] == 1 {
				builder.WriteRune(' ') // full block
				break
			}
		}
	}
	return builder.String()
}
