// Copyright (c) 2024 Nicolas-ggd, released under Apache-2.0 License. See LICENSE file.

package parse

import (
	"strconv"
)

var (
	KILO = 1000
	KIB  = 1024
)

func ToInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return v
}

func ToUint64(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}

	return v
}

func ToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return v
}

func ToFloat64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return v
}

func BytesToKB(bytes int64) float64 {
	return float64(bytes) / 1024
}

func KbToGB(kb int64) float64 {
	return float64(kb) * float64(KIB) / (float64(KILO) * float64(KILO) * float64(KILO))
}

func KBToMib(kb int64) float64 {
	return float64(kb) / 1024
}
