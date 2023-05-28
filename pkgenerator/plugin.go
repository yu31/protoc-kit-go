package pkgenerator

import "google.golang.org/protobuf/compiler/protogen"

// A Plugin provides functionality to add to the output during Go code generation,
// such as to produce RPC stubs.
type Plugin interface {
	// Name identifies the plugin.
	Name() string

	// Version identifies the plugin version.
	Version() string

	// ParamFunc used for accept parameters from the command line,.
	ParamFunc() func(name, value string) error

	// Init is called once before code generated.
	// The `file` will be ignored if return false.
	Init(pp *protogen.Plugin, file *protogen.File) bool

	// Generate to generate codes for specified file.
	// except for the imports, by calling the generator's methods P, In, and Out.
	Generate(g *protogen.GeneratedFile)
}
