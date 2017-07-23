package main

import (
	"reflect"
	"testing"
)

// Clone is a helper method for cloning w for testing purposes.
func (w *BitWriter) Clone() *BitWriter {
	data := make([]byte, len(w.Data))
	copy(data, w.Data)
	return &BitWriter{
		w.ByteIndex,
		w.BitIndex,
		data,
	}
}

// Equals is a helper method for comparing equality between w and x for testing purposes.
func (w *BitWriter) Equals(x *BitWriter) bool {
	return w.ByteIndex == x.ByteIndex &&
		w.BitIndex == x.BitIndex &&
		reflect.DeepEqual(w.Data, x.Data)
}

func TestBitWriter_WriteBool_New(t *testing.T) {
	cases := []struct {
		data     bool
		expected *BitWriter
	}{
		{false, &BitWriter{0, 1, []byte{0x00}}},
		{true, &BitWriter{0, 1, []byte{0x80}}},
	}
	for _, c := range cases {
		got := NewBitWriter()
		got.WriteBool(c.data)
		if !got.Equals(c.expected) {
			t.Errorf("NewBitWriter().WriteBool(%t) = %v, expected %v", c.data, got, c.expected)
		}
	}
}

func TestBitWriter_WriteBool_Continue_1(t *testing.T) {
	cases := []struct {
		data     bool
		expected *BitWriter
	}{
		{true, &BitWriter{0, 1, []byte{0x80}}},
		{true, &BitWriter{0, 2, []byte{0xc0}}},
		{true, &BitWriter{0, 3, []byte{0xe0}}},
		{true, &BitWriter{0, 4, []byte{0xf0}}},
		{true, &BitWriter{0, 5, []byte{0xf8}}},
		{true, &BitWriter{0, 6, []byte{0xfc}}},
		{true, &BitWriter{0, 7, []byte{0xfe}}},
		{true, &BitWriter{0, 8, []byte{0xff}}},
		{true, &BitWriter{1, 1, []byte{0xff, 0x80}}},
		{true, &BitWriter{1, 2, []byte{0xff, 0xc0}}},
		{true, &BitWriter{1, 3, []byte{0xff, 0xe0}}},
		{true, &BitWriter{1, 4, []byte{0xff, 0xf0}}},
		{true, &BitWriter{1, 5, []byte{0xff, 0xf8}}},
		{true, &BitWriter{1, 6, []byte{0xff, 0xfc}}},
		{true, &BitWriter{1, 7, []byte{0xff, 0xfe}}},
		{true, &BitWriter{1, 8, []byte{0xff, 0xff}}},
		{true, &BitWriter{2, 1, []byte{0xff, 0xff, 0x80}}},
	}
	got := NewBitWriter()
	for _, c := range cases {
		prev := got.Clone()
		got.WriteBool(c.data)
		if !got.Equals(c.expected) {
			t.Errorf("%v.WriteBool(%t) = %v, expected %v", prev, c.data, got, c.expected)
			t.Errorf("adasd")
		}
	}
}

func TestBitWriter_WriteBool_Continue_2(t *testing.T) {
	cases := []struct {
		data     bool
		expected *BitWriter
	}{
		{false, &BitWriter{0, 1, []byte{0x00}}},
		{false, &BitWriter{0, 2, []byte{0x00}}},
		{false, &BitWriter{0, 3, []byte{0x00}}},
		{true, &BitWriter{0, 4, []byte{0x10}}},
		{true, &BitWriter{0, 5, []byte{0x18}}},
		{false, &BitWriter{0, 6, []byte{0x18}}},
		{true, &BitWriter{0, 7, []byte{0x1a}}},
		{true, &BitWriter{0, 8, []byte{0x1b}}},
		{true, &BitWriter{1, 1, []byte{0x1b, 0x80}}},
		{true, &BitWriter{1, 2, []byte{0x1b, 0xc0}}},
		{false, &BitWriter{1, 3, []byte{0x1b, 0xc0}}},
		{true, &BitWriter{1, 4, []byte{0x1b, 0xd0}}},
		{true, &BitWriter{1, 5, []byte{0x1b, 0xd8}}},
		{false, &BitWriter{1, 6, []byte{0x1b, 0xd8}}},
		{false, &BitWriter{1, 7, []byte{0x1b, 0xd8}}},
		{false, &BitWriter{1, 8, []byte{0x1b, 0xd8}}},
		{true, &BitWriter{2, 1, []byte{0x1b, 0xd8, 0x80}}},
	}
	got := NewBitWriter()
	for _, c := range cases {
		prev := got.Clone()
		got.WriteBool(c.data)
		if !got.Equals(c.expected) {
			t.Errorf("%v.WriteBool(%t) = %v, expected %v", prev, c.data, got, c.expected)
			t.Errorf("adasd")
		}
	}
}

func TestBitWriter_WriteUint64_New(t *testing.T) {
	cases := []struct {
		data     uint64
		bitCount uint
		expected []byte
	}{
		{0, 0, []byte{}},
		{0, 1, []byte{0x00}},
		{0, 2, []byte{0x00}},
		{0, 3, []byte{0x00}},
		{0, 4, []byte{0x00}},
		{0, 5, []byte{0x00}},
		{0, 6, []byte{0x00}},
		{0, 7, []byte{0x00}},
		{0, 8, []byte{0x00}},
		{0, 9, []byte{0x00, 0x00}},
		{0, 10, []byte{0x00, 0x00}},
		{0, 11, []byte{0x00, 0x00}},
		{0, 12, []byte{0x00, 0x00}},
		{0, 13, []byte{0x00, 0x00}},
		{0, 14, []byte{0x00, 0x00}},
		{0, 15, []byte{0x00, 0x00}},
		{0, 16, []byte{0x00, 0x00}},
		{0, 17, []byte{0x00, 0x00, 0x00}},
		{1, 0, []byte{}},
		{1, 1, []byte{0x80}},
		{1, 2, []byte{0x40}},
		{1, 3, []byte{0x20}},
		{1, 4, []byte{0x10}},
		{1, 5, []byte{0x08}},
		{1, 6, []byte{0x04}},
		{1, 7, []byte{0x02}},
		{1, 8, []byte{0x01}},
		{1, 9, []byte{0x00, 0x80}},
		{1, 10, []byte{0x00, 0x40}},
		{1, 11, []byte{0x00, 0x20}},
		{1, 12, []byte{0x00, 0x10}},
		{1, 13, []byte{0x00, 0x08}},
		{1, 14, []byte{0x00, 0x04}},
		{1, 15, []byte{0x00, 0x02}},
		{1, 16, []byte{0x00, 0x01}},
		{1, 17, []byte{0x00, 0x00, 0x80}},
		{2, 0, []byte{}},
		{2, 1, []byte{0x00}},
		{2, 2, []byte{0x80}},
		{2, 3, []byte{0x40}},
		{2, 4, []byte{0x20}},
		{2, 5, []byte{0x10}},
		{2, 6, []byte{0x08}},
		{2, 7, []byte{0x04}},
		{2, 8, []byte{0x02}},
		{2, 9, []byte{0x01, 0x00}},
		{2, 10, []byte{0x00, 0x80}},
		{2, 11, []byte{0x00, 0x40}},
		{2, 12, []byte{0x00, 0x20}},
		{2, 13, []byte{0x00, 0x10}},
		{2, 14, []byte{0x00, 0x08}},
		{2, 15, []byte{0x00, 0x04}},
		{2, 16, []byte{0x00, 0x02}},
		{2, 17, []byte{0x00, 0x01, 0x00}},
		{3, 0, []byte{}},
		{3, 1, []byte{0x80}},
		{3, 2, []byte{0xc0}},
		{3, 3, []byte{0x60}},
		{3, 4, []byte{0x30}},
		{3, 5, []byte{0x18}},
		{3, 6, []byte{0x0c}},
		{3, 7, []byte{0x06}},
		{3, 8, []byte{0x03}},
		{3, 9, []byte{0x01, 0x80}},
		{3, 10, []byte{0x00, 0xc0}},
		{3, 11, []byte{0x00, 0x60}},
		{3, 12, []byte{0x00, 0x30}},
		{3, 13, []byte{0x00, 0x18}},
		{3, 14, []byte{0x00, 0x0c}},
		{3, 15, []byte{0x00, 0x06}},
		{3, 16, []byte{0x00, 0x03}},
		{3, 17, []byte{0x00, 0x01, 0x80}},
	}
	for _, c := range cases {
		w := NewBitWriter()
		w.WriteUint64(c.data, c.bitCount)
		got := w.Data
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("NewBitWriter().WriteUint64(%d, %d).Data() = %v, expected %v", c.data, c.bitCount, got, c.expected)
		}
	}
}

func TestBitWriter_WriteUint64_Continue_1(t *testing.T) {
	cases := []struct {
		data     uint64
		bitCount uint
		expected []byte
	}{
		{1, 1, []byte{0x80}},
		{1, 1, []byte{0xc0}},
		{1, 1, []byte{0xe0}},
		{1, 1, []byte{0xf0}},
		{1, 1, []byte{0xf8}},
		{1, 1, []byte{0xfc}},
		{1, 1, []byte{0xfe}},
		{1, 1, []byte{0xff}},
		{1, 1, []byte{0xff, 0x80}},
		{1, 1, []byte{0xff, 0xc0}},
		{1, 1, []byte{0xff, 0xe0}},
		{1, 1, []byte{0xff, 0xf0}},
		{1, 1, []byte{0xff, 0xf8}},
		{1, 1, []byte{0xff, 0xfc}},
		{1, 1, []byte{0xff, 0xfe}},
		{1, 1, []byte{0xff, 0xff}},
	}
	w := NewBitWriter()
	for i, c := range cases {
		w.WriteUint64(c.data, c.bitCount)
		got := w.Data
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("Case #%d: %#v.WriteUint64(%d, %d).Data() = %v, expected %v", i, w, c.data, c.bitCount, got, c.expected)
		}
	}
}
