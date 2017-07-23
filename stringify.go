package main

import "fmt"

func stringifyTime(time uint64) string {
	second := time % 60
	time /= 60
	minute := time % 60
	time /= 60
	hour := time % 24
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func stringifyDecimal(decimal uint64) string {
	fractional := decimal % 1000
	decimal /= 1000
	integral := decimal
	if fractional == 0 {
		return fmt.Sprintf("%d", integral)
	}
	result := fmt.Sprintf("%d.%03d", integral, fractional)
	for result[len(result)-1] == '0' {
		result = result[:len(result)-1]
	}
	return result
}

func stringifyEnumFixed(enum uint64) string {
	switch enum {
	case 0:
		return "DUBA"
	case 1:
		return "FFS"
	case 2:
		return "FXCM"
	case 3:
		return "FXDC"
	case 4:
		return "FXDD"
	case 5:
		return "GAIN"
	case 6:
		return "KZ"
	case 7:
		return "PEP"
	case 8:
		return "PFD"
	case 9:
		return "SBD"
	case 10:
		return "SEBS"
	case 11:
		return "TDFX"
	default:
		panic(fmt.Errorf("Unknown enum: %d", enum))
	}
}
