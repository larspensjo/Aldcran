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

package merge

import (
	"math/rand"
	"testing"
)

func TestFindPathEqual(t *testing.T) {
	x := program{1}
	y := x
	cost, p := findShortestPath(x, y)
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}

	x = program{1, 2}
	y = x
	cost, p = findShortestPath(x, y)
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}

	x = program{1, 2, 3, 4, 5}
	y = x
	cost, p = findShortestPath(x, y)
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}
}

// Tests where everything differs
func TestFindPathDifferent(t *testing.T) {
	var x, y program

	x = program{1}
	y = program{99}
	cost, p := findShortestPath(x, y)
	expCost := len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2}
	cost, p = findShortestPath(x, y)
	expCost = len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	// Same as previous, but toggle arguments
	cost, p = findShortestPath(x, y)
	expCost = len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2, 3, 4, 5}
	cost, p = findShortestPath(x, y)
	expCost = len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}
}

func TestFindPathMixed(t *testing.T) {
	var x, y program

	x = program{1, 2}
	y = program{1, 3}
	cost, p := findShortestPath(x, y)
	expCost := 2
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 3}
	y = program{2, 3}
	cost, p = findShortestPath(x, y)
	expCost = 2
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2, 3}
	y = program{1, 4, 5}
	cost, p = findShortestPath(x, y)
	expCost = 4
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2, 3}
	y = program{1, 4, 3}
	cost, p = findShortestPath(x, y)
	expCost = 2
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{'a', 'b', 'c', 'a', 'b', 'b', 'a'}
	y = program{'c', 'b', 'a', 'b', 'a', 'c'}
	cost, p = findShortestPath(x, y)
	expCost = 5
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}
}

func TestMerge(t *testing.T) {
	rand.Seed(1) // Get the same behaviour every time
	x := program{1, 2, 3, 4, 5, 6, 7, 8}
	y := program{1, 9, 3, 0, 5, 0, 7, 0}
	cost, p := findShortestPath(y, x)
	t.Log("Cost between parents", cost)
	for i := 0; i < 10; i++ {
		newProg := p.randomMerge(x, y)
		// Compare the child program to each of the parents. The difference to them should be
		// less or same compared to the difference between the parents
		newCost1, _ := findShortestPath(x, newProg)
		t.Log("Child prog", newProg, "cost", newCost1)

		newCost2, _ := findShortestPath(y, newProg)
		if newCost1+newCost2 != cost {
			t.Error("Invalid cost", newCost1, "+", newCost2, "!=", cost)
		}
		t.Log("Child prog", newProg, "cost", newCost2)
	}
}

func TestDiffLong(t *testing.T) {
	// x := program{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// x := program{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0}
	// y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	x := program{0, 1, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0, 0}
	y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// x := program{0, 0, 1, 1}
	// y := program{0, 0, 0, 0}
	cost, list := findShortestPath2(x, y)
	t.Log("Cost", cost, "vector length", len(list))
	p := interpretPath(x, y, list)
	t.Log(p)
}

func TestMergeLong(t *testing.T) {
	// This is a use case that failed at one point
	// x := program{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	x := program{1, 0}
	// y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	y := program{0, 0}
	numDiffs := 0
	if len(x) != len(y) {
		t.Fatal("This test depends on strings being of equal length")
	}
	for i := range x {
		if x[i] != y[i] {
			numDiffs++
		}
	}
	cost, p := findShortestPath(y, x)
	t.Log("Path:", p)
	if cost != numDiffs*2 {
		t.Error("Cost was", cost, "but expected cost was", numDiffs*2)
	}
	t.Log("Cost between parents", cost)
	newProg := p.randomMerge(x, y)
	if len(newProg) != len(x) {
		t.Error("New prog had length", len(newProg), "while expected", len(x))
	}
	t.Log(x)
	t.Log(y)
	t.Log(newProg)
}
