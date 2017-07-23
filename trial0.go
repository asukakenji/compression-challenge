package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	v0Mask0, v0Shift0 = 0x00000000000fffff, 0
	v1Mask0, v1Shift0 = 0x00000000fff00000, 20
	v2Mask0, v2Shift0 = 0x00000fff00000000, 32
	v3Mask0, v3Shift0 = 0x00fff00000000000, 44
	v4Mask0, v4Shift0 = 0x0f00000000000000, 56
)

func encodeRecord0(v0, v1, v2, v3, v4 uint64) uint64 {
	return ((v0 << v0Shift0) & v0Mask0) |
		((v1 << v1Shift0) & v1Mask0) |
		((v2 << v2Shift0) & v2Mask0) |
		((v3 << v3Shift0) & v3Mask0) |
		((v4 << v4Shift0) & v4Mask0)
}

func compress0() {
	buf := make([]byte, 8)
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	for _, iRecord := range iRecords {
		r := encodeRecord0(
			iRecord.IField0,
			iRecord.IField2-DiffMaxMinIField234,
			iRecord.IField3-DiffMaxMinIField234,
			iRecord.IField4-DiffMaxMinIField234,
			iRecord.IField5,
		)
		buf[0] = byte((r >> 56) & 0xff)
		buf[1] = byte((r >> 48) & 0xff)
		buf[2] = byte((r >> 40) & 0xff)
		buf[3] = byte((r >> 32) & 0xff)
		buf[4] = byte((r >> 24) & 0xff)
		buf[5] = byte((r >> 16) & 0xff)
		buf[6] = byte((r >> 8) & 0xff)
		buf[7] = byte((r >> 0) & 0xff)
		os.Stdout.Write(buf)
	}
}

func decompress0() {
	buf := make([]byte, 8)
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if n != 8 {
			panic("n != 8")
		}
		r := uint64(buf[0])<<56 |
			uint64(buf[1])<<48 |
			uint64(buf[2])<<40 |
			uint64(buf[3])<<32 |
			uint64(buf[4])<<24 |
			uint64(buf[5])<<16 |
			uint64(buf[6])<<8 |
			uint64(buf[7])<<0
		field0 := stringifyTime((r & v0Mask0) >> v0Shift0)
		field1 := field0[0:5]
		field2 := stringifyDecimal((r&v1Mask0)>>v1Shift0 + DiffMaxMinIField234)
		field3 := stringifyDecimal((r&v2Mask0)>>v2Shift0 + DiffMaxMinIField234)
		field4 := stringifyDecimal((r&v3Mask0)>>v3Shift0 + DiffMaxMinIField234)
		field5 := stringifyEnumFixed((r & v4Mask0) >> v4Shift0)
		fmt.Printf(
			"%s %s %s %s %s %s\n",
			field0,
			field1,
			field2,
			field3,
			field4,
			field5,
		)
	}
}
