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

func findShortestPath(gr edit_graph) {

}
