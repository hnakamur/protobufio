// Package binaryio provides wrapper functions which is easier to use than in functions
// in the standard io and encoding/binary packges.
package binaryio

import (
	"encoding/binary"
	"io"
)

// VarintReader is a reader for reading integers in the
// variant-length encoding from the underlying reader.
// See https://golang.org/pkg/encoding/binary/ for the
// variable-length encoding.
type VarintReader struct {
	reader io.Reader
	buf    [1]byte
}

// NewVarintReader creates a VarintReader with the
// underlying reader.
func NewVarintReader(r io.Reader) *VarintReader {
	return &VarintReader{reader: r}
}

// ReadByte read a byte from the underlying reader
// and the byte read.
func (r *VarintReader) ReadByte() (c byte, err error) {
	n, err := r.reader.Read(r.buf[:])
	if n > 0 {
		c = r.buf[0]
	}
	return
}

// ReadVarint reads and returns an integer from the
// underlying reader.
func (r *VarintReader) ReadVarint() (v int64, err error) {
	return binary.ReadVarint(r)
}

// VarintWriter is a writer for writing integers in the
// variant-length encoding to the underlying writer.
// See https://golang.org/pkg/encoding/binary/ for the
// variable-length encoding.
type VarintWriter struct {
	writer io.Writer
	buf    [binary.MaxVarintLen64]byte
}

// NewVarintWriter creates a VarintWriter with the
// underlying writer.
func NewVarintWriter(w io.Writer) *VarintWriter {
	return &VarintWriter{writer: w}
}

// WriteVarint writes an integer in the variable-length
// encoding to the underlying writer. It returns the
// number of bytes written.
func (w *VarintWriter) WriteVarint(v int64) (n int, err error) {
	n = binary.PutVarint(w.buf[:], v)
	return WriteFull(w.writer, w.buf[:n])
}

// WriteFull writes all bytes in the buf to the writer w and
// return the number of bytes written.
func WriteFull(w io.Writer, buf []byte) (n int, err error) {
	for n < len(buf) {
		var n2 int
		n2, err = w.Write(buf[n:])
		n += n2
		if err != nil {
			return
		}
	}
	return
}
