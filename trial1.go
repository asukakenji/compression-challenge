package main

import (
	"encoding/binary"
	"fmt"
)

// Second Trial: 688,407 bytes

func compress1() {
	var size int
	buf := make([]byte, 32)
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	for _, iRecord := range iRecords {
		size = binary.PutUvarint(buf, iRecord.IField0)
		fmt.Printf("%s", buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField2)
		fmt.Printf("%s", buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField3)
		fmt.Printf("%s", buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField4)
		fmt.Printf("%s", buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField5)
		fmt.Printf("%s", buf[0:size])
	}
}
