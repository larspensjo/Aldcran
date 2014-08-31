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
	cost, p := findShortestPathButeForce(graph, 0, 0)
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}

	x = program{1, 2}
	y = x
	graph = findMatchPoints(x, y)
	cost, p = findShortestPathButeForce(graph, 0, 0)
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}

	x = program{1, 2, 3, 4, 5}
	y = x
	graph = findMatchPoints(x, y)
	cost, p = findShortestPathButeForce(graph, 0, 0)
	if cost != 0 || len(p) != len(x) {
		t.Error("Expected cost 0, path length", len(x), "got cost", cost, "with path", p)
	}
}

func TestFindPathDifferent(t *testing.T) {
	x := program{1}
	y := program{2}
	graph := findMatchPoints(x, y)
	cost, p := findShortestPathButeForce(graph, 0, 0)
	if cost != len(x) || p == nil {
		t.Error("Expected cost 0, got", cost)
	}

	x = program{1, 2, 3, 4, 5}
	graph = findMatchPoints(x, y)
	cost, p = findShortestPathButeForce(graph, 0, 0)
	if cost != len(x) || p == nil {
		t.Error("Expected cost 0, got", cost)
	}
	t.Error(cost, p)
}
