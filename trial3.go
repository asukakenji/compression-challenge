package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	v0Mask3, v0Shift3 = 0x000000000000003f, 0
	v1Mask3, v1Shift3 = 0x000000000000ffc0, 6
	v2Mask3, v2Shift3 = 0x0000000003ff0000, 16
	v3Mask3, v3Shift3 = 0x0000000ffc000000, 26
	v4Mask3, v4Shift3 = 0x000000f000000000, 36
)

func encodeRecord3(v0Diff, v1, v2, v3, v4 uint64) uint64 {
	return ((v0Diff << v0Shift3) & v0Mask3) |
		((v1 << v1Shift3) & v1Mask3) |
		((v2 << v2Shift3) & v2Mask3) |
		((v3 << v3Shift3) & v3Mask3) |
		((v4 << v4Shift3) & v4Mask3)
}

func compress3() {
	var size int
	buf := make([]byte, 32)
	previousV0 := uint64(0)
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	for _, iRecord := range iRecords {
		v0Diff := iRecord.IField0 - previousV0
		r := encodeRecord3(
			v0Diff,
			iRecord.IField2-DiffMaxMinIField234,
			iRecord.IField3-DiffMaxMinIField234,
			iRecord.IField4-DiffMaxMinIField234,
			iRecord.IField5,
		)
		size = binary.PutUvarint(buf, r)
		os.Stdout.Write(buf[0:size])
		previousV0 = iRecord.IField0
	}
}

func decompress3() {
	previousV0 := uint64(0)
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	for {
		r, err := binary.ReadUvarint(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		v0 := (r&v0Mask3)>>v0Shift3 + previousV0
		field0 := stringifyTime(v0)
		field1 := field0[0:5]
		field2 := stringifyDecimal((r&v1Mask3)>>v1Shift3 + DiffMaxMinIField234)
		field3 := stringifyDecimal((r&v2Mask3)>>v2Shift3 + DiffMaxMinIField234)
		field4 := stringifyDecimal((r&v3Mask3)>>v3Shift3 + DiffMaxMinIField234)
		field5 := stringifyEnumFixed((r & v4Mask3) >> v4Shift3)
		fmt.Printf(
			"%s %s %s %s %s %s\n",
			field0,
			field1,
			field2,
			field3,
			field4,
			field5,
		)
		previousV0 = v0
	}
}
