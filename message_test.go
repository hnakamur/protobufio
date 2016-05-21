package protobufio

import (
	"bytes"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestReadVarintLenAndMessage(t *testing.T) {
	w := new(bytes.Buffer)
	values := []struct {
		msg ExampleMessage
		n1  int
		n2  int
	}{
		{msg: ExampleMessage{Label: "foo", Type: 1, Reps: []int64{2, 3}}},
		{msg: ExampleMessage{Label: "bar", Type: 2, Reps: nil}},
	}
	bw := NewBytesWriter(w)
	for i, value := range values {
		buf, err := proto.Marshal(&value.msg)
		if err != nil {
			t.Error(err)
		}
		n1, n2, err := bw.WriteVarintLenAndBytes(buf)
		values[i].n1 = n1
		values[i].n2 = n2
		if err != nil {
			t.Error(err)
		}
	}

	var buf []byte
	r := NewMessageReader(bytes.NewBuffer(w.Bytes()))
	for _, value := range values {
		var got ExampleMessage
		var n1, n2 int
		var err error
		buf, n1, n2, err = r.ReadVarintLenAndMessage(&got, buf)
		if err != nil {
			t.Error(err)
		}
		if n1 != value.n1 {
			t.Errorf("unexpected length. got=%v; want=%v", n1, value.n1)
		}
		if n2 != value.n2 {
			t.Errorf("unexpected length. got=%v; want=%v", n2, value.n2)
		}
		if got.Label != value.msg.Label {
			t.Errorf("unexpected value. got=%v; want=%v", got.Label, value.msg.Label)
		}
		if got.Type != value.msg.Type {
			t.Errorf("unexpected value. got=%v; want=%v", got.Type, value.msg.Type)
		}
		if len(got.Reps) != len(value.msg.Reps) {
			t.Errorf("unexpected length. got=%v; want=%v", len(got.Reps), len(value.msg.Reps))
		}
		for i, rep := range got.Reps {
			if rep != value.msg.Reps[i] {
				t.Errorf("unexpected value. got=%v; want=%v", rep, value.msg.Reps[i])
			}
		}
	}
}

func TestWriteVarintLenAndMessage(t *testing.T) {
	values := []struct {
		msg ExampleMessage
		n1  int
		n2  int
	}{
		{msg: ExampleMessage{Label: "foo", Type: 1, Reps: []int64{2, 3}}},
		{msg: ExampleMessage{Label: "bar", Type: 2, Reps: nil}},
	}

	writeBuf := new(bytes.Buffer)
	w := NewMessageWriter(writeBuf)
	for i, value := range values {
		n1, n2, err := w.WriteVarintLenAndMessage(&value.msg)
		values[i].n1 = n1
		values[i].n2 = n2
		if err != nil {
			t.Error(err)
		}
	}

	r := NewBytesReader(bytes.NewReader(writeBuf.Bytes()))
	var buf []byte
	for _, value := range values {
		var n1, n2 int
		var err error
		buf, n1, n2, err = r.ReadVarintLenAndBytes(buf)
		if err != nil {
			t.Error(err)
		}
		if value.n1 != n1 {
			t.Errorf("unexpected length. got=%v; want=%v", value.n1, n1)
		}
		if value.n2 != n2 {
			t.Errorf("unexpected length. got=%v; want=%v", value.n2, n2)
		}
		marshaledBuf, err := proto.Marshal(&value.msg)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(buf[:n2], marshaledBuf) {
			t.Errorf("unexpected value. got=%v; want=%v", buf[:n2], marshaledBuf)
		}
	}
}
