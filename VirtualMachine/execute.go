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

const (
	ParClearThreshold = 100
)

func (p *program) run() {
	for pc := 0; pc < len(p.instructions); pc++ {
		p.instructions[pc].execute(p.memory)
	}
}

func (i *instruction) execute(memory []int) {
	var value int
	if i.clear > ParClearThreshold {
		value = 0
	}
	value *= i.multImmediate
	if ind := i.multIndirect; ind != 0 && int(ind) < len(memory) {
		value *= memory[ind]
	}
	value += i.addImmediate
	if ind := i.addIndirect; ind != 0 && int(ind) < len(memory) {
		value += memory[ind]
	}
	if addr := i.storeAddress; addr != 0 && int(addr) < len(memory) {
		memory[addr] = value
	}
	if ind := i.storeIndirect; ind != 0 && int(ind) < len(memory) && int(memory[ind]) < len(memory) {
		addr := memory[ind]
		memory[addr] = value
	}
}
