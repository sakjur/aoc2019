package aoc2019

type tile int

const (
	empty tile = iota
	wall
	block
	hpaddle
	ball
)

type arcade struct {
	in       chan int
	joystick chan int
	tiles    map[point]tile
}

func StartArcade(in chan int) *arcade {
	arc := arcade{
		in:       in,
		joystick: make(chan int, 1),
		tiles:    make(map[point]tile),
	}

	return &arc
}

func (a *arcade) Run() int {
	score := 0
	ballPos := point{}
	paddlePos := point{}
	ballMov := make(chan struct{})

	var done bool
	go func() {
		for !done {
			<-ballMov
			pad := paddlePos
			bal := ballPos

			if pad.x == bal.x {
				a.joystick <- 0
			} else if pad.x > bal.x {
				a.joystick <- -1
			} else if pad.x < bal.x {
				a.joystick <- 1
			}
		}
		close(a.joystick)
	}()

	for x := range a.in {
		y := <-a.in
		if x == -1 && y == 0 {
			score = <-a.in
			continue
		}
		p := point{x: x, y: y}
		t := tile(<-a.in)
		a.tiles[p] = t

		if t == ball {
			ballPos = p
			ballMov <- struct{}{}
		} else if t == hpaddle {
			paddlePos = p
		}
	}
	close(ballMov)
	done = true
	return score
}

func (a *arcade) Ball() point {
	for p, t := range a.tiles {
		if t == ball {
			return p
		}
	}
	return point{}
}

func (a *arcade) Paddle() point {
	for p, t := range a.tiles {
		if t == hpaddle {
			return p
		}
	}
	return point{}
}

func (a *arcade) CountTiles(obj tile) int {
	c := 0
	for _, t := range a.tiles {
		if t == obj {
			c++
		}
	}
	return c
}
