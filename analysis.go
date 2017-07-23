package main

import "fmt"

type RecordAnalysis struct {
	// MinDiffIField0 is the minimum difference,
	// counting all differences between field 0 of consecutive lines (in seconds).
	MinDiffIField0 uint64
	// MaxDiffIField0 is the maximum difference,
	// counting all differences between field 0 of consecutive lines (in seconds).
	MaxDiffIField0 uint64

	// MinIField234 is the minimum value,
	// counting field 2, field 3, field 4 of all records (in 0.001).
	MinIField234 uint64
	// MaxIField234 is the maximum value,
	// counting field 2, field 3, field 4 of all records (in 0.001).
	MaxIField234 uint64

	// The frequency distribution of field 0.
	// The key is the value of field 0, the value is the frequency.
	IField0FD map[uint64]uint64
	// The frequency distribution of diff of field 0.
	// The key is the diff of field 0, the value is the frequency.
	DiffIField0FD map[uint64]uint64

	// The frequency distribution of field 5.
	// The key is the value of field 5, the value is the frequency.
	IField5FD map[uint64]uint64
}

func analyzeRecords(records []IntegerizedRecord) *RecordAnalysis {
	minDiffIField0 := ^uint64(0)
	maxDiffIField0 := uint64(0)

	minIField234 := ^uint64(0)
	maxIField234 := uint64(0)

	iField0FD := map[uint64]uint64{}
	diffIField0FD := map[uint64]uint64{}

	iField5FD := map[uint64]uint64{}

	previousIField0 := uint64(0)
	for lineNo, record := range records {
		if record.IField0 < previousIField0 {
			panic(fmt.Errorf("Time at line #%d is earlier than previous line", lineNo))
		}
		diffIField0 := record.IField0 - previousIField0
		if diffIField0 < minDiffIField0 {
			minDiffIField0 = diffIField0
		}
		if diffIField0 > maxDiffIField0 {
			maxDiffIField0 = diffIField0
		}
		iField0FD[record.IField0]++
		diffIField0FD[diffIField0]++
		previousIField0 = record.IField0

		if record.IField2 < minIField234 {
			minIField234 = record.IField2
		}
		if record.IField3 < minIField234 {
			minIField234 = record.IField3
		}
		if record.IField4 < minIField234 {
			minIField234 = record.IField4
		}
		if record.IField2 > maxIField234 {
			maxIField234 = record.IField2
		}
		if record.IField3 > maxIField234 {
			maxIField234 = record.IField3
		}
		if record.IField4 > maxIField234 {
			maxIField234 = record.IField4
		}

		iField5FD[record.IField5]++
	}
	return &RecordAnalysis{
		MinDiffIField0: minDiffIField0,
		MaxDiffIField0: maxDiffIField0,
		MinIField234:   minIField234,
		MaxIField234:   maxIField234,
		IField0FD:      iField0FD,
		DiffIField0FD:  diffIField0FD,
		IField5FD:      iField5FD,
	}
}
