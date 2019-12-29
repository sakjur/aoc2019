package aoc2019

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type meteors struct {
	height int
	width  int
	field  []point
}

func ParseMeteorField(r io.Reader) *meteors {
	fields, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fields), "\n")

	m := &meteors{
		height: len(lines),
		width:  len(lines[0]),
		field:  make([]point, 0),
	}

	for i, row := range lines {
		for j, column := range row {
			if column == '#' {
				m.field = append(m.field, point{x: j, y: i})
			}
		}
	}

	return m
}

func (m *meteors) Polars(alpha point) map[float64][]point {
	polars := make(map[float64][]point, len(m.field))

	for _, carth := range m.field {
		dx := float64(alpha.x - carth.x)
		dy := float64(alpha.y - carth.y)

		if dx == 0 && dy == 0 {
			continue
		}

		arc := math.Atan2(dy, dx)*57.2958 - 90
		if arc < 0 {
			arc = 360 + arc
		}

		switch _, exists := polars[arc]; exists {
		case true:
			polars[arc] = append(polars[arc], carth)
		case false:
			polars[arc] = []point{carth}
		}
	}

	return polars
}

func (m *meteors) BestLoc() point {
	max := 0
	maxP := point{}
	for _, p := range m.field {
		n := len(m.Polars(p))
		if n > max {
			max = n
			maxP = p
		}
	}
	return maxP
}

func (m *meteors) Paint() string {
	field := make([][]byte, m.height)
	for i := range field {
		field[i] = make([]byte, m.width)
		for j := range field[i] {
			field[i][j] = '.'
		}
	}

	for _, p := range m.field {
		field[p.y][p.x] = '#'
	}

	return "\n" + string(bytes.Join(field, []byte("\n")))
}

func (m *meteors) PaintWeights() string {
	field := make([][]byte, m.height)
	for i := range field {
		field[i] = make([]byte, m.width)
		for j := range field[i] {
			field[i][j] = '.'
		}
	}

	for _, p := range m.field {
		arcs := len(m.Polars(p))
		field[p.y][p.x] = byte(arcs) + '0'
	}

	return "\n" + string(bytes.Join(field, []byte("\n")))
}

func (m *meteors) BlastOrder(p point) []point {
	blast := make([]point, 0)

	coords := m.Polars(p)
	sortOrder := make([]float64, len(coords))

	i := 0
	for key := range coords {
		sortOrder[i] = key
		i++
	}

	sort.Float64s(sortOrder)

	points := make([][]point, len(coords))
	for i, val := range sortOrder {
		points[i] = coords[val]
		sort.Slice(points[i], func(a, b int) bool {
			return distance(p, points[i][a]) < distance(p, points[i][b])
		})
	}

	sploded := true
	for rot := 0; sploded == true; rot++ {
		sploded = false

		for _, aim := range points {
			if rot < len(aim) {
				blast = append(blast, aim[rot])
				sploded = true
			}
		}
	}

	return blast
}

func distance(a, b point) float64 {
	dx := float64(abs(a.x - b.x))
	dy := float64(abs(a.y - b.y))
	return math.Sqrt(dx*dx + dy*dy)
}
