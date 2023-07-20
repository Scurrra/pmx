package template

import (
	"github.com/Scurrra/docsync/dsconfig"
	"github.com/Scurrra/pmx/dirs"
	"github.com/Scurrra/pmx/pmxconfig"
	. "github.com/eminarican/safetypes"
)

// Basic project structure.
type ProjectStructure struct {
	// Documentation of the project
	Docs Option[dsconfig.Config]

	// Tests for the project code
	Test Option[dirs.Test]

	// Public interface of the library (--lib)
	Include Option[dirs.LibraryCode]

	// Code of the application (--bin)
	Src Option[dirs.SourceCode]

	// Build files of the application
	Build Option[dirs.Build]

	// Configuration of the project
	Config pmxconfig.Config
}
