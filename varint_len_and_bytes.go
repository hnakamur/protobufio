package protobufio

import "io"

// BytesWriter writes a length in the variable-length encoding
// and the following bytes.
type BytesWriter struct {
	io.Writer
}

// NewBytesWriter creates a BytesWriter with the undelying writer.
func NewBytesWriter(w io.Writer) *BytesWriter {
	return &BytesWriter{Writer: w}
}

// WriteVarintLenAndBytes writes a length in the variable-length encoding
// and the following bytes in buf to the underlying writer. It returns the number of
// bytes written for the length and the number of bytes written for bytes in buf.
func (w *BytesWriter) WriteVarintLenAndBytes(buf []byte) (n1, n2 int, err error) {
	n1, err = NewVarintWriter(w).WriteVarint(int64(len(buf)))
	if err != nil {
		return
	}
	n2, err = w.Write(buf)
	return
}

// BytesReader reads a length in the variable-length encoding
// and the following bytes.
type BytesReader struct {
	io.Reader
}

// NewBytesReader creates a BytesReader with the underlying reader.
func NewBytesReader(r io.Reader) *BytesReader {
	return &BytesReader{Reader: r}
}

// ReadVarintLenAndBytes reads a length in the variable-length encoding
// and the following bytes. It uses buf if the length of buf is large
// enough or makes a new buffer and returns it, the number of bytes
// read for the length and the number of bytes read for the following
// bytes. You can get the following bytes with bufOrNewBuf[:n2].
func (r *BytesReader) ReadVarintLenAndBytes(buf []byte) (bufOrNewBuf []byte, n1, n2 int, err error) {
	length, n1, err := NewVarintReader(r).ReadVarint()
	if err != nil {
		return
	}
	if length > int64(len(buf)) {
		bufOrNewBuf = make([]byte, length)
	} else {
		bufOrNewBuf = buf
	}
	n2, err = io.ReadFull(r, bufOrNewBuf[:length])
	return
}
