package aoc2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func stringToIntSlice(input string) ([]int, error) {
	ops := strings.Split(input, ",")

	output := []int{}
	for _, val := range ops {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		output = append(output, i)
	}
	return output, nil
}

func IntCodeString(codes []int) string {
	s := []string{}
	for _, code := range codes {
		s = append(s, strconv.Itoa(code))
	}

	return strings.Join(s, ",")
}

type memory struct {
	reel []int
	base int
}

func IntCode(input []int, keyboard chan int, output chan int) ([]int, error) {
	m := memory{
		base: 0,
	}

	{
		codes := make([]int, len(input))
		copy(codes, input)
		m.reel = codes
	}
	for i := 0; i < len(m.reel); {
		switch m.reel[i] % 100 {
		case 99:
			return m.reel, nil
		case 1: // ADD
			m.set(i, 3, m.val(i, 1)+m.val(i, 2))
			i += 4
		case 2: // MUL
			m.set(i, 3, m.val(i, 1)*m.val(i, 2))
			i += 4
		case 3: // INPUT
			m.set(i, 1, <-keyboard)
			i += 2
		case 4: // OUTPUT
			output <- m.val(i, 1)
			i += 2
		case 5: // JUMP IF TRUE
			if m.val(i, 1) != 0 {
				i = m.val(i, 2)
			} else {
				i += 3
			}
		case 6: // JUMP IF FALSE
			if m.val(i, 1) == 0 {
				i = m.val(i, 2)
			} else {
				i += 3
			}
		case 7: // LT
			if m.val(i, 1) < m.val(i, 2) {
				m.set(i, 3, 1)
			} else {
				m.set(i, 3, 0)
			}
			i += 4
		case 8: // EQ
			if m.val(i, 1) == m.val(i, 2) {
				m.set(i, 3, 1)
			} else {
				m.set(i, 3, 0)
			}
			i += 4
		case 9: // REL BASE
			m.base += m.val(i, 1)
			i += 2
		default:
			return nil, fmt.Errorf("unknown op code '%d'", m.reel[i])
		}
	}
	return nil, fmt.Errorf("program incorrectly terminated")
}

func (m *memory) set(i, param, val int) {
	op := m.reel[i]
	mode := (op / int(math.Pow10(param+1))) % 10
	addr := -1

	switch mode {
	case 0:
		addr = m.reel[i+param]
	case 2:
		addr = m.base + m.reel[i+param]
	}

	m.fit(addr)
	m.reel[addr] = val
}

func (m *memory) val(i int, param int) int {
	op := m.reel[i]
	mode := (op / int(math.Pow10(param+1))) % 10
	addr := -1

	switch mode {
	case 0:
		addr = m.reel[i+param]
	case 1:
		addr = i + param
	case 2:
		addr = m.base + m.reel[i+param]
	default:
		panic(fmt.Errorf("unexpected mode %d", mode))
	}

	m.fit(addr)
	return m.reel[addr]
}

func (m *memory) fit(n int) {
	if n+1 > len(m.reel) {
		tmp := m.reel
		m.reel = make([]int, n+1)
		copy(m.reel, tmp)
	}
}
