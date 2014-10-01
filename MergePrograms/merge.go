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

import "math/rand"

type program []byte

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

func followDiag(x, y int, a, b program) (int, int) {
	for x >= 0 && y >= 0 && a[x] == b[y] {
		x--
		y--
	}
	return x, y
}

func findShortestPath(a, b program) (int, path) {
	cost, list := findShortestPath2(a, b)
	p := interpretPath(a, b, list)
	return cost, p
}

func findShortestPath2(a, b program) (int, []node) {
	x, y := followDiag(len(a)-1, len(b)-1, a, b)
	endNode := node{x: x, y: y, next: -1}
	buffer := []node{endNode}
	if x == -1 && y == -1 {
		return 0, buffer
	}
	start := 0
	end := 1
	// TODO: Don't add nodes with x or y less than -1
	for iter := 1; ; iter++ {
		for i := start; i < end; i++ {
			n := buffer[i]
			x, y = followDiag(n.x-1, n.y, a, b)
			// log.Println("Iter", iter, "adding", x, ",", y)
			buffer = append(buffer, node{x: x, y: y, next: i})
			if x == -1 && y == -1 {
				return iter, buffer
			}

			x, y = followDiag(n.x, n.y-1, a, b)
			// log.Println("Iter", iter, "adding", x, ",", y)
			buffer = append(buffer, node{x: x, y: y, next: i})
			if x == -1 && y == -1 {
				return iter, buffer
			}
		}
		start = end
		end = len(buffer)
	}
}

func interpretPath(a, b program, list []node) path {
	p := path{}
	x, y := -1, -1
	last := len(list) - 1
	for i := list[last].next; i >= 0; {
		n := list[i]
		for x < n.x && y < n.y {
			x++
			y++
			// fmt.Println("Common", a[x], x, y)
			p = append(p, diag)
		}
		if x == n.x {
			// fmt.Println("Adding", b[n.y], n.x, n.y)
			p = append(p, down)
		} else {
			// fmt.Println("Removing", a[n.x], n.x, n.y)
			p = append(p, right)
		}
		x = n.x
		y = n.y
		i = list[i].next
	}
	x++
	y++
	for x < len(a) && y < len(b) {
		// fmt.Println("Common", a[x], x, y)
		p = append(p, diag)
		x++
		y++
	}
	return p
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
