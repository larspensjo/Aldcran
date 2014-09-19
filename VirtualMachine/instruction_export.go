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
