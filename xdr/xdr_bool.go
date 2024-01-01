package xdr

import (
	"encoding/binary"
	"io"
)

type Bool bool

func (b *Bool) Serialize(writer io.Writer) error {
	var val uint32
	if *b {
		val = 1
	}
	return binary.Write(writer, binary.BigEndian, val)
}

func (b *Bool) Deserialize(reader io.Reader) error {
	var val uint32
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return err
	}
	*b = val > 0
	return nil
}
