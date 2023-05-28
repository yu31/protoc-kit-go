package pkident

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SliceFloat32 returns the string to use for a Go identifier for type of float32 slice.
func SliceFloat32(vv []float32) string {
	var s strings.Builder

	s.WriteString("[]float32{")
	for i, v := range vv {
		s.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 64))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")

	return s.String()
}

// SliceFloat64 returns the string to use for a Go identifier for type of float64 slice.
func SliceFloat64(vv []float64) string {
	var s strings.Builder

	s.WriteString("[]float64{")
	for i, v := range vv {
		s.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")

	return s.String()
}

// SliceInt32 returns the string to use for a Go identifier for type of int32 slice.
func SliceInt32(vv []int32) string {
	var s strings.Builder

	s.WriteString("[]int32{")
	for i, v := range vv {
		s.WriteString(strconv.FormatInt(int64(v), 10))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")
	return s.String()
}

// SliceInt64 returns the string to use for a Go identifier for type of int64 slice.
func SliceInt64(vv []int64) string {
	var s strings.Builder

	s.WriteString("[]int64{")
	for i, v := range vv {
		s.WriteString(strconv.FormatInt(v, 10))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")
	return s.String()
}

// SliceUint32 returns the string to use for a Go identifier for type of uint32 slice.
func SliceUint32(vv []uint32) string {
	var s strings.Builder

	s.WriteString("[]uint32{")
	for i, v := range vv {
		s.WriteString(strconv.FormatUint(uint64(v), 10))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")
	return s.String()
}

// SliceUint64 returns the string to use for a Go identifier for type of uint64 slice.
func SliceUint64(vv []uint64) string {
	var s strings.Builder

	s.WriteString("[]uint64{")
	for i, v := range vv {
		s.WriteString(strconv.FormatUint(v, 10))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")

	return s.String()
}

// SliceBool returns the string to use for a Go identifier for type of bool slice.
func SliceBool(vv []bool) string {
	var s strings.Builder

	s.WriteString("[]bool{")
	for i, v := range vv {
		s.WriteString(strconv.FormatBool(v))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")
	return s.String()
}

// SliceString returns the string to use for a Go identifier for type of string slice.
func SliceString(vv []string) string {
	var s strings.Builder

	s.WriteString("[]string{")
	for i, v := range vv {
		s.WriteString(strconv.Quote(v))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")

	return s.String()
}

// SliceBytes returns the string to use for a Go identifier for type of bytes slice.
func SliceBytes(vv [][]byte) string {
	var ss strings.Builder
	ss.WriteString("[][]byte")
	ss.WriteString("{")

	for i, iv := range vv {
		ss.WriteString("{")
		for j, jv := range iv {
			ss.WriteString(strconv.FormatInt(int64(jv), 10))
			if j < len(iv)-1 {
				ss.WriteString(", ")
			}
		}
		ss.WriteString("}")

		if i < len(vv)-1 {
			ss.WriteString(", ")
		}
	}
	ss.WriteString("}")
	return ss.String()
}

// SliceEnum returns the string to use for a Go identifier for type of enum slice.
func SliceEnum(goType string, vv []int32) string {
	var s strings.Builder

	s.WriteString("[]")
	s.WriteString(goType)
	s.WriteString("{")
	for i, v := range vv {
		s.WriteString(strconv.FormatInt(int64(v), 10))
		if i != len(vv)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")

	return s.String()
}

func SliceDuration(goType string, vv []*durationpb.Duration) string {
	var s strings.Builder

	s.WriteString("[]*")
	s.WriteString(goType)
	s.WriteString("{")

	for i, v := range vv {
		if v != nil {
			s.WriteString(fmt.Sprintf("{Seconds: %d, Nanos: %d}", v.Seconds, v.Nanos))
			//s.WriteString(fmt.Sprintf("&%s{Seconds: %d, Nanos: %d}", goType, v.Seconds, v.Nanos))
		}

		if i < len(vv)-1 {
			s.WriteString(", ")
		}
	}

	s.WriteString("}")

	return s.String()
}

func SliceTimestamp(goType string, vv []*timestamppb.Timestamp) string {
	var s strings.Builder

	s.WriteString("[]*")
	s.WriteString(goType)
	s.WriteString("{")

	for i, v := range vv {
		if v != nil {
			s.WriteString(fmt.Sprintf("{Seconds: %d, Nanos: %d}", v.Seconds, v.Nanos))
			//s.WriteString(fmt.Sprintf("&%s{Seconds: %d, Nanos: %d}", goType, v.Seconds, v.Nanos))
		}

		if i < len(vv)-1 {
			s.WriteString(", ")
		}
	}

	s.WriteString("}")

	return s.String()
}
