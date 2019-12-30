package aoc2019

import (
	"fmt"
	"strings"
)

type body struct {
	pos point3d
	vel point3d
}

type system []*body

func (s system) Add(b *body) system {
	return append(s, b)
}

func (s *system) Step(n int) {
	for i := 0; i < n; i++ {
		s.Gravity()
		s.Movement()
	}
}

func (s *system) Gravity() {
	for _, b := range *s {
		for _, o := range *s {
			b.Gravity(o)
		}
	}
}

func (s *system) Movement() {
	for _, b := range *s {
		b.pos.x += b.vel.x
		b.pos.y += b.vel.y
		b.pos.z += b.vel.z
	}
}

func (s *system) Energy() int {
	sum := 0
	for _, b := range *s {
		sum += b.Energy()
	}
	return sum
}

func (s *system) String() string {
	bodies := make([]string, len(*s))
	for i, b := range *s {
		bodies[i] = b.String()
	}
	return strings.Join(bodies, "\n")
}

func (b *body) Gravity(o *body) {
	b.vel.x += cmp(o.pos.x, b.pos.x)
	b.vel.y += cmp(o.pos.y, b.pos.y)
	b.vel.z += cmp(o.pos.z, b.pos.z)
}

func (b *body) Energy() int {
	return b.Potential() * b.Kinetic()
}

func (b *body) Potential() int {
	return abs(b.pos.x) + abs(b.pos.y) + abs(b.pos.z)
}

func (b *body) Kinetic() int {
	return abs(b.vel.x) + abs(b.vel.y) + abs(b.vel.z)
}

func (b *body) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", b.pos, b.vel)
}

func cmp(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
