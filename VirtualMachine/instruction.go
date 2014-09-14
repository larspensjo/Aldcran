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

// Implement a virtual machine used by Aldcran.
// The propeties of this is:
// 1. Use a very wide instruction set
// 2. Only one instruction, which means no opcode is needed
// 3. A program shall be resilient to bit mutations. It shall change the behaviour, but have low risk of breaking the program completely
package vm

import (
	"fmt"
	"github.com/larspensjo/go-monotonic-graycode"
)

type instruction struct {
	clear         int // Clear operand if > ParClearThreshold
	addImmediate  int
	addIndirect   int
	multImmediate int
	multIndirect  int
	storeAddress  int
	storeIndirect int
}

type program struct {
	instructions []instruction
	grayCode     *mgc.Mgc
	memory       []int
	// The sum of all penalties during execution.
	// This will be used to evaluate the success/failure indicator of the algorithm
	penalties int
}

type subroutine struct {
	id int32 // Identifying the subroutine
	pr program
}

func main() {
	fmt.Println("Hello World!")
}
