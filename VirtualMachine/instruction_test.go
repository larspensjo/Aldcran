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

package vm

import (
	"testing"
)

func TestNoop(t *testing.T) {
	var p program
	p.run() // Running without memory or any instructions
	if p.penalties != 0 {
		t.Error("Expected no errors from empty program")
	}
	p.instructions = []instruction{noop}
	p.run()
	if p.penalties != 0 {
		t.Error("Expected no errors from nop")
	}
}

func TestAdd(t *testing.T) {
	var p program
	p.instructions = []instruction{{addImmediate: 1, storeAddress: 1}}
	p.run() // Running with no memory installed
	if p.penalties == 0 {
		t.Error("Should be memory error")
	}
	p.penalties = 0
	p.memory = make([]int, 2)
	p.run()
	if p.penalties > 0 {
		t.Error("Shouldn't be error from addimmediate")
	}
	if p.memory[1] != 1 {
		t.Error("addimmediate should give 1, but had", p.memory[1])
	}
}
