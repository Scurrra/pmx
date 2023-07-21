package dirs

import (
	. "github.com/Scurrra/pmx/common"
)

// Information about the dependency
type Dependency struct {
	// Name of the dependency
	Name string

	// Version of the dependency installed
	Version string

	// Langs the dependency is written in
	PLang PLanguage

	// Type of the dependency: interface, object, static or shared
	Type LibType

	// Path to the dependency file
	Path string

	// No granddepednencies needed
}
