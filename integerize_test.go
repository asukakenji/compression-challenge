package main

import "testing"

func expectPanicUint64String(t *testing.T, fname string, f func(string) uint64, s string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s(%s) did not panic", fname, s)
		}
	}()
	f(s)
}

func TestIntegerizeTime(t *testing.T) {
	cases := []struct {
		time     string
		expected uint64
	}{
		{"00:00:00", 0},
		{"00:00:01", 1},
		{"00:01:00", 60},
		{"00:01:00", 60},
		{"00:01:01", 61},
		{"01:00:00", 3600},
		{"01:00:01", 3601},
		{"01:01:00", 3660},
		{"01:01:01", 3661},
		{"23:59:59", 86399},
	}
	for _, c := range cases {
		got := integerizeTime(c.time)
		if got != c.expected {
			t.Errorf("integerizeTime(%s) = %d, expected %d", c.time, got, c.expected)
		}
	}
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "hh:00:00")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "24:00:00")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "-1:00:00")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "00:mm:00")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "00:60:00")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "00:-1:00")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "00:00:ss")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "00:00:60")
	expectPanicUint64String(t, "integerizeTime", integerizeTime, "00:00:-1")
}

func TestIntegerizeDecimal(t *testing.T) {
	cases := []struct {
		decimal  string
		expected uint64
	}{
		{".", 0},
		{"0", 0},
		{"0.", 0},
		{".0", 0},
		{"0.0", 0},
		{".001", 1},
		{"0.001", 1},
		{".01", 10},
		{"0.01", 10},
		{".010", 10},
		{"0.010", 10},
		{".1", 100},
		{"0.1", 100},
		{".10", 100},
		{"0.10", 100},
		{".100", 100},
		{"0.100", 100},
		{"1", 1000},
		{"1.", 1000},
		{"1.1", 1100},
	}
	for _, c := range cases {
		got := integerizeDecimal(c.decimal)
		if got != c.expected {
			t.Errorf("integerizeDecimal(%s) = %d, expected %d", c.decimal, got, c.expected)
		}
	}
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "-1")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "-1.0")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "i")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "i.")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "0.f")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "i.f")
	expectPanicUint64String(t, "integerizeDecimal", integerizeDecimal, "0.1000")
}
