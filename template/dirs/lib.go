package dirs

import (
	"fmt"
	"os"
	"strings"

	. "github.com/Scurrra/pmx/common"
	. "github.com/eminarican/safetypes"
)

// include directory contents a.k.a. public
type LibraryCode struct {
	// All private headers
	Headers Option[Headers]

	// C++20 module files
	Modules Option[Modules] // idk, but let it be here

	// ...
	Binaries Option[Binaries]

	// All types of allowed files in include/
	CodeFiles []CodeFile
}

// Create empty directory `include` with subdirectories
func (lc LibraryCode) CreateDirs() error {
	err_dir := os.Mkdir("src", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	err_dir = os.Mkdir("headers", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	err_dir = os.Mkdir("modules", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	err_dir = os.Mkdir("binaries", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	return nil
}

// Create empty directory `include` with subdirectories and template for header only library
func (lc LibraryCode) CreateTemplateHeader(name string, plang PLanguage) error {
	if _, err := os.Stat("include"); os.IsNotExist(err) {
		err_dir := lc.CreateDirs()
		if err_dir != nil {
			return err_dir
		}
	}

	switch plang {
	case C:
		NAME := strings.ToUpper(name) + "_LIB"
		lib_template := fmt.Sprintf(`
// Public API of your library
#ifdef %s
#define %s

// Write your code here

#endif %s
		`, NAME, NAME, NAME)

		f, err_f := os.Create(fmt.Sprintf("include/%s.h", name))
		if err_f != nil {
			return err_f
		}
		defer f.Close()

		_, err_f = f.Write([]byte(lib_template))
		if err_f != nil {
			return err_f
		}
	case CPP:
		lib_template := `
// Public API of your library
#pragma once
// Write your code here
		`

		f, err_f := os.Create(fmt.Sprintf("include/%s.hpp", name))
		if err_f != nil {
			return err_f
		}
		defer f.Close()

		_, err_f = f.Write([]byte(lib_template))
		if err_f != nil {
			return err_f
		}
	case CC:
		NAME := strings.ToUpper(name) + "_LIB"
		lib_template := fmt.Sprintf(`
// Public API of your library
#ifdef %s
#define %s

// Write your code here

#endif %s
		`, NAME, NAME, NAME)

		f, err_f := os.Create(fmt.Sprintf("include/%s.hxx", name))
		if err_f != nil {
			return err_f
		}
		defer f.Close()

		_, err_f = f.Write([]byte(lib_template))
		if err_f != nil {
			return err_f
		}
	}

	return nil
}

// TODO: templates for other types of libraries
