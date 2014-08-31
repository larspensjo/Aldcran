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

type instruction int
type program []instruction
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
	if x == len(gr) {
		p = path{}
		for ; y < len(gr); y++ {
			cost++
			p = append(p, down)
		}
		return
	}
	if y == len(gr)-1 && x == len(gr[0])-1 {
		p = path{diag}
		return
	}
	if gr[y][x] {
		addCost, rest := findShortestPathButeForce(gr, x+1, y+1)
		cost += addCost
		p = append(path{diag}, rest...)
		return
	}
	addCost1 := len(gr[0]) + 1
	addCost2 := len(gr) + 1
	var rest1, rest2 path
	if y == len(gr) {
		// Can't go further down
		addCost1, rest1 = findShortestPathButeForce(gr, x+1, y)
	} else if x == len(gr[0]) {
		// Can't go further right
		addCost2, rest2 = findShortestPathButeForce(gr, x, y+1)
	} else {
		addCost1, rest1 = findShortestPathButeForce(gr, x+1, y)
		addCost2, rest2 = findShortestPathButeForce(gr, x, y+1)
	}
	if addCost1 < addCost2 {
		cost = addCost1 + 1
		p = append(path{right}, rest1...)
	} else {
		cost = addCost2 + 1
		p = append(path{down}, rest2...)
	}
	return
}
