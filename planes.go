package aoc2019

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type point3d struct {
	x int
	y int
	z int
}

func Points3D(input string) []point3d {
	points := []point3d{}

	for i := 0; i < len(input); i++ {
		readNumber := func() (int, int) {
			i += 2
			end := strings.IndexAny(input[i:], ",>")
			n, err := strconv.Atoi(input[i : i+end])
			if err != nil {
				panic(err)
			}
			i = i + end
			return i, n
		}

		p := point3d{}

		if input[i] != '<' {
			continue
		}
		// gulp the <
		i++
	inner:
		for {
			switch input[i] {
			case '>':
				break inner
			default:
				i++
			case 'x':
				i, p.x = readNumber()
			case 'y':
				i, p.y = readNumber()
			case 'z':
				i, p.z = readNumber()
			}
		}

		points = append(points, p)
	}

	return points
}

func (p point3d) String() string {
	return fmt.Sprintf("<x=%2d, y=%2d, z=%2d>", p.x, p.y, p.z)
}
