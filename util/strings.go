package util

import (
	"math"
	"strconv"
	"strings"
)

func EqualAny(str string, any []string) bool {
	for _, s := range any {
		if s == str {
			return true
		}
	}
	return false
}

func HasAny(str string, any []string) (string, bool) {
	for _, s := range any {
		if s == str {
			return s, true
		}
	}
	return "", false
}

func NumberParts(v int64, delim string) string {
	sign := ""

	// Min int64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], delim)
}
