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

func TestFindPathMixed(t *testing.T) {
	var p program
	p.grayCode = mgc.New(16)
	i := instruction{clear: 0, addImmediate: 1, addIndirect: 2, multImmediate: 3, multIndirect: 5, storeAddress: 6, storeIndirect: 7}
	p.instructions = append(p.instructions, i)
	data, err := p.MarshalBinary()
	if err != nil {
		t.Error("Failed to Marshal", data)
	}
}
