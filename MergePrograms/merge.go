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
	"log"
	"math"
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
	diag  step = iota // No change needed, both values are the same
	right             // Skip x-value from the string at current position
	down              // Add the y-value to the string at current position
)

// A path defines the changes (steps) needed to transform one string into another.
// Example, given the two strings x="00" and y="10". The path will now become "210" (steps).
// This will transform from x to y as follows:
// Step 2 (down): Add first character from y (1).
// Step 1 (right): Skip first character from x (0).
// Step 0: (diag): Add second character from x or y (0).
type path []step

type node struct {
	x, y int
	next int // The node that follows on this node
}

func findShortestPath2(a, b program) (cost int) {
	endNode := node{len(a), len(b), 0}
	buffer := []node{endNode}
	start := 0
	end := 1
	for iter := 0; ; iter++ {
		for i := start; i < end; i++ {
			n := buffer[i]
			x := n.x - 1
			y := n.y
			for x < len(a) && y < len(b) && x >= 0 && y >= 0 && a[x] == b[y] {
				x--
				y--
			}
			if x <= 0 && y <= 0 {
				return iter
			}
			// log.Println("Iter", iter, "adding", x, ",", y)
			buffer = append(buffer, node{x, y, 0})

			x = n.x
			y = n.y - 1
			for x < len(a) && y < len(b) && x >= 0 && y >= 0 && a[x] == b[y] {
				x--
				y--
			}
			if x <= 0 && y <= 0 {
				return iter
			}
			// log.Println("Iter", iter, "adding", x, ",", y)
			buffer = append(buffer, node{x, y, 0})
		}
		start = end
		end = len(buffer)
	}
}

func (gr edit_graph) findShortestPath() (cost int, p path) {
	for maxCost := 0; ; maxCost++ {
		cost, p, abort := findShortestPathBruteForce(gr, 0, 0, maxCost)
		if !abort {
			if cost != maxCost {
				log.Fatal("Expect costs to be the same (", cost, " and ", maxCost, ")")
			}
			return cost, p
		}
	}
	return
}

func findShortestPathBruteForce(gr edit_graph, x, y, maxCost int) (cost int, p path, abort bool) {
	if maxCost < 0 {
		// Cancel this attempt. 0 cost could still work, though.
		return 0, nil, true
	}
	if y == len(gr) {
		if len(gr[0])-x > maxCost {
			abort = true
			return
		}
		for ; x < len(gr[0]); x++ {
			cost++
			p = append(p, right)
		}
		return
	}
	if x == len(gr[0]) {
		if len(gr)-y > maxCost {
			abort = true
			return
		}
		for ; y < len(gr); y++ {
			cost++
			p = append(p, down)
		}
		return
	}
	bestCost := math.MaxUint32
	bestPath := path{}
	bestDir := right
	abort = true
	testCost, testPath, testAbort := findShortestPathBruteForce(gr, x+1, y, maxCost-1)
	if !testAbort {
		bestCost = testCost
		bestPath = testPath
		abort = false
	}
	testCost, testPath, testAbort = findShortestPathBruteForce(gr, x, y+1, maxCost-1)
	if !testAbort && testCost < bestCost {
		bestCost = testCost
		bestPath = testPath
		bestDir = down
		abort = false
	}
	if gr[y][x] {
		testCost, testPath, testAbort = findShortestPathBruteForce(gr, x+1, y+1, maxCost)
		if !testAbort && testCost-1 < bestCost {
			bestCost = testCost - 1
			bestPath = testPath
			bestDir = diag
			abort = false
		}
	}
	if abort {
		// Nothing found
		return 0, nil, true
	}
	cost = bestCost + 1
	p = append(path{bestDir}, bestPath...)
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
	cost, p := graph.findShortestPath()
	fmt.Println("Testing, cost", cost, "path", p)
	for i := 0; i < 20; i++ {
		newProg := p.randomMerge(x, y)
		fmt.Println("Randomized:", newProg)
	}
}
