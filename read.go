package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Record struct {
	Field0 string
	//Field1 string
	Field2 string
	Field3 string
	Field4 string
	Field5 string
}

func nextField(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return scanner.Text()
	}
	panic("Call to Scan() failed")
}

func readRecords() []Record {
	lineNo := 1
	records := []Record{}
	lineScanner := bufio.NewScanner(os.Stdin)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		field0 := nextField(wordScanner)
		field1 := nextField(wordScanner)
		field2 := nextField(wordScanner)
		field3 := nextField(wordScanner)
		field4 := nextField(wordScanner)
		field5 := nextField(wordScanner)
		if field1 != field0[0:5] {
			err := fmt.Errorf("Unexpected field1 at line #%d: field0: %s, field1: %s", lineNo, field0, field1)
			panic(err)
		}
		records = append(records, Record{
			field0,
			field2,
			field3,
			field4,
			field5,
		})
		lineNo++
	}
	return records
}
