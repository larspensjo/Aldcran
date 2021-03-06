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

// These are parameters that can be tuned for better efficiency
const (
	parClearThreshold                  = 100
	parMultScaling                     = 100
	parPenaltyMultIllegalAddress       = 100
	parPenaltyAddIllegalAddress        = 100
	parPenaltyStoreIllegalAddress      = 100
	parPenaltyStoreIndirIllegalAddress = 100
)

// Run a program, and return the number of instructions that were executed
func (p *program) run() (cost int) {
	for pc := 0; pc < len(p.instructions); pc++ {
		cost++
		p.instructions[pc].execute(p)
	}
	return
}

func (i *instruction) execute(p *program) {
	memory := p.virtualMachine.memory
	var value int
	if i.clear > parClearThreshold {
		value = 0
	}
	// Convert to float while doing computation. Don't multiply with the exact argument,
	// use a down scaled value added with 1 to minimize impact
	value = int(float64(value)*(1+float64(i.multImmediate)/parMultScaling) + 0.5)
	if ind := i.multIndirect; ind != 0 {
		if int(ind) < len(memory) {
			value = int(float64(value)*(1+float64(memory[ind])/parMultScaling) + 0.5)
		} else {
			p.addPenalty(parPenaltyMultIllegalAddress)
		}
	}
	value += i.addImmediate
	if ind := i.addIndirect; ind != 0 {
		if int(ind) < len(memory) {
			value += memory[ind]
		} else {
			p.addPenalty(parPenaltyAddIllegalAddress)
		}
	}
	if addr := i.storeAddress; addr != 0 {
		if int(addr) < len(memory) {
			memory[addr] = value
		} else {
			p.addPenalty(parPenaltyStoreIllegalAddress)
		}
	}
	if ind := i.storeIndirect; ind != 0 {
		if int(ind) < len(memory) && int(memory[ind]) < len(memory) {
			addr := memory[ind]
			memory[addr] = value
		} else {
			p.addPenalty(parPenaltyStoreIndirIllegalAddress)
		}
	}
}

func (p *program) addPenalty(penalty int) {
	p.penalties += penalty
}
