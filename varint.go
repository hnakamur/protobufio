package protobufio

import (
	"encoding/binary"
	"io"
)

// VarintReader is a reader for reading integers in the
// variant-length encoding from the underlying reader.
// See https://golang.org/pkg/encoding/binary/ for the
// variable-length encoding.
type VarintReader struct {
	io.Reader
	buf       [1]byte
	bytesRead int
}

// NewVarintReader creates a VarintReader with the
// underlying reader.
func NewVarintReader(r io.Reader) *VarintReader {
	return &VarintReader{Reader: r}
}

// ReadByte read a byte from the underlying reader
// and the byte read.
func (r *VarintReader) ReadByte() (c byte, err error) {
	n, err := r.Read(r.buf[:])
	if n > 0 {
		c = r.buf[0]
		r.bytesRead++
	}
	return
}

// ReadVarint reads and returns an integer from the
// underlying reader and the number of bytes read.
func (r *VarintReader) ReadVarint() (v int64, n int, err error) {
	r.bytesRead = 0
	v, err = binary.ReadVarint(r)
	n = r.bytesRead
	return
}

// VarintWriter is a writer for writing integers in the
// variant-length encoding to the underlying writer.
// See https://golang.org/pkg/encoding/binary/ for the
// variable-length encoding.
type VarintWriter struct {
	io.Writer
	buf [binary.MaxVarintLen64]byte
}

// NewVarintWriter creates a VarintWriter with the
// underlying writer.
func NewVarintWriter(w io.Writer) *VarintWriter {
	return &VarintWriter{Writer: w}
}

// WriteVarint writes an integer in the variable-length
// encoding to the underlying writer. It returns the
// number of bytes written.
func (w *VarintWriter) WriteVarint(v int64) (n int, err error) {
	n = binary.PutVarint(w.buf[:], v)
	return w.Write(w.buf[:n])
}
