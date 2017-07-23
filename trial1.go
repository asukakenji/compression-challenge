package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func compress1() {
	var size int
	buf := make([]byte, 32)
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	for _, iRecord := range iRecords {
		iRecord.IField2 -= DiffMaxMinIField234
		iRecord.IField3 -= DiffMaxMinIField234
		iRecord.IField4 -= DiffMaxMinIField234
		size = binary.PutUvarint(buf, iRecord.IField0)
		os.Stdout.Write(buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField2)
		os.Stdout.Write(buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField3)
		os.Stdout.Write(buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField4)
		os.Stdout.Write(buf[0:size])
		size = binary.PutUvarint(buf, iRecord.IField5)
		os.Stdout.Write(buf[0:size])
	}
}

func decompress1() {
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	for {
		iField0, err := binary.ReadUvarint(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		iField2, err := binary.ReadUvarint(reader)
		if err != nil {
			panic(err)
		}
		iField3, err := binary.ReadUvarint(reader)
		if err != nil {
			panic(err)
		}
		iField4, err := binary.ReadUvarint(reader)
		if err != nil {
			panic(err)
		}
		iField5, err := binary.ReadUvarint(reader)
		if err != nil {
			panic(err)
		}
		field0 := stringifyTime(iField0)
		field1 := field0[0:5]
		field2 := stringifyDecimal(iField2 + DiffMaxMinIField234)
		field3 := stringifyDecimal(iField3 + DiffMaxMinIField234)
		field4 := stringifyDecimal(iField4 + DiffMaxMinIField234)
		field5 := stringifyEnumFixed(iField5)
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
