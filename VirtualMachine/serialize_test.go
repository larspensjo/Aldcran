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
	"testing"
)

func TestSerialization(t *testing.T) {
	var p program
	p.grayCode = mgc.New(16)
	i := instruction{clear: 1, addImmediate: 2, addIndirect: 3, multImmediate: 4, multIndirect: 5, storeAddress: 6, storeIndirect: 7}
	p.instructions = append(p.instructions, i)
	data, err := p.MarshalBinary()
	if err != nil {
		t.Error("Failed to Marshal", data)
	}
	var p2 program
	p2.grayCode = p.grayCode
	p2.UnmarshalBinary(data)
	if len(p2.instructions) != 1 {
		t.Error("Expected one instruction")
	}
	if p2.instructions[0] != i {
		t.Error("Failed to serialize/deserialize:", i, p2.instructions[0])
	}

}
