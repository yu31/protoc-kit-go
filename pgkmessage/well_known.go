package pgkmessage

import (
	"strings"
)

//// WellKnownTypePackage is the proto package name where all Well Known Types
//// currently reside.
//const WellKnownTypePackage string = "google.protobuf"

// WellKnownType (WKT) encapsulates the Name of a Parser from the
// `google.protobuf` package. Most official protoc plugins special case code
// generation on these messages.
type WellKnownType string

// Name converts the WellKnownType to a Name. This is a convenience method.
func (wkt WellKnownType) Name() string { return string(wkt) }

// Valid returns true if the WellKnownType is recognized by this library.
func (wkt WellKnownType) Valid() bool {
	_, ok := wktLookup[wkt.Name()]
	return ok
}

// 1-to-1 mapping of the WKT names to WellKnownTypes.
const (
	// WKTUnknown indicates that the type is not a known WKT. This value may be
	// returned erroneously mapping a Name to a WellKnownType or if a WKT is
	// added to the `google.protobuf` package but this library is outdated.
	WKTUnknown WellKnownType = "Unknown"

	WKTAny       WellKnownType = "Any"
	WKTTimestamp WellKnownType = "Timestamp"
	WKTDuration  WellKnownType = "Duration"
	WKTEmpty     WellKnownType = "Empty"
	WKTStruct    WellKnownType = "Struct"
	WKTValue     WellKnownType = "Value"

	WKTListValue   WellKnownType = "ListValue"
	WKTDoubleValue WellKnownType = "DoubleValue"
	WKTFloatValue  WellKnownType = "FloatValue"
	WKTInt64Value  WellKnownType = "Int64Value"
	WKTUInt64Value WellKnownType = "UInt64Value"
	WKTInt32Value  WellKnownType = "Int32Value"
	WKTUInt32Value WellKnownType = "UInt32Value"
	WKTBoolValue   WellKnownType = "BoolValue"
	WKTStringValue WellKnownType = "StringValue"
	WKTBytesValue  WellKnownType = "BytesValue"
)

var wktLookup = map[string]WellKnownType{
	"Any":         WKTAny,
	"Timestamp":   WKTTimestamp,
	"Duration":    WKTDuration,
	"Empty":       WKTEmpty,
	"Struct":      WKTStruct,
	"Value":       WKTValue,
	"ListValue":   WKTListValue,
	"DoubleValue": WKTDoubleValue,
	"FloatValue":  WKTFloatValue,
	"Int64Value":  WKTInt64Value,
	"UInt64Value": WKTUInt64Value,
	"Int32Value":  WKTInt32Value,
	"UInt32Value": WKTUInt32Value,
	"BoolValue":   WKTBoolValue,
	"StringValue": WKTStringValue,
	"BytesValue":  WKTBytesValue,
}

// LookupWKT returns the WellKnownType related to the provided Name. If the
// name is not recognized, UnknownWKT is returned.
func LookupWKT(name string) WellKnownType {
	n := strings.TrimPrefix(name, "google.protobuf.")
	if wkt, ok := wktLookup[n]; ok {
		return wkt
	}
	return WKTUnknown
}
