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
	"fmt"
	"github.com/larspensjo/go-monotonic-graycode"
)

type VirtualMachine struct {
	graycode *mgc.Mgc
	memory   []int
}

func New(width uint32, memorySize uint32) *VirtualMachine {
	var vm VirtualMachine
	vm.graycode = mgc.New(width)
	vm.memory = make([]int, memorySize)
	return &vm
}

// Create a program data structure
// Instructions have to be added afterwards
func (vm *VirtualMachine) NewProgram() *program {
	var p program
	p.virtualMachine = vm
	return &p
}

// Convert an instruction to a pretty string
func (i *instruction) String() (ret string) {
	if i.clear > parClearThreshold {
		ret = "Clear, "
	}
	if i.multImmediate != 1 {
		ret += fmt.Sprintf("*%d, ", i.multImmediate)
	}
	if ind := i.multIndirect; ind != 0 {
		ret += fmt.Sprintf("*mem[%d], ", i.multImmediate)
	}
	if i.addImmediate != 0 {
		ret += fmt.Sprintf("+%d, ", i.addImmediate)
	}
	if ind := i.addIndirect; ind != 0 {
		ret += fmt.Sprintf("+mem[%d], ", i.multImmediate)
	}
	if addr := i.storeAddress; addr != 0 {
		ret += fmt.Sprintf("store[%d], ", i.storeAddress)
	}
	if ind := i.storeIndirect; ind != 0 {
		ret += fmt.Sprintf("store[*%d], ", i.storeAddress)
	}

	if ret == "" {
		ret = "noop"
	}
	return
}

func (p *program) String() (ret string) {
	for i, ins := range p.instructions {
		ret += fmt.Sprintln(i, ": ", ins.String())
	}
	return
}
