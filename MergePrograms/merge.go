// Copyright 2014 Lars Pensjö
//
// This file is part of Aldcran.
//
// Aldcran is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// Aldcran is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Aldcran.  If not, see <http://www.gnu.org/licenses/>.
//

// Based on http://www.xmailserver.org/diff2.pdf
package merge

import (
	"fmt"
	"math/rand"
)

type program []byte
type edit_graph [][]bool

func findMatchPoints(x program, y program) edit_graph {
	rows := make([][]bool, len(y))
	for yi, y_instr := range y {
		rows[yi] = make([]bool, len(x))
		for xi, x_instr := range x {
			if x_instr == y_instr {
				rows[yi][xi] = true
			}
		}
	}
	return rows
}

type step int

const (
	right step = iota
	down
	diag
)

type path []step

// Simple algorithm, but not optimal
func findShortestPathButeForce(gr edit_graph, x, y int) (cost int, p path) {
	if y == len(gr) {
		p = path{}
		for ; x < len(gr[0]); x++ {
			cost++
			p = append(p, right)
		}
		return
	}
	if x == len(gr[0]) {
		p = path{}
		for ; y < len(gr); y++ {
			cost++
			p = append(p, down)
		}
		return
	}
	if gr[y][x] {
		addCost, rest := findShortestPathButeForce(gr, x+1, y+1)
		cost += addCost
		p = append(path{diag}, rest...)
		return
	}
	addCost1, rest1 := findShortestPathButeForce(gr, x+1, y)
	addCost2, rest2 := findShortestPathButeForce(gr, x, y+1)
	if addCost1 < addCost2 {
		cost = addCost1 + 1
		p = append(path{right}, rest1...)
	} else {
		cost = addCost2 + 1
		p = append(path{down}, rest2...)
	}
	return
}

// Take a path, do a random merge transform on two programs, and return the new program
func (p path) randomMerge(p1, p2 program) (newProg program) {
	commonPath := true
	newProg = program{}
	var chooseX bool
	var x, y int
	for i := range p {
		if p[i] == diag {
			commonPath = true
			newProg = append(newProg, p1[x])
			x++
			y++
			continue
		}
		if commonPath {
			// Detected a deviation that starts here
			commonPath = false
			if rand.Float32() > 0.5 {
				chooseX = true
			} else {
				chooseX = false
			}
		}
		if p[i] == right {
			if chooseX {
				newProg = append(newProg, p2[x])
			}
			x++
		} else if p[i] == down {
			if !chooseX {
				newProg = append(newProg, p1[y])
			}
			y++
		}
	}
	return
}

func Test() {
	rand.Seed(1) // Get the same behaviour every time
	x := program{0, 1, 2, 3, 4, 5}
	fmt.Println("Program 1:", x)
	y := program{0, 1, 6, 3}
	fmt.Println("Program 2:", y)
	graph := findMatchPoints(y, x)
	cost, p := findShortestPathButeForce(graph, 0, 0)
	fmt.Println("Testing, cost", cost, "path", p)
	for i := 0; i < 20; i++ {
		newProg := p.randomMerge(x, y)
		fmt.Println("Randomized:", newProg)
	}
}
