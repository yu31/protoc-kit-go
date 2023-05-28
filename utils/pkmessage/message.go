package pkmessage

import "google.golang.org/protobuf/compiler/protogen"

// LoadValidMessages get a valid message lists recursively from giving message lists.
func LoadValidMessages(messages []*protogen.Message) []*protogen.Message {
	var loadMessages func(messages []*protogen.Message)

	var validMessages []*protogen.Message

	loadMessages = func(messages []*protogen.Message) {
		for _, msg := range messages {
			if msg.Desc.IsMapEntry() {
				// Ignore protobuf's map entry.
				continue
			}

			validMessages = append(validMessages, msg)

			loadMessages(msg.Messages)
		}
	}

	loadMessages(messages)

	return validMessages
}

// LoadFieldLists returns valid field lists in message.
func LoadFieldLists(message *protogen.Message) []*protogen.Field {
	fields := make([]*protogen.Field, 0)
	for _, field := range message.Fields {
		if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() && field.Oneof.Fields[0] != field {
			continue // only generate for first appearance
		}
		fields = append(fields, field)
	}
	return fields
}
