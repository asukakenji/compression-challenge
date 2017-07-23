package main

import (
	"encoding/binary"
	"fmt"
)

// Fourth Trial: 439,681 bytes

const (
	v0Mask3, v0Shift3 = 0x000000000000003f, 0
	v1Mask3, v1Shift3 = 0x000000000001ffc0, 6
	v2Mask3, v2Shift3 = 0x000000000ffe0000, 17
	v3Mask3, v3Shift3 = 0x0000007ff0000000, 28
	v4Mask3, v4Shift3 = 0x0000078000000000, 39
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
			iRecord.IField2,
			iRecord.IField3,
			iRecord.IField4,
			iRecord.IField5,
		)
		size = binary.PutUvarint(buf, r)
		fmt.Printf("%s", buf[0:size])
		previousV0 = iRecord.IField0
	}
}
