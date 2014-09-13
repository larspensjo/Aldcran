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

func (p *program) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	for b.Len() > 0 {
		var i instruction
		err := i.decode(b, p.grayCode)
		if err != nil {
			return err
		}
		p.instructions = append(p.instructions, i)
	}
	return nil
}

// Take a binary number, convert it to mgc, and encode it into a 4-byte array
func encodeMgc(number int32, b *bytes.Buffer, m *mgc.Mgc) {
	converted := int16(m.GetMgc(number))
	err := binary.Write(b, binary.LittleEndian, converted)
	if err != nil {
		log.Fatalln("encodeMgc failed", err)
	}
}

// Given a byte string, convert to binary number
func decodeMgc(b *bytes.Buffer, m *mgc.Mgc) (int32, error) {
	var number int16
	err := binary.Read(b, binary.LittleEndian, &number)
	if err != nil {
		return 0, err
	}
	return m.GetInt(mgc.MgcNumber(number)), nil
}

func (i *instruction) encode(b *bytes.Buffer, m *mgc.Mgc) {
	encodeMgc(i.storeAddress, b, m)
	encodeMgc(i.storeIndirect, b, m)
	encodeMgc(i.clear, b, m)
	encodeMgc(i.addImmediate, b, m)
	encodeMgc(i.addIndirect, b, m)
	encodeMgc(i.multImmediate, b, m)
	encodeMgc(i.multIndirect, b, m)
}

func (i *instruction) decode(b *bytes.Buffer, m *mgc.Mgc) error {
	var err error
	// This has to be done in the opposite order to encode()

	i.multIndirect, err = decodeMgc(b, m)
	if err != nil {
		return err
	}
	i.multImmediate, err = decodeMgc(b, m)
	if err != nil {
		return err
	}
	i.clear, err = decodeMgc(b, m)
	if err != nil {
		return err
	}
	i.addImmediate, err = decodeMgc(b, m)
	if err != nil {
		return err
	}
	i.addIndirect, err = decodeMgc(b, m)
	if err != nil {
		return err
	}
	i.multImmediate, err = decodeMgc(b, m)
	if err != nil {
		return err
	}
	i.multIndirect, err = decodeMgc(b, m)
	return err
}
