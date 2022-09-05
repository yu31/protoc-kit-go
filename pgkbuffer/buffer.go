package pgkbuffer

import (
	"unsafe"
)

type Buffer struct {
	buf []byte
}

// New returns a new Buffer with default cap.
func New() *Buffer {
	return &Buffer{buf: make([]byte, 0, 64)}
}

func (x *Buffer) IsEmpty() bool {
	return x == nil || len(x.buf) == 0
}

func (x *Buffer) WriteStringLn(ss ...string) {
	x.WriteString(ss...)
	x.buf = append(x.buf, '\n')
}

func (x *Buffer) WriteString(ss ...string) {
	for _, s := range ss {
		x.buf = append(x.buf, s...)
	}
}

func (x *Buffer) WriteBytesLn(vv ...[]byte) {
	x.WriteBytes(vv...)
	x.buf = append(x.buf, '\n')
}

func (x *Buffer) WriteBytes(vv ...[]byte) {
	for _, v := range vv {
		x.buf = append(x.buf, v...)
	}
}

func (x *Buffer) Bytes() []byte {
	if x == nil {
		return nil
	}
	return x.buf
}

func (x *Buffer) String() string {
	if x == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&x.buf))
}
