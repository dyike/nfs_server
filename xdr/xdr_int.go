package xdr

import (
	"encoding/binary"
	"io"
)

// I32,I64,U32,U64

// i32 类型的序列化和反序列化
type I32 int32

func (i *I32) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.BigEndian, int32(*i))
}

func (i *I32) Deserialize(reader io.Reader) error {
	var val int32
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return err
	}
	*i = I32(val)
	return nil
}

// i64 类型的序列化和反序列化
type I64 int64

func (i *I64) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.BigEndian, int64(*i))
}

func (i *I64) Deserialize(reader io.Reader) error {
	var val int64
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return err
	}
	*i = I64(val)
	return nil
}

// u32 类型的序列化和反序列化
type U32 uint32

func (u *U32) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.BigEndian, uint32(*u))
}

func (u *U32) Deserialize(reader io.Reader) error {
	var val uint32
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return err
	}
	*u = U32(val)
	return nil
}

// u64 类型的序列化和反序列化
type U64 uint64

func (u *U64) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.BigEndian, uint64(*u))
}

func (u *U64) Deserialize(reader io.Reader) error {
	var val uint64
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return err
	}
	*u = U64(val)
	return nil
}
