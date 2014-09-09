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
	"bytes"
	"fmt"
	"github.com/larspensjo/go-monotonic-graycode"
)

const (
	ParClearThreshold = 100
)

type instruction struct {
	clear         int32 // Clear operand if > ParClearThreshold
	addImmediate  int32
	addIndirect   int32
	multImmediate int32
	multIndirect  int32
}

type program struct {
	instructions []instruction
}

type subroutine struct {
	id int32 // Identifying the subroutine
	pr program
}

func (p *program) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	for _, ins := range p.instructions {
		ins.encode(buf)
	}
	return buf.Bytes(), nil
}

func (i *instruction) encode(b *bytes.Buffer) {
	mgc.GetMgc(i.clear)
}

func main() {
	fmt.Println("Hello World!")
}
