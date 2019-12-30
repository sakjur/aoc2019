package aoc2019

import "testing"

func TestDay12_Simple(t *testing.T) {
	s := &system{
		{pos: point3d{x: -1, y: 0, z: 2}},
		{pos: point3d{x: 2, y: -10, z: -7}},
		{pos: point3d{x: 4, y: -8, z: 8}},
		{pos: point3d{x: 3, y: 5, z: -1}},
	}

	s.Step(10)
	expected := `pos=<x= 2, y= 1, z=-3>, vel=<x=-3, y=-2, z= 1>
pos=<x= 1, y=-8, z= 0>, vel=<x=-1, y= 1, z= 3>
pos=<x= 3, y=-6, z= 1>, vel=<x= 3, y= 2, z=-3>
pos=<x= 2, y= 0, z= 4>, vel=<x= 1, y=-1, z=-1>`
	if expected != s.String() {
		t.Errorf("expected:\n%s\n\ngot:\n%s", expected, s.String())
	}

	if s.Energy() != 179 {
		t.Errorf("expected energy to be 179, got %d", s.Energy())
	}

	s = &system{
		{pos: point3d{x: -1, y: 0, z: 2}},
		{pos: point3d{x: 2, y: -10, z: -7}},
		{pos: point3d{x: 4, y: -8, z: 8}},
		{pos: point3d{x: 3, y: 5, z: -1}},
	}
	if s.Periodicity() != 2772 {
		t.Fail()
	}
}

func TestDay12_Parse(t *testing.T) {
	input := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
	points := Points3D(input)
	s := &system{}
	for _, p := range points {
		tmp := (*s).Add(&body{pos: p})
		s = &tmp
	}

	expected := `pos=<x=-1, y= 0, z= 2>, vel=<x= 0, y= 0, z= 0>
pos=<x= 2, y=-10, z=-7>, vel=<x= 0, y= 0, z= 0>
pos=<x= 4, y=-8, z= 8>, vel=<x= 0, y= 0, z= 0>
pos=<x= 3, y= 5, z=-1>, vel=<x= 0, y= 0, z= 0>`
	if s.String() != expected {
		t.Errorf("expected:\n%s\n\ngot:\n%s", expected, s.String())
	}
}

func TestDay12_Task(t *testing.T) {
	s := taskSystem()
	s.Step(1000)
	if s.Energy() != 8454 {
		t.Logf("expected energy to be 8454, got %d", s.Energy())
	}

	// task 2 takes a ton of time.
	if false {
		s = taskSystem()
		t.Log(s.Periodicity()) // 362 336 016 722 948 iterations
	}
}

func taskSystem() *system {
	in := `<x=-10, y=-13, z=7>
<x=1, y=2, z=1>
<x=-15, y=-3, z=13>
<x=3, y=7, z=-4>`
	points := Points3D(in)
	s := &system{}
	for _, p := range points {
		tmp := (*s).Add(&body{pos: p})
		s = &tmp
	}
	return s
}
