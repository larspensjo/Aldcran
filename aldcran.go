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

package main

import (
	"Aldcran/MergePrograms"
	"Aldcran/VirtualMachine"
	"fmt"
)

var virtualMachine = vm.New(16, 10)

func main() {
	merge.Test()
	fmt.Println(virtualMachine)
}
