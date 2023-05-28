package pkerror

import "fmt"

type Error struct {
	reason string
}

func (e *Error) Error() string {
	return e.reason
}

func New(format string, a ...interface{}) *Error {
	return &Error{reason: fmt.Sprintf(format, a...)}
}

func Panic(format string, a ...interface{}) {
	err := &Error{reason: fmt.Sprintf(format, a...)}
	panic(err)
}
