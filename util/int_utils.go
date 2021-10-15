package util

import "strconv"

const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

func CompareMax(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func IsDigital(s string) (int, bool) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, false
	}
	return int(i), true
}
