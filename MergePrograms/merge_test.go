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

func compareEditGraphs(a, b edit_graph) bool {
	if len(a) != len(b) {
		return false
	}
	for y := range a {
		if len(b[y]) != len(a[y]) {
			return false
		}
		for x := range a[y] {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}

func TestEditGraph(t *testing.T) {
	x := program{1, 2}
	y := program{1, 2, 3}
	graph := findMatchPoints(x, y)
	expect := edit_graph{{true, false}, {false, true}, {false, false}}
	if !compareEditGraphs(expect, graph) {
		t.Error("Expected:", expect, "got", graph)
	}
}

func TestFindPathEqual(t *testing.T) {
	x := program{1}
	y := x
	graph := findMatchPoints(x, y)
	cost, p := graph.findShortestPath()
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}

	x = program{1, 2}
	y = x
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}

	x = program{1, 2, 3, 4, 5}
	y = x
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}
}

// Tests where everything differs
func TestFindPathDifferent(t *testing.T) {
	var x, y program
	var graph edit_graph

	x = program{1}
	y = program{99}
	graph = findMatchPoints(x, y)
	cost, p := graph.findShortestPath()
	expCost := len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2}
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	expCost = len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	// Same as previous, but toggle arguments
	graph = findMatchPoints(y, x)
	cost, p = graph.findShortestPath()
	expCost = len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2, 3, 4, 5}
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	expCost = len(x) + len(y)
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}
}

func TestFindPathMixed(t *testing.T) {
	var x, y program
	var graph edit_graph

	x = program{1, 2}
	y = program{1, 3}
	graph = findMatchPoints(x, y)
	cost, p := graph.findShortestPath()
	expCost := 2
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 3}
	y = program{2, 3}
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	expCost = 2
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2, 3}
	y = program{1, 4, 5}
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	expCost = 4
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{1, 2, 3}
	y = program{1, 4, 3}
	graph = findMatchPoints(x, y)
	cost, p = graph.findShortestPath()
	expCost = 2
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}

	x = program{'a', 'b', 'c', 'a', 'b', 'b', 'a'}
	y = program{'c', 'b', 'a', 'b', 'a', 'c'}
	graph = findMatchPoints(y, x)
	cost, p = graph.findShortestPath()
	expCost = 5
	if cost != expCost || p == nil {
		t.Error("Expected cost", expCost, "got", cost)
	}
}

func TestMerge(t *testing.T) {
	rand.Seed(1) // Get the same behaviour every time
	x := program{1, 2, 3, 4, 5, 6, 7, 8}
	y := program{1, 9, 3, 0, 5, 0, 7, 0}
	graph := findMatchPoints(y, x)
	cost, p := graph.findShortestPath()
	t.Log("Cost between parents", cost)
	for i := 0; i < 10; i++ {
		newProg := p.randomMerge(x, y)
		// Compare the child program to each of the parents. The difference to them should be
		// less or same compared to the difference between the parents
		graph := findMatchPoints(x, newProg)
		newCost1, _ := graph.findShortestPath()
		t.Log("Child prog", newProg, "cost", newCost1)

		graph = findMatchPoints(y, newProg)
		newCost2, _ := graph.findShortestPath()
		if newCost1+newCost2 != cost {
			t.Error("Invalid cost", newCost1, "+", newCost2, "!=", cost)
		}
		t.Log("Child prog", newProg, "cost", newCost2)
	}
}

func TestDiffLong(t *testing.T) {
	x := program{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// x := program{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0}
	// y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// x := program{0, 1, 0, 0, 0, 2, 0, 0, 3, 4, 0, 0}
	// y := program{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	cost := findShortestPath2(x, y)
	t.Log("Cost", cost)
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
	graph := findMatchPoints(y, x)
	t.Log("Graph:", graph)
	cost, p := graph.findShortestPath()
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
