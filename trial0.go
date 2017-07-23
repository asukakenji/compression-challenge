package main

import "fmt"

// First Trial: 738,242 bytes

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
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	for _, iRecord := range iRecords {
		r := encodeRecord0(
			iRecord.IField0,
			iRecord.IField2,
			iRecord.IField3,
			iRecord.IField4,
			iRecord.IField5,
		)
		fmt.Printf(
			"%c%c%c%c%c%c%c%c",
			(r>>56)&0xff,
			(r>>48)&0xff,
			(r>>40)&0xff,
			(r>>32)&0xff,
			(r>>24)&0xff,
			(r>>16)&0xff,
			(r>>8)&0xff,
			(r>>0)&0xff,
		)
	}
}
