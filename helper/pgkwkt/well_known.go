package pgkwkt

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// WKT (Well Known Type) encapsulates the Name of a Parser from the `google.protobuf` package.
// Most official protoc plugins special case code generation on these messages.
type WKT string

// Name converts the Type to a Name. This is a convenience method.
func (wkt WKT) Name() string { return string(wkt) }

// Valid returns true if the WKT is recognized by this package.
func (wkt WKT) Valid() bool {
	_, ok := wktLookup[wkt.Name()]
	return ok
}

// 1-to-1 mapping of the  names to WellKnownTypes.
const (
	// Unknown indicates that the type is not a known . This value may be
	// returned erroneously mapping a Name to a Type or if a  is
	// added to the `google.protobuf` package but this library is outdated.
	Unknown WKT = "Unknown"

	Any       WKT = "google.protobuf.Any"
	Timestamp WKT = "google.protobuf.Timestamp"
	Duration  WKT = "google.protobuf.Duration"
	Empty     WKT = "google.protobuf.Empty"
	Struct    WKT = "google.protobuf.Struct"
	Value     WKT = "google.protobuf.Value"

	ListValue   WKT = "google.protobuf.ListValue"
	DoubleValue WKT = "google.protobuf.DoubleValue"
	FloatValue  WKT = "google.protobuf.FloatValue"
	Int64Value  WKT = "google.protobuf.Int64Value"
	UInt64Value WKT = "google.protobuf.UInt64Value"
	Int32Value  WKT = "google.protobuf.Int32Value"
	UInt32Value WKT = "google.protobuf.UInt32Value"
	BoolValue   WKT = "google.protobuf.BoolValue"
	StringValue WKT = "google.protobuf.StringValue"
	BytesValue  WKT = "google.protobuf.BytesValue"
)

var wktLookup = map[string]WKT{
	"google.protobuf.Any":       Any,
	"google.protobuf.Timestamp": Timestamp,
	"google.protobuf.Duration":  Duration,
	"google.protobuf.Empty":     Empty,
	"google.protobuf.Struct":    Struct,
	"google.protobuf.Value":     Value,

	"google.protobuf.ListValue":   ListValue,
	"google.protobuf.DoubleValue": DoubleValue,
	"google.protobuf.FloatValue":  FloatValue,
	"google.protobuf.Int64Value":  Int64Value,
	"google.protobuf.UInt64Value": UInt64Value,
	"google.protobuf.Int32Value":  Int32Value,
	"google.protobuf.UInt32Value": UInt32Value,
	"google.protobuf.BoolValue":   BoolValue,
	"google.protobuf.StringValue": StringValue,
	"google.protobuf.BytesValue":  BytesValue,
}

// Lookup returns the Type related to the provided Name. If the
// name is not recognized, Unknown is returned.
func Lookup(name string) WKT {
	//n := strings.TrimPrefix(name, "google.protobuf.")
	if mkt, ok := wktLookup[name]; ok {
		return mkt
	}
	return Unknown
}

// Valid returns true if the field is contains in this package.
// Notice: We also return true (Due to inability to make judgments?) if the field type is `oneof` and it is the parent of 'oneof'.
func Valid(field *protogen.Field) bool {
	var msgName string
	switch {
	case field.Desc.IsMap() && field.Desc.MapValue().Kind() == protoreflect.MessageKind:
		msgName = string(field.Message.Fields[1].Message.Desc.FullName())
	case !field.Desc.IsMap() && field.Desc.Kind() == protoreflect.MessageKind:
		msgName = string(field.Message.Desc.FullName())
	}
	return Lookup(msgName).Valid()
}
