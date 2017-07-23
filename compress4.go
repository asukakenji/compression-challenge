package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

func compress4() {
	previousV0 := uint64(0)
	records := readRecords()
	iRecords := integerizeRecords(integerizeEnumFixed, records)
	w := NewBitWriter()
	for _, iRecord := range iRecords {
		w.WriteUint64(iRecord.IField0-previousV0, 6)
		w.WriteUint64(iRecord.IField2-DiffMaxMinIField234, 10)
		w.WriteUint64(iRecord.IField3-DiffMaxMinIField234, 10)
		w.WriteUint64(iRecord.IField4-DiffMaxMinIField234, 10)
		w.WriteUint64(iRecord.IField5, 4)
		previousV0 = iRecord.IField0
	}
	buf := make([]byte, 32)
	size := binary.PutUvarint(buf, uint64(len(records)))
	os.Stdout.Write(buf[0:size])
	os.Stdout.Write(w.Data)
}

func decompress4() {
	previousV0 := uint64(0)
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	maxLineNo, err := binary.ReadUvarint(reader)
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	r := NewBitReader(buf)
	for lineNo := uint64(0); lineNo < maxLineNo; lineNo++ {
		v0 := r.ReadUint64(6) + previousV0
		field0 := stringifyTime(v0)
		field1 := field0[0:5]
		field2 := stringifyDecimal(r.ReadUint64(10) + DiffMaxMinIField234)
		field3 := stringifyDecimal(r.ReadUint64(10) + DiffMaxMinIField234)
		field4 := stringifyDecimal(r.ReadUint64(10) + DiffMaxMinIField234)
		field5 := stringifyEnumFixed(r.ReadUint64(4))
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
