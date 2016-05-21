package binaryio

import (
	"io"

	"github.com/gogo/protobuf/proto"
)

// ProtoBufWriter writes a length in the variable-length encoding
// and the following bytes in the protocol buffer encoding.
type ProtoBufWriter struct {
	writer *BytesWriter
}

// NewProtoBufWriter creates a ProtoBufWriter with the underlying
// writer.
func NewProtoBufWriter(w io.Writer) *ProtoBufWriter {
	return &ProtoBufWriter{writer: NewBytesWriter(w)}
}

// WriteVarintLenAndBytes writes a length in the variable-length encoding
// and the following encoded bytes to the underlying writer. It returns the number of
// bytes written for the length and the number of bytes written for pb.
func (w *ProtoBufWriter) WriteVarintLenAndProtoBuf(pb proto.Message) (n1, n2 int, err error) {
	buf, err := proto.Marshal(pb)
	if err != nil {
		return
	}
	n1, n2, err = w.writer.WriteVarintLenAndBytes(buf)
	return
}

// ProtoBufReader reads a length in the variable-length encoding
// and the following bytes in the protocol buffer encoding.
type ProtoBufReader struct {
	reader *BytesReader
}

// NewProtoBufReader creates a ProtoBufReader with the underlying
// reader.
func NewProtoBufReader(r io.Reader) *ProtoBufReader {
	return &ProtoBufReader{reader: NewBytesReader(r)}
}

// ReadVarintLenAndProtoBuf reads a length in the variable-length encoding
// and the following encoded bytes. It uses buf if the length of buf is large
// enough or makes a new buffer and returns it, the number of bytes
// read for the length and the number of bytes read for pb.
func (r *ProtoBufReader) ReadVarintLenAndProtoBuf(pb proto.Message, buf []byte) (bufOrNewBuf []byte, n1, n2 int, err error) {
	bufOrNewBuf, n1, n2, err = r.reader.ReadVarintLenAndBytes(buf)
	if err != nil {
		return
	}
	err = proto.Unmarshal(bufOrNewBuf[:n2], pb)
	return
}
