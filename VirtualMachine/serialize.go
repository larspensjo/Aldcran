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
	"encoding/binary"
	"github.com/larspensjo/go-monotonic-graycode"
	"log"
)

func (p *program) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	for _, ins := range p.instructions {
		ins.encode(buf, p.grayCode)
	}
	return buf.Bytes(), nil
}

func (p *program) UnmarshalBinary(data []byte) {
	b := bytes.NewBuffer(data)
	for b.Len() > 0 {
		var i instruction
		i.decode(b, p.grayCode)
		p.instructions = append(p.instructions, i)
	}
}

// Take a binary number, convert it to mgc, and encode it into a 4-byte array
func encodeMgc(number int, b *bytes.Buffer, m *mgc.Mgc) {
	// Convert the signed number to an unsigned that can be used for MGC.
	var unsigned uint32 = uint32(number)
	if number < 0 {
		unsigned = uint32(0x10000 - number)
	}
	converted := uint16(m.GetMgc(unsigned))
	err := binary.Write(b, binary.LittleEndian, converted)
	if err != nil {
		log.Fatalln("encodeMgc failed", err)
	}
}

// Given a byte string, convert to binary number
func decodeMgc(b *bytes.Buffer, m *mgc.Mgc) int {
	var number uint16
	err := binary.Read(b, binary.LittleEndian, &number)
	if err != nil {
		log.Fatal("failed to read number from stream")
	}
	var tmp uint32 = m.GetInt(mgc.MgcNumber(number))
	// MGC only handles unsigned, convert to signed int
	ret := int(tmp)
	if tmp > 0x3FFF {
		ret = int(tmp - 0x10000)
	}
	return ret
}

func (i *instruction) encode(b *bytes.Buffer, m *mgc.Mgc) {
	encodeMgc(i.clear, b, m)
	encodeMgc(i.addImmediate, b, m)
	encodeMgc(i.addIndirect, b, m)
	encodeMgc(i.multImmediate, b, m)
	encodeMgc(i.multIndirect, b, m)
	encodeMgc(i.storeAddress, b, m)
	encodeMgc(i.storeIndirect, b, m)
}

func (i *instruction) decode(b *bytes.Buffer, m *mgc.Mgc) {
	i.clear = decodeMgc(b, m)
	i.addImmediate = decodeMgc(b, m)
	i.addIndirect = decodeMgc(b, m)
	i.multImmediate = decodeMgc(b, m)
	i.multIndirect = decodeMgc(b, m)
	i.storeAddress = decodeMgc(b, m)
	i.storeIndirect = decodeMgc(b, m)
}
