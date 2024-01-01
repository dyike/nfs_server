package xdr

import (
	"encoding/binary"
	"io"
)

// https://datatracker.ietf.org/doc/html/rfc1057

type XDR interface {
	Serialize(writer io.Writer) error
	Deserialize(reader io.Reader) error
}

func WriteBytesAutoPad(writer io.Writer, dat []byte) error {
	if err := binary.Write(writer, binary.BigEndian, dat); err != nil {
		return err
	}
	padLen := Pad(len(dat))
	if padLen > 0 {
		pad := make([]byte, padLen)
		return binary.Write(writer, binary.BigEndian, pad)
	}
	return nil
}
