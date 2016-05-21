package protobufio

import (
	"io"

	"github.com/gogo/protobuf/proto"
)

// MessageWriter writes a length in the variable-length encoding
// and the following bytes in the protocol buffer encoding.
type MessageWriter struct {
	*BytesWriter
}

// NewMessageWriter creates a MessageWriter with the underlying
// writer.
func NewMessageWriter(w io.Writer) *MessageWriter {
	return &MessageWriter{BytesWriter: NewBytesWriter(w)}
}

// WriteVarintLenAndMessage writes a length in the variable-length encoding
// and the following encoded bytes to the underlying writer. It returns the number of
// bytes written for the length and the number of bytes written for pb.
func (w *MessageWriter) WriteVarintLenAndMessage(pb proto.Message) (n1, n2 int, err error) {
	buf, err := proto.Marshal(pb)
	if err != nil {
		return
	}
	n1, n2, err = w.WriteVarintLenAndBytes(buf)
	return
}

// MessageReader reads a length in the variable-length encoding
// and the following bytes in the protocol buffer encoding.
type MessageReader struct {
	*BytesReader
}

// NewMessageReader creates a MessageReader with the underlying
// reader.
func NewMessageReader(r io.Reader) *MessageReader {
	return &MessageReader{BytesReader: NewBytesReader(r)}
}

// ReadVarintLenAndMessage reads a length in the variable-length encoding
// and the following encoded bytes. It uses buf if the length of buf is large
// enough or makes a new buffer and returns it, the number of bytes
// read for the length and the number of bytes read for pb.
func (r *MessageReader) ReadVarintLenAndMessage(pb proto.Message, buf []byte) (bufOrNewBuf []byte, n1, n2 int, err error) {
	bufOrNewBuf, n1, n2, err = r.ReadVarintLenAndBytes(buf)
	if err != nil {
		return
	}
	err = proto.Unmarshal(bufOrNewBuf[:n2], pb)
	return
}
