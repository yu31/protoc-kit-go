package pgkerror

import (
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func Recover(pluginName string, file *protogen.File, message *protogen.Message, field *protogen.Field,
	fn func()) (catch bool) {
	defer func() {
		if r := recover(); r != nil {
			catch = true

			msgName := string(message.Desc.FullName())
			if index := strings.Index(msgName, "."); index > 0 {
				msgName = msgName[index+1:]
			}

			fieldName := string(field.Desc.Name())
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				fieldName = string(field.Oneof.Desc.Name())
			}

			id := fmt.Sprintf(
				"%s: { file: %s | message: %s | field: %s }",
				pluginName, file.Desc.Path(), msgName, fieldName,
			)

			if e, ok := r.(*Error); ok {
				println(fmt.Sprintf("ERROR - %s - %s", id, e.Error()))
				return
			}

			println(fmt.Sprintf("unexpected panic on -> %s", id))

			println(fmt.Sprintf("stack: %v", r))
			buf := make([]byte, 8192)
			_ = runtime.Stack(buf, true)

			println(string(buf))
			return
		}
	}()

	fn()
	return
}
