package counter

import (
	"io"
	"testing"
)

func TestWriteCounter(t *testing.T) {
	counter := NewWriteCounter(io.Discard)

	data := []byte("Hello NFS")
	counter.Write(data)

	t.Log("Bytes written: ", counter.BytesWritten())
}
