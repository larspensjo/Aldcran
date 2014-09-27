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
	"log"
	"testing"
)

func TestMutation(t *testing.T) {
	p := program{virtualMachine: vmTest}
	for i := 0; i < 20; i++ {
		p.instructions = append(p.instructions, noop)
	}
	bin, err := p.MarshalBinary()
	if err != nil {
		t.Error("MarshalBinary returned ", err)
	}
	mutate(bin, 0.03)
	p2 := program{virtualMachine: vmTest}
	p2.UnmarshalBinary(bin)
	log.Print(p2.String())
}
