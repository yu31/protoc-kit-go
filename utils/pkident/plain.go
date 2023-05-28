package pkident

import (
	"strconv"
	"strings"
)

// Float32 returns the string to use for a Go identifier for type of float32.
func Float32(v float32) string {
	return strconv.FormatFloat(float64(v), 'f', -1, 64)
}

// Float64 returns the string to use for a Go identifier for type of float64.
func Float64(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// Int32 returns the string to use for a Go identifier for type of int32.
func Int32(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

// UInt32 returns the string to use for a Go identifier for type of uint32.
func UInt32(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Int64 returns the string to use for a Go identifier for type of int64.
func Int64(v int64) string {
	return strconv.FormatInt(v, 10)
}

// UInt64 returns the string to use for a Go identifier for type of uint64.
func UInt64(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// Bool returns the string to use for a Go identifier for type of bool.
func Bool(v bool) string {
	return strconv.FormatBool(v)
}

// String returns the string to use for a Go identifier for type of string.
func String(v string) string {
	return strconv.Quote(v)
}

// Bytes returns the string to use for a Go identifier for type of bytes.
func Bytes(vv []byte) string {
	var s strings.Builder
	s.WriteString("[]byte")
	s.WriteString("{")

	for i, v := range vv {
		s.WriteString(strconv.FormatInt(int64(v), 10))
		if i < len(vv)-1 {
			s.WriteString(",")
		}
	}
	s.WriteString("}")
	return s.String()
}
