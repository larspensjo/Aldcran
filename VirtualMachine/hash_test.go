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
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"testing"
)

const size = 1e6

func BenchmarkAdler32(t *testing.B) {
	vector := make([]byte, size)
	hash := adler32.New()
	for i := 0; i < t.N; i++ {
		hash.Write(vector)
		hash.Sum32()
	}
}

func BenchmarkCrc32(t *testing.B) {
	vector := make([]byte, size)
	hash := crc32.NewIEEE()
	for i := 0; i < t.N; i++ {
		hash.Write(vector)
		hash.Sum32()
	}
}

func BenchmarkFnv32(t *testing.B) {
	vector := make([]byte, size)
	hash := fnv.New32()
	for i := 0; i < t.N; i++ {
		hash.Write(vector)
		hash.Sum32()
	}
}
