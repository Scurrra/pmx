package pmxconfig

type ProjectType string

const (
	LIB ProjectType = "library"
	BIN             = "binary"
)

// Build options
type BuildOptions struct {
	// Build command, if external tool is used. Runs on `pmx build`.
	// If not specified, internal build system is used.
	Cmd string
	// TODO: smth like preffered compiler, standard, custom options
}

// Project meta information
type ProjectMeta struct {
	Name        string   // `toml:"name"`
	Authors     []string // `toml:"authors"`
	ProjectType ProjectType
	Version     string // `toml:"version"`
	License     string // `toml:"license"`
	Description string // `toml:"description"`
	Repository  string // `toml:"repository"`
}

// Dependencies of the project.
type Dependency struct {
	// Minimal compatible version of the dependency
	Version string

	// Features of the dependency that should be enabled
	Features []string // only for header-only deps I guess

	// Is this dependency optional, i. e. is it enabled for default set of the project dependencies
	Optional bool
}

// Feature specification. If element of a slice starts with "dep:", specified dependency is included,
// otherwise the features automatically enables the other feature with name <element>.
type Feature []string

// Contents of the pmx.toml file.
type Config struct {
	Project      ProjectMeta
	Build        BuildOptions
	Dependencies map[string]Dependency
	Feature      map[string][]Feature
}

func Configure(name string, authors []string, pt ProjectType, version, license, description, repository string) error {
	return nil
}
