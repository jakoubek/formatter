package formatter

import (
	"strconv"
	"strings"
)

// FormatInteger formats an int64 value as a number
// with two decimal places, delimited by a comma (,).
// A period (.) is used as a thousand separator for numbers
// >= 1.000.
func FormatInteger(value int64) string {
	sign := ""

	if value < 0 {
		sign = "-"
		// value = 0 - value  ???
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for value > 999 {
		parts[j] = strconv.FormatInt(value%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		value = value / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(value))

	return sign + strings.Join(parts[j:], ".")
}