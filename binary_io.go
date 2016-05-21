// Package binaryio provides wrapper functions which is easier to use than in functions
// in the standard io and encoding/binary packges.
package binaryio

import (
	"encoding/binary"
	"io"
)

type nonBufferedByteReader struct {
	reader    io.Reader
	buf       []byte
	bytesRead int
}

func newNonBufferedByteReader(r io.Reader) *nonBufferedByteReader {
	return &nonBufferedByteReader{
		reader: r,
		buf:    make([]byte, 1),
	}
}

func (r *nonBufferedByteReader) ReadByte() (c byte, err error) {
	n, err := r.reader.Read(r.buf)
	r.bytesRead += n
	if n > 0 {
		c = r.buf[0]
	}
	return
}

func (r *nonBufferedByteReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	r.bytesRead += n
	return
}

func (r *nonBufferedByteReader) BytesRead() int {
	return r.bytesRead
}

// ReadVarint reads an integer in the variable-length encoding from a reader
// and returns the value read, the number of bytes read.
// See https://golang.org/pkg/encoding/binary/ for the variable-length encoding.
func ReadVarint(r io.Reader) (v int64, n int, err error) {
	br := newNonBufferedByteReader(r)
	v, err = binary.ReadVarint(br)
	return v, br.BytesRead(), err
}

// ReadVarintLenAndBytes reads a length in the variable-length encoding
// and following bytes of that length. It returns the following bytes,
// the total length of bytes read.
func ReadVarintLenAndBytes(r io.Reader) (buf []byte, n int, err error) {
	l, n, err := ReadVarint(r)
	if err != nil {
		return nil, n, err
	}
	buf = make([]byte, l)
	n2, err := io.ReadFull(r, buf)
	return buf, n + n2, err
}

// WriteVariant writes an int64 value in the variable-length encoding
// and returns the number of bytes written.
// See https://golang.org/pkg/encoding/binary/ for the variable-length encoding.
func WriteVariant(w io.Writer, v int64) (int, error) {
	buf := make([]byte, binary.Size(v))
	n := binary.PutVarint(buf, v)
	return w.Write(buf[:n])
}

// WriteFull writes all bytes in the buf to the writer w and
// return the number of bytes written.
func WriteFull(w io.Writer, buf []byte) (n int, err error) {
	for len(buf) > 0 {
		var n2 int
		n2, err = w.Write(buf)
		n += n2
		if err != nil {
			return
		}
		buf = buf[n2:]
	}
	return
}

// WriteVarintLenAndBytes writes a length of the buf in the variable-length
// encoding followed by all bytes in buf. It returns the total number of
// bytes written.
func WriteVarintLenAndBytes(w io.Writer, buf []byte) (n int, err error) {
	n, err = WriteVariant(w, int64(len(buf)))
	if err != nil {
		return
	}
	n2, err := WriteFull(w, buf)
	return n + n2, err
}
