package main

import "fmt"

// stringifyTime is the reverse of integerizeTime.
// It converts an integer into a time string.
// 0 <= time < 86,400
// "00:00:00" <= return value <= "23:59:59"
func stringifyTime(time uint64) string {
	if time >= 86400 {
		panic("time >= 86400")
	}
	second := time % 60
	time /= 60
	minute := time % 60
	time /= 60
	hour := time
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

// stringifyDecimal is the reverse of integerizeDecimal.
// It convers an integer into a decimal string.
// 0 <= decimal < 1,000,000
// "0.000" <= return value <= "999.999"
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

// stringifyEnumFixed is the reverse of integerizeEnumFixed.
// It converts an integer into an enum string.
// 0 <= enum <= 11
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
