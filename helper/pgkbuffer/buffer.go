package pgkbuffer

import (
	"bytes"
)

type Buffer struct {
	BB bytes.Buffer
}

// New returns a new Buffer with default cap.
func New() *Buffer {
	return &Buffer{BB: bytes.Buffer{}}
}

func (x *Buffer) IsEmpty() bool {
	return x == nil || x.BB.Len() == 0
}

func (x *Buffer) WriteStringLn(ss ...string) {
	x.WriteString(ss...)
	_ = x.BB.WriteByte('\n')
}

func (x *Buffer) WriteString(ss ...string) {
	for _, s := range ss {
		_, _ = x.BB.WriteString(s)
	}
}

func (x *Buffer) WriteBytesLn(vv ...[]byte) {
	x.WriteBytes(vv...)
	_ = x.BB.WriteByte('\n')
}

func (x *Buffer) WriteBytes(vv ...[]byte) {
	for _, v := range vv {
		_, _ = x.BB.Write(v)
	}
}

func (x *Buffer) Bytes() []byte {
	if x == nil {
		return nil
	}
	return x.BB.Bytes()
}

func (x *Buffer) String() string {
	if x == nil {
		return ""
	}
	return x.BB.String()
}
