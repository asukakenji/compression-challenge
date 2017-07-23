package main

import (
	"encoding/binary"
	"fmt"
)

// Third Trial: 561,906 bytes

const (
	v0Mask2, v0Shift2 = 0x000000000001ffff, 0
	v1Mask2, v1Shift2 = 0x000000000ffe0000, 17
	v2Mask2, v2Shift2 = 0x0000007ff0000000, 28
	v3Mask2, v3Shift2 = 0x0003ff8000000000, 39
	v4Mask2, v4Shift2 = 0x003c000000000000, 50
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
		r := encodeRecord2(
			iRecord.IField0,
			iRecord.IField2,
			iRecord.IField3,
			iRecord.IField4,
			iRecord.IField5,
		)
		size = binary.PutUvarint(buf, r)
		fmt.Printf("%s", buf[0:size])
	}
}
