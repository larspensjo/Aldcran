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
)

type numberVM int32
type register int32

const (
	ParClearThreshold numberVM = 100
)

type instruction struct {
	clear         numberVM // Clear operand if > ParClearThreshold
	addImmediate  numberVM
	addIndirect   register
	multImmdeiate numberVM
	multIndirect  register
}

type subroutine struct {
	id     int32 // Identifying the subroutine
	instrs []instruction
}

func main() {
	fmt.Println("Hello World!")
}
