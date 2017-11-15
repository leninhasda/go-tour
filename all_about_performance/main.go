package main

import (
	"math"
	"strconv"
)

// format takes 2 arguments and output a formatted string
// sample: format(1,2.0) => "formated-int(1)-float(2.00)" the float number shall always have 2 digits after decimal point.
func format(intValue int64, floatValue float64) string {
	// si := strconv.FormatInt(intValue, 10)
	si := strconv.Itoa(int(intValue))
	// sf := strconv.FormatFloat(floatValue, 'f', 2, 64)
	// sf := fmt.Sprintf("%v", floatValue)
	fi, fd := math.Modf(floatValue)
	fii, fid := int(fi), int(fd)
	sfii := strconv.Itoa(fii)
	sfid := strconv.Itoa(fid)
	if sfid == "0" {
		sfid = "00"
	}
	// sfii, sfid := string(fii), string(fid)
	// return fmt.Sprintf("formatted-int(%d)-float(%d.%s)", intValue, fii, sfid)
	return "formatted-int(" + si + ")-float(" + sfii + "." + sfid + ")"
	// return "formatted-int(" + si + ")-float(100.00)"
	// return "formatted-int(" + si + ")-float(" + sf + ")"
	// return "formatted-int(" + strconv.FormatInt(intValue, 10) + ")-float(" + strconv.FormatFloat(floatValue, 'f', 2, 64) + ")"
}
