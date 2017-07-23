package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntegerizedRecord struct {
	IField0 uint64
	//IField1 uint64
	IField2 uint64
	IField3 uint64
	IField4 uint64
	IField5 uint64
}

// integerizeTime converts a time string into an integer.
// time in the format: "HH:MM:SS", where
//     HH is the hour and 00 <= HH <= 23,
//     MM is the minute and 00 <= MM <= 59,
//     SS is the second and 00 <= SS <= 59.
// 0 <= return value < 86,400
func integerizeTime(time string) uint64 {
	hour := time[0:2]
	minute := time[3:5]
	second := time[6:8]
	hh, err := strconv.Atoi(hour)
	if err != nil {
		panic(err)
	}
	if hh < 0 || hh > 23 {
		panic(fmt.Errorf("invalid hour: %s", hour))
	}
	mm, err := strconv.Atoi(minute)
	if err != nil {
		panic(err)
	}
	if mm < 0 || mm > 59 {
		panic(fmt.Errorf("invalid minute: %s", minute))
	}
	ss, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}
	if ss < 0 || ss > 59 {
		panic(fmt.Errorf("invalid second: %s", second))
	}
	return uint64(hh)*3600 + uint64(mm)*60 + uint64(ss)
}

// integerizeDecimal converts a decimal string into an integer.
// decimal in the format: "[[i]i]i[.f[f[f]]]", where
// i is an integral digit, and there are at most 3 digits,
// the decimal point is optional,
// f is a fractional digit, and there are at most 3 digits.
// 0 <= return value < 1,000,000
func integerizeDecimal(decimal string) uint64 {
	dotIndex := strings.Index(decimal, ".")
	var i, f int
	var err error
	if dotIndex == -1 {
		i, err = strconv.Atoi(decimal)
		if err != nil {
			panic(err)
		}
		if i < 0 || 1 >= 1000 {
			panic(fmt.Errorf("invalid integral part: %s", decimal))
		}
		f = 0
	} else {
		integral := decimal[0:dotIndex]
		if integral == "" {
			i = 0
		} else {
			i, err = strconv.Atoi(integral)
			if err != nil {
				panic(err)
			}
			if i < 0 || i >= 1000 {
				panic(fmt.Errorf("invalid integral part: %s", integral))
			}
		}
		if dotIndex == len(decimal)-1 {
			f = 0
		} else {
			fractional := decimal[dotIndex+1:]
			for len(fractional) < 3 {
				fractional += "0"
			}
			f, err = strconv.Atoi(fractional)
			if err != nil {
				panic(err)
			}
			if f < 0 || f >= 1000 {
				panic(fmt.Errorf("invalid fractional part: %s", fractional))
			}
		}
	}
	return uint64(i*1000) + uint64(f)
}

// integerizeDecimalFixed converts a decimal string into an integer.
// decimal in the format: "[[i]i]i[.f[f[f]]]", where
// i is an integral digit, and there are at most 3 digits,
// the decimal point is optional,
// f is a fractional digit, and there are at most 3 digits.
// 0 <= return value < 1,000,000
func integerizeDecimalFixed(decimal string) uint64 {
	return integerizeDecimal(decimal) - 108009
}

// integerizeEnumFixed converts an enum string into an integer.
// enum is just a string.
// 0 <= return value <= 11
func integerizeEnumFixed(enum string) uint64 {
	switch enum {
	case "DUBA":
		return 0
	case "FFS":
		return 1
	case "FXCM":
		return 2
	case "FXDC":
		return 3
	case "FXDD":
		return 4
	case "GAIN":
		return 5
	case "KZ":
		return 6
	case "PEP":
		return 7
	case "PFD":
		return 8
	case "SBD":
		return 9
	case "SEBS":
		return 10
	case "TDFX":
		return 11
	default:
		panic(fmt.Errorf("Unknown enum: %s", enum))
	}
}

var (
	enumMap   = map[string]uint64{}
	enumIndex = uint64(0)
)

// integerizeEnum converts an enum string into an integer.
// enum is just a string.
// The range of the return value depends on how many different strings are read.
func integerizeEnum(enum string) uint64 {
	if v, ok := enumMap[enum]; ok {
		return v
	}
	returnValue := enumIndex
	enumMap[enum] = enumIndex
	enumIndex++
	return returnValue
}

// integerizeRecords converts a Record into an IntegerizedRecord.
func integerizeRecords(integerizeEnumFunc func(string) uint64, records []Record) []IntegerizedRecord {
	iRecords := make([]IntegerizedRecord, len(records))
	for i, record := range records {
		iRecords[i] = IntegerizedRecord{
			IField0: integerizeTime(record.Field0),
			//IField1: integerizeTime(record.Field1),
			IField2: integerizeDecimal(record.Field2),
			IField3: integerizeDecimal(record.Field3),
			IField4: integerizeDecimal(record.Field4),
			IField5: integerizeEnumFunc(record.Field5),
		}
	}
	return iRecords
}
