package protobufio

import (
	"bytes"
	"io"
	"testing"
)

func TestReadVarintLenAndBytes(t *testing.T) {
	w := new(bytes.Buffer)
	values := []struct {
		text string
		n1   int
		n2   int
	}{
		{text: "hello, world."},
		{text: "goodbye, world."},
	}
	vw := NewVarintWriter(w)
	for i, value := range values {
		n1, err := vw.WriteVarint(int64(len(value.text)))
		values[i].n1 = n1
		if err != nil {
			t.Error(err)
		}
		n2, err := w.Write([]byte(value.text))
		values[i].n2 = n2
		if err != nil {
			t.Error(err)
		}
	}

	var buf []byte
	r := NewBytesReader(bytes.NewBuffer(w.Bytes()))
	for _, value := range values {
		buf, n1, n2, err := r.ReadVarintLenAndBytes(buf)
		if err != nil {
			t.Error(err)
		}
		if n1 != value.n1 {
			t.Errorf("unexpected length. got=%v; want=%v", n1, value.n1)
		}
		if n2 != value.n2 {
			t.Errorf("unexpected length. got=%v; want=%v", n2, value.n2)
		}
		got := string(buf[:n2])
		if got != value.text {
			t.Errorf("unexpected value. got=%v; want=%v", got, value.text)
		}
	}
}

func TestWriteVarintLenAndBytes(t *testing.T) {
	writeBuf := new(bytes.Buffer)
	w := NewBytesWriter(writeBuf)
	values := []struct {
		text string
		n1   int
		n2   int
	}{
		{text: "hello, world."},
		{text: "goodbye, world."},
	}
	for i, value := range values {
		n1, n2, err := w.WriteVarintLenAndBytes([]byte(value.text))
		values[i].n1 = n1
		values[i].n2 = n2
		if err != nil {
			t.Error(err)
		}
	}

	r := bytes.NewReader(writeBuf.Bytes())
	vr := NewVarintReader(r)
	var buf []byte
	for _, value := range values {
		l, n1, err := vr.ReadVarint()
		if err != nil {
			t.Error(err)
		}
		if value.n1 != n1 {
			t.Errorf("unexpected length. got=%v; want=%v", value.n1, n1)
		}
		if l > int64(len(buf)) {
			buf = make([]byte, l)
		}
		n2, err := io.ReadFull(r, buf)
		if err != nil {
			t.Error(err)
		}
		if value.n2 != n2 {
			t.Errorf("unexpected length. got=%v; want=%v", value.n2, n2)
		}
		got := string(buf[:n2])
		if got != value.text {
			t.Errorf("unexpected value. got=%v; want=%v", got, value.text)
		}
	}
}
