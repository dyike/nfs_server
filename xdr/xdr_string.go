package xdr

import (
	"encoding/binary"
	"io"
)

type Str string

func (i *Str) Serialize(writer io.Writer) error {
	l := len(*i)
	if err := binary.Write(writer, binary.BigEndian, uint32(l)); err != nil {
		return err
	}
	dat := []byte(*i)
	return WriteBytesAutoPad(writer, dat)
}

func (i *Str) Deserialize(reader io.Reader) error {
	return nil
}
