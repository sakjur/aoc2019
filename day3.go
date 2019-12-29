package aoc2019

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
)

type grid map[point]int

func CollisionDetection(r io.Reader) (int, int) {
	directions, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	paths := bytes.Split(directions, []byte("\n"))

	var lines []grid

	for _, p := range paths {
		lines = append(lines, pathIntoLine(p))
	}

	if len(lines) < 2 {
		panic("must have two paths")
	}

	intersection := []point{}
	for p := range lines[0] {
		if _, exists := lines[1][p]; exists {
			intersection = append(intersection, p)
		}
	}

	shortestDistance := math.MaxInt64
	signalDistance := math.MaxInt64

	for _, i := range intersection {
		distance := manhattanDistance(point{}, i)
		if distance < shortestDistance {
			shortestDistance = distance
		}
		signal := lines[0][i] + lines[1][i]
		if signal < signalDistance {
			signalDistance = signal
		}
	}

	return shortestDistance, signalDistance
}

func pathIntoLine(p []byte) grid {
	var direction byte
	var length int
	var current point
	var step = 0

	line := grid{}

	for i := 0; i < len(p); i++ {
		for i < len(p) && p[i] != ',' {
			b := p[i]
			if direction == 0 {
				direction = b
			}

			if b >= '0' && b <= '9' {
				f := b - '0'
				length = length*10 + int(f)
			}

			i++
		}
		points := movements[direction](current, length)
		for _, p := range points {
			current = p
			step++
			if _, exists := line[current]; !exists {
				line[current] = step
			}
		}
		direction = 0
		length = 0
	}

	return line
}

var movements = map[byte]func(p point, n int) []point{
	'U': func(p point, n int) []point {
		points := []point{}
		for i := 0; i < n; i++ {
			p.y++
			points = append(points, p)
		}
		return points
	},
	'D': func(p point, n int) []point {
		points := []point{}
		for i := 0; i < n; i++ {
			p.y--
			points = append(points, p)
		}
		return points
	},
	'R': func(p point, n int) []point {
		points := []point{}
		for i := 0; i < n; i++ {
			p.x++
			points = append(points, p)
		}
		return points
	},
	'L': func(p point, n int) []point {
		points := []point{}
		for i := 0; i < n; i++ {
			p.x--
			points = append(points, p)
		}
		return points
	},
}

func manhattanDistance(a point, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
