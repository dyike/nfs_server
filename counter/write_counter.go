package counter

import "io"

type WriteCounter struct {
	inner io.Writer
	count int
}

func NewWriteCounter(inner io.Writer) *WriteCounter {
	return &WriteCounter{inner: inner, count: 0}
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n, err := wc.inner.Write(p)
	wc.count += n
	return n, err
}

func (wc *WriteCounter) BytesWritten() int {
	return wc.count
}
