package protobufio

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

func TestReadVarint(t *testing.T) {
	buf := make([]byte, 4*binary.MaxVarintLen64)
	n1 := binary.PutVarint(buf, math.MaxInt8)
	n2 := binary.PutVarint(buf[n1:], math.MaxInt16)
	n3 := binary.PutVarint(buf[n1+n2:], math.MaxInt32)
	n4 := binary.PutVarint(buf[n1+n2+n3:], math.MaxInt64)

	r := NewVarintReader(bytes.NewBuffer(buf))
	v, n, err := r.ReadVarint()
	if err != nil {
		t.Error(err)
	}
	if n != n1 {
		t.Errorf("unexpected number of read bytes. got=%v; want=%v", n, n1)
	}
	if v != math.MaxInt8 {
		t.Errorf("unexpected value. got=%v; want=%v", v, math.MaxInt8)
	}

	v, n, err = r.ReadVarint()
	if err != nil {
		t.Error(err)
	}
	if n != n2 {
		t.Errorf("unexpected number of read bytes. got=%v; want=%v", n, n2)
	}
	if v != math.MaxInt16 {
		t.Errorf("unexpected value. got=%v; want=%v", v, math.MaxInt16)
	}

	v, n, err = r.ReadVarint()
	if err != nil {
		t.Error(err)
	}
	if n != n3 {
		t.Errorf("unexpected number of read bytes. got=%v; want=%v", n, n3)
	}
	if v != math.MaxInt32 {
		t.Errorf("unexpected value. got=%v; want=%v", v, math.MaxInt32)
	}

	v, n, err = r.ReadVarint()
	if err != nil {
		t.Error(err)
	}
	if n != n4 {
		t.Errorf("unexpected number of read bytes. got=%v; want=%v", n, n4)
	}
	if v != math.MaxInt64 {
		t.Errorf("unexpected value. got=%v; want=%v", v, math.MaxInt64)
	}
}

func TestWriteVarint(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewVarintWriter(buf)
	cases := []struct {
		v int64
		n int
	}{
		{v: math.MaxInt8, n: 2},
		{v: math.MaxInt16, n: 3},
		{v: math.MaxInt32, n: 5},
		{v: math.MaxInt64, n: 10},
	}
	for _, c := range cases {
		n, err := w.WriteVarint(c.v)
		if err != nil {
			t.Errorf("error in WriteVariant. err=%v", err)
		}
		if n != c.n {
			t.Errorf("unexpected number of written bytes. got=%v; want=%v", n, c.n)
		}
	}
	b := buf.Bytes()
	for _, c := range cases {
		v, n := binary.Varint(b)
		if n != c.n {
			t.Errorf("unexpected number of read bytes. got=%v; want=%v", n, c.n)
		}
		if v != c.v {
			t.Errorf("unexpected read integer. got=%v; want=%v", v, c.v)
		}
		b = b[n:]
	}
}
