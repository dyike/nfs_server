package xdr

import "io"

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(data []byte) (int, error) {
	return w.w.Write(data)
}
