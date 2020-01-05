package aoc2019

import (
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type isWall bool
type dir int8

const (
	north dir = 1
	south     = 2
	west      = 3
	east      = 4
)

type repairDroid struct {
	pos          point
	explored     map[point]isWall
	pause        sync.Mutex
	oxygenSystem point
}

func (r *repairDroid) start(resp chan int) chan int {
	out := make(chan int)
	r.explored = make(map[point]isWall)
	r.explored[r.pos] = false

	go func() {
		d := randomDir()
		out <- int(d)
		for in := range resp {
			r.pause.Lock()
			switch in {
			case 0:
				r.explored[r.move(d)] = true
			case 1:
				r.pos = r.move(d)
				r.explored[r.pos] = false
			case 2:
				r.pos = r.move(d)
				r.explored[r.pos] = false
				r.oxygenSystem = r.pos
			}
			r.pause.Unlock()
			d = randomDir()
			out <- int(d)
		}

		close(out)
	}()

	return out
}

func (r *repairDroid) String() string {
	r.pause.Lock()
	defer r.pause.Unlock()

	min, max := math.MaxInt64, math.MinInt64

	for p := range r.explored {
		if min > p.x {
			min = p.x
		}
		if min > p.y {
			min = p.y
		}
		if max < p.x {
			max = p.x
		}
		if max < p.y {
			max = p.y
		}
	}

	sz := abs(min) + max + 1
	grid := make([][]rune, sz)

	for row := range grid {
		grid[row] = make([]rune, sz)
		for column := range grid[row] {
			grid[row][column] = '?'
		}
	}

	for p, t := range r.explored {
		p2 := toGridPos(p, min, max)
		if t {
			grid[p2.y][p2.x] = '#'
		} else {
			grid[p2.y][p2.x] = '.'
		}
	}

	water := toGridPos(r.oxygenSystem, min, max)
	if water.y != 0 || water.x != 0 {
		grid[water.y][water.x] = 'H'
	}

	origo := toGridPos(point{}, min, max)
	grid[origo.y][origo.x] = 'S'

	res := []string{}
	for _, row := range grid {
		res = append(res, string(row))
	}

	return strings.Join(res, "\n")
}

func toGridPos(p point, min, max int) point {
	var y int
	if p.y < 0 {
		y = max + abs(p.y)
	} else {
		y = max - p.y
	}

	return point{x: abs(min) + p.x, y: y}
}

func (r *repairDroid) move(d dir) point {
	p := r.pos
	switch d {
	case north:
		p.y++
	case south:
		p.y--
	case west:
		p.x--
	case east:
		p.x++
	}
	return p
}

func randomDir() dir {
	rand.Seed(time.Now().UnixNano())
	return dir(rand.Intn(4) + 1)
}
