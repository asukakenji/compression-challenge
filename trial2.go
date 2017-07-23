package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	v0Mask2, v0Shift2 = 0x000000000001ffff, 0
	v1Mask2, v1Shift2 = 0x0000000007fe0000, 17
	v2Mask2, v2Shift2 = 0x0000001ff8000000, 27
	v3Mask2, v3Shift2 = 0x00007fe000000000, 37
	v4Mask2, v4Shift2 = 0x0007800000000000, 47
)

func encodeRecord2(v0, v1, v2, v3, v4 uint64) uint64 {
	return ((v0 << v0Shift2) & v0Mask2) |
		((v1 << v1Shift2) & v1Mask2) |
		((v2 << v2Shift2) & v2Mask2) |
		((v3 << v3Shift2) & v3Mask2) |
		((v4 << v4Shift2) & v4Mask2)
}

func compress2() {
	var size int
	buf := make([]byte, 32)
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	for _, iRecord := range iRecords {
		iRecord.IField2 -= DiffMaxMinIField234
		iRecord.IField3 -= DiffMaxMinIField234
		iRecord.IField4 -= DiffMaxMinIField234
		r := encodeRecord2(
			iRecord.IField0,
			iRecord.IField2,
			iRecord.IField3,
			iRecord.IField4,
			iRecord.IField5,
		)
		size = binary.PutUvarint(buf, r)
		os.Stdout.Write(buf[0:size])
	}
}

func decompress2() {
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	for {
		r, err := binary.ReadUvarint(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		field0 := stringifyTime((r & v0Mask2) >> v0Shift2)
		field1 := field0[0:5]
		field2 := stringifyDecimal((r&v1Mask2)>>v1Shift2 + DiffMaxMinIField234)
		field3 := stringifyDecimal((r&v2Mask2)>>v2Shift2 + DiffMaxMinIField234)
		field4 := stringifyDecimal((r&v3Mask2)>>v3Shift2 + DiffMaxMinIField234)
		field5 := stringifyEnumFixed((r & v4Mask2) >> v4Shift2)
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
