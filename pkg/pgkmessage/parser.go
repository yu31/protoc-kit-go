package pgkmessage

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Parser for parser the protogen.Message
type Parser struct {
	// The proto message.
	message *protogen.Message

	// plainMap store all fields except that in `oneof` labels.
	plainMap map[string]*protogen.Field

	// oneOfMap store all fields that in `oneof` labels.
	oneOfMap map[string]*protogen.Field

	// fields store the all field in message.
	fields []*protogen.Field
}

func NewParser(message *protogen.Message) *Parser {
	x := &Parser{
		message: message,
	}
	x.init()
	return x
}

func (x *Parser) init() {
	message := x.message
	nLen := len(message.Fields)

	// plainMap store all fields except that in `oneof` labels.
	plainMap := make(map[string]*protogen.Field, nLen)
	// oneOfMap store all fields that in `oneof` labels.
	oneOfMap := make(map[string]*protogen.Field)
	// fields store the all field in message.
	fields := make([]*protogen.Field, 0, nLen)

	for _, field := range message.Fields {
		if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
			if field.Oneof.Fields[0] != field {
				continue // only generate for first appearance
			}
			plainMap[string(field.Oneof.Desc.Name())] = field
			for _, oneOfField := range field.Oneof.Fields {
				oneOfMap[string(oneOfField.Desc.Name())] = oneOfField
			}
		} else {
			plainMap[string(field.Desc.Name())] = field
		}
		fields = append(fields, field)
	}

	x.plainMap = plainMap
	x.oneOfMap = oneOfMap
	x.fields = fields
}

// Fields returns all fields except those in "oneof".
func (x *Parser) Fields() []*protogen.Field {
	return x.fields
}

// LookupByName get the field by name. returns nil if not found.
func (x *Parser) LookupByName(name string) (field *protogen.Field, inOneOf bool) {
	var ok bool
	if field, ok = x.plainMap[name]; ok {
		return
	}
	if field, ok = x.oneOfMap[name]; ok {
		inOneOf = true
		return
	}
	return
}
