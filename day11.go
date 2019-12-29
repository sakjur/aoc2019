package aoc2019

import (
	"bytes"
	"image/color"
)

type robot struct {
	in      chan int
	out     chan int
	painted map[point]color.Color
}

func NewRobot() *robot {
	return &robot{
		in:      make(chan int),
		out:     make(chan int),
		painted: make(map[point]color.Color),
	}
}

func (r *robot) run(start color.Color) {
	r.out <- colorCode(start)

	pos := point{}
	dir := 0

	for v := range r.in {
		var c color.Color
		switch v {
		case 0:
			c = color.Black
		case 1:
			c = color.White
		}
		r.painted[pos] = c

		switch rot := <-r.in; rot {
		case 0:
			dir = (4 + dir - 1) % 4
		case 1:
			dir = (4 + dir + 1) % 4
		}

		switch dir {
		case 0:
			pos.y++
		case 1:
			pos.x++
		case 2:
			pos.y--
		case 3:
			pos.x--
		}

		if col, exists := r.painted[pos]; exists {
			r.out <- colorCode(col)
		} else {
			r.out <- 0
		}
	}
	close(r.out)
}

func (r *robot) PrintPlate() string {
	var min, max point
	for sq, _ := range r.painted {
		if min.x > sq.x {
			min.x = sq.x
		} else if max.x < sq.x {
			max.x = sq.x
		}

		if min.y > sq.y {
			min.y = sq.y
		} else if max.y < sq.y {
			max.y = sq.y
		}
	}

	plate := make([][]byte, abs(min.y)+abs(max.y)+1)
	for i := range plate {
		plate[i] = make([]byte, abs(min.x)+abs(max.x)+1)
		for j := range plate[i] {
			plate[i][j] = 'x'
		}
	}

	for sq, col := range r.painted {
		var c byte
		switch col {
		case color.Black:
			c = 'x'
		case color.White:
			c = ' '
		default:
			panic("unknown color")
		}

		row := (len(plate) - 1) - (sq.y + abs(min.y))
		column := sq.x + abs(min.x)

		plate[row][column] = c
	}

	return "\n" + string(bytes.Join(plate, []byte("\n")))
}

func colorCode(c color.Color) int {
	switch c {
	case color.Black:
		return 0
	case color.White:
		return 1
	}
	return 0
}
