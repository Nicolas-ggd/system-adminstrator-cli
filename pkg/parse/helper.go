package parse

import "strconv"

func ToUint64(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}
