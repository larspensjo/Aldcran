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
	"log"
	"testing"
)

func TestMutation(t *testing.T) {
	var vm VirtualMachine
	vm.graycode = mgc.New(16)
	p := program{virtualMachine: &vm}
	for i := 0; i < 10; i++ {
		p.instructions = append(p.instructions, noop)
	}
	log.Print(p.String())
	_, err := p.MarshalBinary()
	if err != nil {
		t.Error("MarshalBinary returned ", err)
	}
	// mutate(bin, 0.1)
}
